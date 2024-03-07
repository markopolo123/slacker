/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/enescakir/emoji"
	internal "github.com/markopolo123/slacker/internal"
	models "github.com/markopolo123/slacker/models"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "List your shortcuts",
	Long:  `List all shortcuts in your config file in a nice table format.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config *models.Config
		dir, _ := os.UserHomeDir()
		path := filepath.Join(dir, ".config", "slacker", "config.toml")

		config, err := internal.LoadConfig(path)
		if err != nil {
			fmt.Println("Error loading config:", err)
			os.Exit(1)
		}
		show(config.Shortcuts)
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

const (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")
)

func show(shortcuts []models.Shortcut) {

	var (
		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Left)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle = lipgloss.NewStyle().Padding(0, 1)
		// OddRowStyle is the lipgloss style used for odd-numbered table rows.
		OddRowStyle = CellStyle.Copy().Foreground(gray)
		// EvenRowStyle is the lipgloss style used for even-numbered table rows.
		EvenRowStyle = CellStyle.Copy().Foreground(lightGray)
		// BorderStyle is the lipgloss style used for the table border.
		BorderStyle = lipgloss.NewStyle().Foreground(purple)
	)
	emojiColumnWidth := 15

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			style := CellStyle
			if row == 0 {
				style = HeaderStyle
			} else if row%2 == 0 {
				style = EvenRowStyle
			} else {
				style = OddRowStyle
			}
			// Apply minimum width to the emoji column
			if col == 2 { // Assuming emoji is the third column (0-indexed)
				style = style.Width(emojiColumnWidth).Align(lipgloss.Center)
			}

			return style
		}).
		Headers("Shortcut", "Text", "Emoji", "Duration").
		Rows()

	// Add rows
	for _, sc := range shortcuts {
		sc := sc
		// fmt.Println(sc)

		symbol := emoji.Parse(sc.Emoji)
		t.Row(sc.Name, sc.Text, symbol, sc.Duration)
	}

	fmt.Println(t)
}
