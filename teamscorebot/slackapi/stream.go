package slackapi
import "github.com/nlopes/slack"

var SlackStream slack.RTM

func init() {
	SlackStream = *Client.NewRTM()
	go SlackStream.ManageConnection()
}
