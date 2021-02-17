package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"github.com/yakkun/slack-thread-stopper/config"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Unable to load .env file. (Ignore me if you don't handle .env file)")
	}

	conf := config.New()
	if err := conf.Load(); err != nil {
		log.Fatalf("Unable to load config: %#v", err)
	}

	if conf.Debugging == true {
		log.Print("Debug mode is active.")
	}
	if conf.SlackBotToken == "" {
		log.Fatal("SlackBotToken is not set, must set it with Env-vars or .env")
	}
	if conf.SlackAppToken == "" {
		log.Fatal("SlackAppToken is not set, must set it with Env-vars or .env")
	}

	client := slack.New(
		conf.SlackBotToken,
		slack.OptionAppLevelToken(conf.SlackAppToken),
		slack.OptionDebug(conf.Debugging),
	)

	socketMode := socketmode.New(
		client,
		socketmode.OptionDebug(conf.Debugging),
	)

	authTest, err := client.AuthTest()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to authenticate: %v\n", err)
		os.Exit(1)
	}

	selfUserId := authTest.UserID

	go func() {
		for envelope := range socketMode.Events {
			if envelope.Type != socketmode.EventTypeEventsAPI {
				continue
			}

			socketMode.Ack(*envelope.Request)
			eventPayload, _ := envelope.Data.(slackevents.EventsAPIEvent)
			if eventPayload.Type != slackevents.CallbackEvent {
				continue
			}

			switch event := eventPayload.InnerEvent.Data.(type) {
			case *slackevents.MessageEvent:
				if event.User == selfUserId {
					continue
				}
				if event.ThreadTimeStamp == "" {
					continue
				}
				_, _, err := client.PostMessage(
					event.Channel,
					slack.MsgOptionText(
						fmt.Sprintf(":no_entry: <@%v> :no_entry:\n真・スレッドストッパー。。。(￣ー￣)ﾆﾔﾘｯ\n\nこのチャンネルではスレッドの利用は推奨されていません。必要な場合は、チャンネルを新たに作り、話題を分けられるか検討してください。", event.User),
						false,
					),
					slack.MsgOptionTS(event.ThreadTimeStamp),
				)
				if err != nil {
					log.Printf("Unable to reply: %v", err)
				}
			}
		}
	}()

	socketMode.Run()
}
