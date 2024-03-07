/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/huh"
	models "github.com/markopolo123/slacker/models"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// create config.toml in ~/.config/slacker

var path = "~/.config/slacker"
var file = "config.toml"
var filePath = path + "/" + file
var overwrite = false

func initConfig() error {
	expandedPath, err := expandHomeDir(filePath)
	if err != nil {
		return err
	}
	_, err = os.Stat(expandedPath)
	if err != nil {
		os.Mkdir(path, 0700)
		templateConfig()
		fmt.Println("File created at", filePath)
		return nil
	} else {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("File already exists... overwrite?").
					Value(&overwrite),
			),
		)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		if overwrite {
			templateConfig()
		}
		return nil
	}
}

func templateConfig() {
	shortcuts := []models.Shortcut{
		{Name: "desk", Text: "At my Desk", Emoji: ":desktop_computer:", Duration: "1h"},
		{Name: "lunch", Text: "Gone for lunch", Emoji: ":pizza:", Duration: "60m"},
		{Name: "dnd", Text: "Do Not Disturb", Emoji: ":warning:", Duration: "60m"},
	}
	config := models.Config{Shortcuts: shortcuts}

	err := writeTomlFile(filePath, config)

	if err != nil {
		fmt.Println("Error writing TOML file:", err)
	}
}

func expandHomeDir(path string) (string, error) {
	if path[:2] == "~/" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		dir := usr.HomeDir
		path = filepath.Join(dir, path[2:])
	}
	return path, nil
}

func writeTomlFile(filePath string, config models.Config) error {
	expandedPath, err := expandHomeDir(filePath)
	if err != nil {
		return err
	}

	file, err := os.Create(expandedPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
