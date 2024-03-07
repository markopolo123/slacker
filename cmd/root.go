/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	internal "github.com/markopolo123/slacker/internal"
	models "github.com/markopolo123/slacker/models"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "slacker",
	Short: "Set Slack statuses quicker than a cheetah on a caffeine rush",
	Long: `helper utility for setting Slack statuses quickly and easily:

Set your desired statuses in the config.toml file and then call slacker set $name.
For example, if you have a status called "lunch" in your config.toml file, you can
set that status by running "slacker lunch".`,
}

func Execute() {
	var config *models.Config
	dir, _ := os.UserHomeDir()
	path := filepath.Join(dir, ".config", "slacker", "config.toml")

	config, err := internal.LoadConfig(path)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// Dynamically create and add commands based on the loaded configuration
	internal.CreateShortcutCommands(rootCmd, config.Shortcuts)

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
