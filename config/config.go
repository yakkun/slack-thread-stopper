package config

import (
	"os"
)

type Config struct {
	Debugging     bool
	SlackBotToken string
	SlackAppToken string
}

func New() *Config {
	return &Config{
		Debugging:     false,
		SlackBotToken: "",
		SlackAppToken: "",
	}
}

func (c *Config) Load() error {
	c.loadDebugging()
	if err := c.loadSlackBotToken(); err != nil {
		return err
	}
	if err := c.loadSlackAppToken(); err != nil {
		return err
	}
	return nil
}

func (c *Config) loadDebugging() {
	v := os.Getenv("DEBUGGING")
	if v == "true" {
		c.Debugging = true
	}
}

func (c *Config) loadSlackBotToken() error {
	c.SlackBotToken = os.Getenv("SLACK_BOT_TOKEN")
	return nil
}

func (c *Config) loadSlackAppToken() error {
	c.SlackAppToken = os.Getenv("SLACK_APP_TOKEN")
	return nil
}
