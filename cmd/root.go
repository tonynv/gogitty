package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	c "github.com/tonynv/gogitty/config"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gogitty",
	Short: "A GitHub Utility",
	Long:  `A GitHub Utility to interact with the API written in Golang`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}
var cfgFile string
var programName = "GOGITTY"
var configuration c.Configurations

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gogitty.yaml)")
}

func initConfig() {
	fmt.Println("Initializing... Gitty up!!")
	// Set the file name of the configurations file
	viper.SetConfigName(".gogitty")

	// Get the UserHomeDir
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to User Home Directory , %v", err)
	} else {
		viper.AddConfigPath(home)
	}

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("GitHubCfg.ApiEndpoint", "github.com")

	cfgErr := viper.Unmarshal(&configuration)
	if cfgErr != nil {
		fmt.Printf("Unable to decode into struct, %v", cfgErr)
	}

}
