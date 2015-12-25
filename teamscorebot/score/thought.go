package score
import (
	"fmt"
	"strconv"
	"teamscorebot/slackapi"
	"teamscorebot/thoughtfactory/types"
	"teamscorebot/thoughtfactory"
)

func init() {
	s := &ScoreThought{};
	s.RegExLiteral = `^(?P<symbol>\+|\-)(?P<amount>[\d]+) points (to|from) (?P<team>[a-zA-Z]+)`
	thoughtfactory.RegisterThought(s)
}

type ScoreThought struct {
	types.SentenceThought
}

func (s *ScoreThought) Process() {
/*	user, err := slackapi.Client.GetUserInfo(s.SlackEv.User)
	if err != nil {
		fmt.Println("Issue determining user for score.")
	}

	fmt.Println(user.RealName)*/

	sym := ""
	amt, err :=  strconv.Atoi(s.Results["amount"])
	if err != nil {
		fmt.Println("I dont know what number that is...")
	}
	tofro := ""
	team := s.Results["team"]

	switch (s.Results["symbol"]) {
	case "+":
		sym = "Adding"
		tofro = "to"
	case "-":
		sym = "Taking away"
		tofro = "from"
	}

	fmt.Printf("%s %d points %s %s", sym, amt, tofro, team)

	slackapi.SlackStream.SendMessage(
		slackapi.SlackStream.NewOutgoingMessage(
			fmt.Sprintf("%s %d points %s %s", sym, amt, tofro, team),
			s.SlackEv.Channel,
		),
	)
}
