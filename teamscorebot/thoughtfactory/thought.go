package thoughtfactory

import (
	"regexp"
	"fmt"
)

var thoughts = []Thoughtifier{}

type Thoughtifier interface {
	Understood(input string) bool
	Process() string
}

type Thought struct {
	RegExLiteral string;
	RegEx regexp.Regexp
	Results map[string]string
}

func (e *Thought) Understood(input string) bool {
	re, err := regexp.Compile(e.RegExLiteral)
	if err != nil {
		fmt.Println("Regex is invalid. Could not compile.")
		return false
	}

	e.RegEx = *re;

	n1 := re.SubexpNames()
	r2 := re.FindAllStringSubmatch(input, -1)

	if len(r2) == 0 {
		return false
	}

	e.Results = map[string]string{}
	for i, n := range r2[0] {
		e.Results[n1[i]] = n
	}

	return true
}

func ThinkAbout(stimulus string) (Thoughtifier) {
	for _, thought := range thoughts {
		if thought.Understood(stimulus) {
			return thought
		}
	}

	return &EmptyThought{}
}