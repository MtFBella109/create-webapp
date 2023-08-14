package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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
	cmd := exec.Command("nativefier", "--name", webappName, webappURL)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	if config.generate_desktop_file == true {
		exec.Command("cp", cwd+"/templates/template.desktop", "./"+webappName+".desktop")
		desktop_file := webappName + ".desktop"
		file, _ := os.Open(desktop_file)
		filedata, _ := io.ReadAll(file)
		filedata = []byte(strings.ReplaceAll(string(filedata), "$NAME", webappName))
		filedata = []byte(strings.ReplaceAll(string(filedata), "$PATH", config.webapps_directory+"/"+webappName+"-linux-x64/"))
		os.WriteFile(webappName+".desktop", filedata, 0644)
		if config.systemwide_desktop_entry == true {
			exec.Command("sudo", "mv", "./*.desktop", "/usr/share/applications/")
		} else {
			exec.Command("mv", "*.desktop", "~/.local/share/applications/")
		}
		fmt.Println("WebApp was sucesfully created")
	}
	exec.Command("cd %s", cwd)
}
