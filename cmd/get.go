/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/enescakir/emoji"
	internal "github.com/markopolo123/slacker/internal"
	models "github.com/markopolo123/slacker/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your current Slack status",
	Long:  `Get your current Slack status`,
	Run: func(cmd *cobra.Command, args []string) {
		spin()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func spin() {
	// Function to be called by the spinner
	getStatus := func() {
		getstatus()
	}
	// Create and run the spinner
	err := spinner.New().
		Title("Fetching your status...").
		Action(getStatus).
		Run()
	if err != nil {
		log.Fatal("Failed to update status: ", err)
	}
}

func getstatus() {
	api, userID := internal.SetupAPI()
	user, err := api.GetUserInfo(userID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	var status models.Status

	status.Text = user.Profile.StatusText
	status.Emoji = user.Profile.StatusEmoji
	status.Expiration = user.Profile.StatusExpiration
	// convert int to string
	expiration := internal.CalcTime(status.Expiration)
	symbol := emoji.Parse(status.Emoji)
	wave := emoji.WavingHand
	string := fmt.Sprintf("%s Hi %s!     \n\nYour Slack status:\n\nStatus: %s\nEmoji: %s\nExpiration: %s", wave, user.RealName, status.Text, symbol, expiration)
	var style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63"))
	fmt.Println(style.Render(string))

}
