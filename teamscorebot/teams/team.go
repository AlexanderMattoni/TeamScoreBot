package teams
import "teamscorebot/score"

type Team struct {
	ID string			`json:"id"`
	Name string			`json:"name"`
	Members	[]Member	`json:"members"`
	Points score.Points	`json:"points"`
}

type TeamSaver interface {
	SaveTeam(t *Team)
}

type TeamLoader interface {
	LoadTeam(id string) Team
}

func LoadTeam(id string, l TeamLoader) {
	return l.LoadTeam(id)
}

func (t *Team) Save(s TeamSaver) {
	s.SaveTeam(t)
}

func (t *Team) AddMember(m Member) {
	append(t.Members, m);
}