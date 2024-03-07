package internal

import (
	"fmt"

	"github.com/BurntSushi/toml"
	models "github.com/markopolo123/slacker/models"
	"github.com/spf13/cobra"
)

// LoadConfig loads and parses the config.toml file
func LoadConfig(configFile string) (*models.Config, error) {
	var config models.Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// createShortcutCommands creates Cobra commands from shortcuts and adds them to the parent command
func CreateShortcutCommands(parent *cobra.Command, shortcuts []models.Shortcut) {
	for _, sc := range shortcuts {
		sc := sc
		// Create a new command for each shortcut
		cmd := &cobra.Command{
			Use:   sc.Name,
			Short: fmt.Sprintf("A shortcut command for %s", sc.Name),
			Run: func(cmd *cobra.Command, args []string) {
				SetStatus(sc)
			},
		}

		// Add the new command as a subcommand
		parent.AddCommand(cmd)
	}
}

func SetStatus(sc models.Shortcut) {
	fmt.Println("Setting status:", sc.Name)
	api, userID := SetupAPI()
	expiration := CalcUnixTime(sc.Duration)

	// set the slack status
	err := api.SetUserCustomStatusWithUser(userID, sc.Text, sc.Emoji, expiration)
	if err != nil {
		fmt.Printf("Error setting status: %s\n", err)
		return
	}
	fmt.Println("Slack status updated successfully!")
}

func ClearStatus() {
	fmt.Println("Clearing status...")
	api, userID := SetupAPI()

	// set the slack status
	err := api.SetUserCustomStatusWithUser(userID, "", "", 0)
	if err != nil {
		fmt.Printf("Error setting status: %s\n", err)
		return
	}
	fmt.Println("Slack status cleared!")
}
