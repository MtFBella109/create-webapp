package main

import (
	"fmt"
	"reflect"

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

	fmt.Println(reflect.TypeOf(config))
	return &config
}
