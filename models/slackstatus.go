package models

// Status is a struct that represents a Slack status
type Status struct {
	// The text of the status
	Text string `json:"text"`
	// The emoji of the status
	Emoji      string `json:"emoji"`
	Expiration int    `json:"expiration"`
}
