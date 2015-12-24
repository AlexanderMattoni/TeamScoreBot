package slackapi
import "github.com/nlopes/slack"

var Client *slack.Client

func init() {
	Client = slack.New("xoxb-17294767315-UQPhtIpjeODry9uqssYpB3k0")
	Client.SetDebug(false)
}
