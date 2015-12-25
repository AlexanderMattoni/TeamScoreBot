package score
import (
	"time"
	"github.com/nlopes/slack"
	//"teamscorebot/teams"
)

type PointsEvent struct {
	Type string			`json:"type"`
	Date time.Time		`json:"date"`
	Reason string		`json:"reason"`
	Giver slack.User	`json:"giver"`
}

type Points struct {
	Total int               `json:"total"`
	Events []PointsEvent    `json:"events"`
}
//
//type PointSaver interface {
//	SavePoints(p *PointsEvent, t *teams.Team)
//}
//
//func (p *PointsEvent) Save(s PointSaver, t *teams.Team) {
//	s.SavePoints(p, t)
//}