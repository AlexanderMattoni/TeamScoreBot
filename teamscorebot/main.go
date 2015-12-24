package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"teamscorebot/thoughtfactory"
	"teamscorebot/monologue"
)

func main() {
	api := slack.New("xoxb-17294767315-UQPhtIpjeODry9uqssYpB3k0")
	api.SetDebug(false)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
			monologue.Output("I think, therefore...")

			case *slack.ConnectedEvent:
				monologue.Output(fmt.Sprintf("Oh, someone new! I'm talking with %d coders!", ev.ConnectionCount))

			case *slack.MessageEvent:
				thought := thoughtfactory.ThinkAbout(ev.Text);
				rtm.SendMessage(rtm.NewOutgoingMessage(thought.Process(), ev.Channel))

			case *slack.RTMError:
				monologue.Output(fmt.Sprintf("Oh no, something went wrong!! Let's see..%s", ev.Error()))

			case *slack.InvalidAuthEvent:
				monologue.Output("Hey, that's the wrong key! I can't get in...I'll just wait for someone to give me the right one.")
				break Loop

			default:
			}
		}
	}
}