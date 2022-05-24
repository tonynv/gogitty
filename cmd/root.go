package cmd

// RootCmd represents the base command when called without any subcommands

import (
	"github.com/spf13/cobra"
)

const version = "0.0.0-alpha1"

var RootCmd = &cobra.Command{
	Version: version,
	Use:     "gogitty",
	Short:   "A simple gogitty",
	Long:    "gogitty",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute)
}
func init() {}
