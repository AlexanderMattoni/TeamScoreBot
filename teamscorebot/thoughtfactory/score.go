package thoughtfactory
import (
	"fmt"
	"strconv"
)

func init() {
	s := &ScoreThought{};
	s.RegExLiteral = `^(?P<symbol>\+|\-)(?P<amount>[\d]+) points (to|from) (?P<team>[a-zA-Z]+)`
	thoughts = append(thoughts, s)
}

type ScoreThought struct {
	Thought
}

func (s *ScoreThought) Process() string {
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

	return fmt.Sprintf("%s %d points %s %s", sym, amt, tofro, team)
}
