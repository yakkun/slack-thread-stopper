package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	want := &Config{
		Debugging:     false,
		SlackBotToken: "",
		SlackAppToken: "",
	}
	if got.Debugging != want.Debugging {
		t.Errorf("config.Debugging() = %#v; want: %#v", got.Debugging, want.Debugging)
	}
	if got.SlackBotToken != want.SlackBotToken {
		t.Errorf("config.SlackBotToken() = %#v; want: %#v", got.SlackBotToken, want.SlackBotToken)
	}
	if got.SlackAppToken != want.SlackAppToken {
		t.Errorf("config.SlackAppToken() = %#v; want: %#v", got.SlackAppToken, want.SlackAppToken)
	}
}

func TestConfig_Load(t *testing.T) {
	config := New()
	err := config.Load()
	if err != nil {
		t.Errorf("Unexpected error: %#v", err)
	}
}

func TestConfig_loadDebugging(t *testing.T) {
	if value := os.Getenv("DEBUGGING"); value != "" {
		os.Unsetenv("DEBUGGING")
		defer func() { os.Setenv("DEBUGGING", value) }()
	}

	want := true
	os.Setenv("DEBUGGING", "true")
	config := New()
	config.loadDebugging()
	if config.Debugging != want {
		t.Errorf("config.Debugging = %#v; want: %#v", config.Debugging, want)
	}
}

func TestConfig_loadSlackBotToken(t *testing.T) {
	if value := os.Getenv("SLACK_BOT_TOKEN"); value != "" {
		os.Unsetenv("SLACK_BOT_TOKEN")
		defer func() { os.Setenv("SLACK_BOT_TOKEN", value) }()
	}

	want := "some-slack-bot-token"
	os.Setenv("SLACK_BOT_TOKEN", want)
	config := New()
	if err := config.loadSlackBotToken(); err != nil {
		t.Fatalf("%#v", err)
	}
	if config.SlackBotToken != want {
		t.Errorf("config.SlackBotToken = %#v; want: %#v", config.SlackBotToken, want)
	}
}

func TestConfig_loadSlackAppToken(t *testing.T) {
	if value := os.Getenv("SLACK_APP_TOKEN"); value != "" {
		os.Unsetenv("SLACK_APP_TOKEN")
		defer func() { os.Setenv("SLACK_APP_TOKEN", value) }()
	}

	want := "some-slack-app-token"
	os.Setenv("SLACK_APP_TOKEN", want)
	config := New()
	if err := config.loadSlackAppToken(); err != nil {
		t.Fatalf("%#v", err)
	}
	if config.SlackAppToken != want {
		t.Errorf("config.SlackAppToken = %#v; want: %#v", config.SlackAppToken, want)
	}
}
