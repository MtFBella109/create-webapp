package main

import (
  "fmt"
  "io"
  "os"
  "os/exec"
  "path/filepath"
  "strings"
  "syscall"
)


func createWebApp(config *configstruct, webappName, webappURL string) {
  cwd, _ := syscall.Getwd()
  syscall.Chdir(config.webapps_directory)
  defer syscall.Chdir(cwd)

  if !strings.HasPrefix(webappURL, "https://") && !strings.HasPrefix(webappURL, "http://") {
    webappURL = "https://" + webappURL
  }

  err := os.Mkdir(webappName, 0755)
  if err != nil {
    fmt.Println(err)
    return
  }

  exec.Command("cp", filepath.Join(cwd, "templates", "WebAppTemplate", "index.js"), filepath.Join(config.webapps_directory, webappName)).Run()
  exec.Command("cp", filepath.Join(cwd, "templates", "WebAppTemplate", "package.json"), filepath.Join(config.webapps_directory, webappName)).Run()
  err = os.Chdir(filepath.Join(config.webapps_directory, webappName))
  if err != nil {
    fmt.Println(err)
    return
  }

  indexfile, err := os.Open("index.js")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer indexfile.Close()

  indexfiledata, err := io.ReadAll(indexfile)
  if err != nil {
    fmt.Println(err)
    return
  }

  indexfiledata = []byte(strings.ReplaceAll(string(indexfiledata), "WBTITLE", webappName))
  indexfiledata = []byte(strings.ReplaceAll(string(indexfiledata), "WBURL", webappURL))

  err = os.WriteFile("index.js", indexfiledata, 0755)
  if err != nil {
    fmt.Println(err)
    return
  }

  exec.Command("npm", "install").Run()

  if config.generate_desktop_file && config.specialos == "none" {
    exec.Command("cp", filepath.Join(cwd, "templates", "template.desktop"), filepath.Join(config.webapps_directory, webappName+".desktop")).Run()
    desktopFile := webappName + ".desktop"
    file, err := os.Open(desktopFile)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer file.Close()

    filedata, err := io.ReadAll(file)
    if err != nil {
      fmt.Println(err)
      return
    }

    filedata = []byte(strings.ReplaceAll(string(filedata), "$NAME", webappName))
    filedata = []byte(strings.ReplaceAll(string(filedata), "$PATH", filepath.Join(config.webapps_directory, webappName)))

    err = os.WriteFile(webappName+".desktop", filedata, 0644)
    if err != nil {
      fmt.Println(err)
      return
    }

    if config.systemwide_desktop_entry {
      err = exec.Command("sudo", "mv", filepath.Join(config.webapps_directory, webappName+".desktop"), "/usr/share/applications/").Run()
      if err != nil {
        fmt.Println(err)
        return
      }
    } else {
      err = exec.Command("mv", filepath.Join(config.webapps_directory, webappName+".desktop"), filepath.Join("~/.local/share/applications/")).Run()
      if err != nil {
        fmt.Println(err)
        return
      }
    }

    fmt.Println("WebApp was successfully created")
  } else if config.generate_desktop_file && config.specialos == "nixos" {
    fmt.Println("Function to add NixOS support is not yet implemented")
  }
}


func main() {
  config := Config()

  var webappName string
  var webappURL string

  fmt.Println("What's the Name of the WebApp?")
  fmt.Scanln(&webappName)
  fmt.Println("What's the URL of the WebApp?")
  fmt.Scanln(&webappURL)

  createWebApp(config, webappName, webappURL)
}
