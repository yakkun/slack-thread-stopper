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
		ThreadStopMessage: ":no_entry: <@%v> :no_entry:\n" +
			"真・スレッドストッパー。。。(￣ー￣)ﾆﾔﾘｯ\n" +
			"\n" +
			"このチャンネルではスレッドの利用は推奨されていません。必要な場合は、チャンネルを新たに作り、話題を分けられるか検討してください。",
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
	if got.ThreadStopMessage != want.ThreadStopMessage {
		t.Errorf("config.ThreadStopMessage() = %#v; want: %#v", got.ThreadStopMessage, want.ThreadStopMessage)
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

func TestConfig_loadThreadStopMessage(t *testing.T) {
	if value := os.Getenv("THREAD_STOP_MESSAGE"); value != "" {
		os.Unsetenv("THREAD_STOP_MESSAGE")
		defer func() { os.Setenv("THREAD_STOP_MESSAGE", value) }()
	}

	want := "some\nspecial\nmessage"
	os.Setenv("THREAD_STOP_MESSAGE", want)
	config := New()
	if err := config.loadThreadStopMessage(); err != nil {
		t.Fatalf("%#v", err)
	}
	if config.ThreadStopMessage != want {
		t.Errorf("config.ThreadStopMessage = %#v; want: %#v", config.ThreadStopMessage, want)
	}
}
