package models

// Shortcut represents the command configuration from TOML
type Shortcut struct {
	Name     string
	Text     string
	Emoji    string
	Duration string
}

// Config represents the TOML config structure
type Config struct {
	Shortcuts []Shortcut `toml:"shortcut"`
}
