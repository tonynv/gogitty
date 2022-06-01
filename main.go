/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tonynv/gogitty/cmd"
	c "github.com/tonynv/gogitty/config"
	"os"
)

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName(".gogitty")

	// Set the path to look for the configurations file
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to User Home Directory , %v", err)
	}
	viper.AddConfigPath(home)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("GitHubCfg.ApiEndpoint", "github.com")

	cfgErr := viper.Unmarshal(&configuration)
	if cfgErr != nil {
		fmt.Printf("Unable to decode into struct, %v", cfgErr)
	}

	// Reading variables using the model
	//fmt.Println("Reading variables using the model..")
	//fmt.Println("GitHub Endpoint is\t", configuration.GitHubCfg.ApiEndpoint)
	//fmt.Println("GitHub Token is\t", configuration.GitHubCfg.ApiToken)

	// Reading variables without using the model
	//fmt.Println("\nReading variables without using the model..")
	//fmt.Println("Github hostname is\t", viper.Get("GitHubApi.Endpoint"))
	//fmt.Println("Git Access Token is\t", viper.Get("GitHubApi.token"))
	cmd.Execute()

}
