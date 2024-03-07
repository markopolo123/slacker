package internal

import (
	"os"

	"github.com/slack-go/slack"
)

var api *slack.Client
var userID string

func SetupAPI() (*slack.Client, string) {
	token := os.Getenv("SLACK_TOKEN")
	userID = os.Getenv("SLACK_USER_ID")
	api = slack.New(token)
	return api, userID
}
