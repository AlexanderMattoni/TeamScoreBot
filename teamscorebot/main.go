package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"teamscorebot/thoughtfactory"
	"teamscorebot/slackapi"
	_ "teamscorebot/score"
)

func main() {
	Loop:
	for {
		select {
		case msg := <-slackapi.SlackStream.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
			fmt.Println("I think, therefore...")

			case *slack.ConnectedEvent:
				fmt.Println(fmt.Sprintf("Oh, someone new! I'm talking with %d coders!", ev.ConnectionCount))

			case *slack.MessageEvent:
				thought := thoughtfactory.ThinkAbout(ev);
				thought.Process()

			case *slack.RTMError:
				fmt.Println("Oh no, something went wrong!! Let's see..%s", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Println("Hey, that's the wrong key! I can't get in...I'll just wait for someone to give me the right one.")
				break Loop

			default:
				
			}
		}
	}
}