package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type configstruct struct {
	first_launch             bool
	generate_desktop_file    bool
	systemwide_desktop_entry bool
	webapps_directory        string
	locale                   string
}

func Config() *configstruct {
	configfile := viper.New()
	configfile.SetConfigFile("config.toml")
	configfile.ReadInConfig()

	config := configstruct{
		first_launch:             configfile.GetBool("general.first_launch"),
		generate_desktop_file:    configfile.GetBool("general.generate_desktop_file"),
		systemwide_desktop_entry: configfile.GetBool("general.systemwide_desktop_entry"),
		webapps_directory:        configfile.GetString("general.webapps_directory"),
		locale:                   configfile.GetString("general.locale"),
	}
	if config.first_launch == true {
		newconfig := writeConfig(&config)
		config = *newconfig
		config.first_launch = false
		configfile.Set("general.first_launch", config.first_launch)
		configfile.Set("general.generate_desktop_file", config.generate_desktop_file)
		configfile.Set("general.systemwide_desktop_entry", config.systemwide_desktop_entry)
		configfile.Set("general.webapps_directory", config.webapps_directory)
		configfile.Set("general.locale", config.locale)
		configfile.WriteConfig()
	}

	return &config
}

func writeConfig(config *configstruct) *configstruct {
	fmt.Println("1: Generate Desktop File (Default is true)")
	fmt.Println("2: Make the Desktop entry systemwide available (Default is false)")
	fmt.Println("3: In which Directory should all WebApps go in? ATTENTION: This Directory has to exist, otherwise the Program will not work correctly (Default is ~/WebApps)")
	fmt.Println("4: Change the locale (Default is en)")
	editconf := true
	for editconf == true {
		fmt.Println("Type the Number of the Config you want to change or type 'done' if everything is set correctly")
		var confuserinput string
		switch fmt.Scanln(&confuserinput); confuserinput {
		case "done":
			editconf = false
		case "1":
			if config.generate_desktop_file == true {
				fmt.Println("Changed now Generate Desktop file to False")
				config.generate_desktop_file = false
			} else {
				fmt.Println("Changed now Generate Desktop file to True")
				config.generate_desktop_file = true
			}
		case "2":
			if config.systemwide_desktop_entry == true {
				fmt.Println("Changed now Generate Desktop file to False")
				config.systemwide_desktop_entry = false
			} else {
				fmt.Println("Changed now Generate Desktop file to True")
				config.systemwide_desktop_entry = true
			}
		case "3":
			fmt.Println("Please type the Directory where the WebApps should be created")
			var directory string
			fmt.Scanln(&directory)
			config.webapps_directory = directory
			fmt.Println("Changed to " + directory)
		case "4":
			fmt.Println("Type your wanted locale You can choose: 'de', 'en'")
			var locale string
			fmt.Scanln(&locale)
			config.locale = locale
			fmt.Println("Changed to " + locale)
		}
	}
	return config
}
