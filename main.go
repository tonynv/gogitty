/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	c "github.com/tonynv/gogitty/config"
	"os"
)

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName(".gogitty")

	// Set the path to look for the configurations file
	home, _err_homedir := os.UserHomeDir()
	cobra.CheckErr(_err_homedir)
	viper.AddConfigPath(home)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("endpoint.hostname", "github.com")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("GitHub Endpoint is\t", configuration.Endpoint.Hostname)
	fmt.Println("GitHub Token is\t", configuration.GIT_ACCESS_TOKEN)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Githib hostname is\t", viper.Get("Endpoint.hostname"))
	fmt.Println("Git Access Token is\t", viper.Get("GIT_ACCESS_TOKEN"))
	//cmd.Execute()

}
