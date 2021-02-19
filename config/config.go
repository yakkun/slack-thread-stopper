package config

import (
	"os"
)

type Config struct {
	Debugging         bool
	SlackBotToken     string
	SlackAppToken     string
	ThreadStopMessage string
}

func New() *Config {
	return &Config{
		Debugging:     false,
		SlackBotToken: "",
		SlackAppToken: "",
		ThreadStopMessage: ":no_entry: <@%v> :no_entry:\n" +
			"真・スレッドストッパー。。。(￣ー￣)ﾆﾔﾘｯ\n" +
			"\n" +
			"このチャンネルではスレッドの利用は推奨されていません。必要な場合は、チャンネルを新たに作り、話題を分けられるか検討してください。",
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
	if err := c.loadThreadStopMessage(); err != nil {
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

func (c *Config) loadThreadStopMessage() error {
	v := os.Getenv("THREAD_STOP_MESSAGE")
	if v != "" {
		c.ThreadStopMessage = v
	}
	return nil
}
