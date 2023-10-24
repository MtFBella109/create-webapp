package main
//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,de

import (
	"fmt"
	"os/user"
	"strings"
  "os"
  "golang.org/x/text/language"
  "golang.org/x/text/message"
	"github.com/spf13/viper"
)

type configstruct struct {
	first_launch             bool
	generate_desktop_file    bool
	systemwide_desktop_entry bool
	webapps_directory        string
	locale                   string
  specialos string
}

func Config() *configstruct {
  configfile := viper.New()
  configfile.SetConfigFile("config/config.toml")
  configfile.ReadInConfig()

  config := &configstruct{
    first_launch:             configfile.GetBool("general.first_launch"),
    generate_desktop_file:    configfile.GetBool("general.generate_desktop_file"),
    systemwide_desktop_entry: configfile.GetBool("general.systemwide_desktop_entry"),
    webapps_directory:        configfile.GetString("general.webapps_directory"),
    locale:                   configfile.GetString("general.locale"),
    specialos:                configfile.GetString("advanced.specialos"),
  }

  if config.first_launch {
                writeConfig(config, configfile)
  }

  return config
}

func writeConfig(config *configstruct, configfile *viper.Viper) {
  p := message.NewPrinter(language.German)
  p.Println("1: Generate Desktop File (Default is true)")
  p.Println("2: Make the Desktop entry systemwide available (Default is false)")
  p.Println("3: In which Directory should all WebApps go in? ATTENTION: This Directory has to exist, otherwise the Program will not work correctly (Default is ~/WebApps)")
  p.Println("4: Change the locale (Default is en)")
  p.Println("5: Change SpecialOS, you can only choose between NixOS and none")
  editconf := true
  for editconf {
    p.Println("Type the Number of the Config you want to change or type 'done' if everything is set correctly")
    var confuserinput string
    switch fmt.Scanln(&confuserinput); confuserinput {
    case "done":
      editconf = false
    case "1":
      if config.generate_desktop_file == true {
        p.Println("Changed now Generate Desktop file to False")
        config.generate_desktop_file = false
      } else {
        p.Println("Changed now Systemwide Desktop file to True")
        config.generate_desktop_file = true
      }
    case "2":
      if config.systemwide_desktop_entry == true {
        p.Println("Changed now Systemwide Desktop file to False")
        config.systemwide_desktop_entry = false
      } else {
        p.Println("Changed now Generate Desktop file to True")
        config.systemwide_desktop_entry = true
      }
    case "3":
      p.Println("Please type the Directory where the WebApps should be created")
      var directory string
      fmt.Scanln(&directory)
      if strings.HasPrefix(directory, "~/") {
        User, err := user.Current()
        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }
        strings.Replace(directory, "~", "/home/" + User.Username + "/", 1)
      }
      config.webapps_directory = directory
      p.Println("Changed to " + directory)
    case "4":
      p.Println("Type your wanted locale You can choose: 'de', 'en'")
      var locale string
      fmt.Scanln(&locale)
      config.locale = locale
      p.Println("Changed to " + locale)
    case "5":
      p.Println("Type your wanted SpecialOS You can choose: 'NixOS', 'none'")
      var SpecialOS string
      fmt.Scanln(&SpecialOS)
      SpecialOS = strings.ToLower(SpecialOS)
      if SpecialOS == "nixos" {
        config.specialos = "nixos"
      } else {
        config.specialos = "none"
      }
      p.Println("Changed to " + SpecialOS)
    }
  }
  configfile.Set("general.first_launch", false)
  configfile.Set("general.generate_desktop_file", config.generate_desktop_file)
  configfile.Set("general.systemwide_desktop_entry", config.systemwide_desktop_entry)
  configfile.Set("general.webapps_directory", config.webapps_directory)
  configfile.Set("general.locale", config.locale)
  configfile.Set("advanced.specialos", config.specialos)
  configfile.WriteConfig()
}
