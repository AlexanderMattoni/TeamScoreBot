package teams
import "github.com/nlopes/slack"

type Member struct {
	*slack.User
	Team Team
}
