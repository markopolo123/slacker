/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/markopolo123/slacker/internal"
	models "github.com/markopolo123/slacker/models"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config *models.Config
		dir, _ := os.UserHomeDir()
		path := filepath.Join(dir, ".config", "slacker", "config.toml")

		config, err := internal.LoadConfig(path)
		if err != nil {
			fmt.Println("Error loading config:", err)
			os.Exit(1)
		}
		set(config.Shortcuts)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var pickedShortcut string

func set(shortcuts []models.Shortcut) {
	var options []huh.Option[string]

	for _, sc := range shortcuts {
		option := huh.NewOption[string](sc.Name, sc.Name)
		options = append(options, option)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your status...").
				Options(options...).
				Value(&pickedShortcut),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	// Find the corresponding Shortcut object by name
	for _, sc := range shortcuts {
		if sc.Name == pickedShortcut {
			internal.SetStatus(sc)
			return
		}
	}
}
