package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("[%s] Show Configuration\n", programName)
		// Reading variables using the model
		fmt.Println("Reading variables using the model..")
		fmt.Println("GitHub Endpoint is\t", configuration.GitHubCfg.ApiEndpoint)
		fmt.Println("GitHub Token is\t", configuration.GitHubCfg.ApiToken)

		// Reading variables without using the model
		//fmt.Println("\nReading variables without using the model..")
		//fmt.Println("Github hostname is\t", viper.Get("GitHubApi.Endpoint"))
		//fmt.Println("Git Access Token is\t", viper.Get("GitHubApi.token"))

	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
