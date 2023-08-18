package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	// "runtime"
)

func main() {
	var webappName string
	var webappURL string

	config := Config()
	cwd, _ := syscall.Getwd()
	syscall.Chdir(config.webapps_directory)
	fmt.Println("Whats the Name of the WebApp?")
	fmt.Scanln(&webappName)
	fmt.Println("Whats the URL of the WebApp?")
	fmt.Scanln(&webappURL)
	os.Chdir(filepath.Join(config.webapps_directory))
	exec.Command("nativefier", "--name", webappName, webappURL).Run()
	if config.generate_desktop_file == true {
		exec.Command("cp", filepath.Join(cwd, "templates", "template.desktop"), filepath.Join(config.webapps_directory, webappName+".desktop")).Run()
		desktop_file := webappName + ".desktop"
		file, _ := os.Open(desktop_file)
		filedata, _ := io.ReadAll(file)
		filedata = []byte(strings.ReplaceAll(string(filedata), "$NAME", webappName))
		filedata = []byte(strings.ReplaceAll(string(filedata), "$PATH", config.webapps_directory+"/"+webappName+"-linux-x64"))
		os.WriteFile(webappName+".desktop", filedata, 0644)
		if config.systemwide_desktop_entry == true {
			cmd := exec.Command("sudo", "mv", filepath.Join(config.webapps_directory, webappName+".desktop"), "/usr/share/applications/")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			cmd := exec.Command("mv", filepath.Join(config.webapps_directory, webappName+".desktop"), "~/.local/share/applications/")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("WebApp was sucesfully created")
	}
	exec.Command("cd %s", cwd).Run()
}
