package thoughtfactory

import (
	"github.com/nlopes/slack"
	"teamscorebot/thoughtfactory/types"
)

var thoughts = []Thoughtifier{}

type Thoughtifier interface {
	Understood(*slack.MessageEvent) bool
	Process()
}

func RegisterThought(t Thoughtifier) {
	thoughts = append(thoughts, t)
}

func ThinkAbout(stimulus *slack.MessageEvent) (Thoughtifier) {
	for _, thought := range thoughts {
		if thought.Understood(stimulus) {
			return thought
		}
	}

	return &types.EmptyThought{}
}