package main

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,de

import (
  "fmt"
  "io"
  "os"
  "os/exec"
  "path/filepath"
  "strings"
  "syscall"
  "golang.org/x/text/language"
  "golang.org/x/text/message"
)


func createWebApp(config *configstruct, webappName, webappURL string, P *message.Printer) {
  cwd, _ := syscall.Getwd()
  err := syscall.Chdir(config.webapps_directory)
  errCheck(err)

  if !strings.HasPrefix(webappURL, "https://") && !strings.HasPrefix(webappURL, "http://") {
    webappURL = "https://" + webappURL
  }

  err = os.Mkdir(webappName, 0755)
  errCheck(err)

  exec.Command("cp", filepath.Join(cwd, "templates", "WebAppTemplate", "index.js"), filepath.Join(config.webapps_directory, webappName)).Run()
  exec.Command("cp", filepath.Join(cwd, "templates", "WebAppTemplate", "package.json"), filepath.Join(config.webapps_directory, webappName)).Run()
  err = os.Chdir(filepath.Join(config.webapps_directory, webappName))
  errCheck(err)

  indexfile, err := os.Open("index.js")
  errCheck(err)
  defer indexfile.Close()

  indexfiledata, err := io.ReadAll(indexfile)
  errCheck(err)

  indexfiledata = []byte(strings.ReplaceAll(string(indexfiledata), "WBTITLE", webappName))
  indexfiledata = []byte(strings.ReplaceAll(string(indexfiledata), "WBURL", webappURL))

  err = os.WriteFile("index.js", indexfiledata, 0755)
  errCheck(err)

  err = exec.Command("npm", "install").Run()
  if err != nil {
    panic(err)
  }

  if config.generate_desktop_file && config.specialos == "none" {
    err =exec.Command("cp", filepath.Join(cwd, "templates", "template.desktop"), filepath.Join(config.webapps_directory, webappName+".desktop")).Run()
    errCheck(err)
    desktopFile := webappName + ".desktop"
    file, err := os.Open(desktopFile)
    errCheck(err)
    defer file.Close()

    filedata, err := io.ReadAll(file)
    errCheck(err)

    filedata = []byte(strings.ReplaceAll(string(filedata), "$NAME", webappName))
    filedata = []byte(strings.ReplaceAll(string(filedata), "$PATH", filepath.Join(config.webapps_directory, webappName)))

    err = os.WriteFile(webappName+".desktop", filedata, 0644)
    errCheck(err)

    if config.systemwide_desktop_entry {
      err = exec.Command("sudo", "mv", filepath.Join(config.webapps_directory, webappName+".desktop"), "/usr/share/applications/").Run()
      errCheck(err)
    } else {
      err = exec.Command("mv", filepath.Join(config.webapps_directory, webappName+".desktop"), filepath.Join("~/.local/share/applications/")).Run()
      errCheck(err)
    }

    P.Println("WebApp was successfully created")
  } else if config.generate_desktop_file && config.specialos == "nixos" {
    P.Println("Function to add NixOS support is not yet implemented")
  }
  defer syscall.Chdir(cwd)
}

func errCheck(err error) {
  if err != nil {
    fmt.Println(err)
    return
  }
}

func main() {
  config := Config()

  var webappName string
  var webappURL string

  P := message.NewPrinter(language.German)

  P.Println("What's the Name of the WebApp?")
  fmt.Scanln(&webappName)
  P.Println("What's the URL of the WebApp?")
  fmt.Scanln(&webappURL)

  createWebApp(config, webappName, webappURL, P)
}
