package discord

import (
	"fmt"
	"regexp"
)

type bot struct {
	name                 string
	id                   string
	directedMessageRegex *regexp.Regexp
}

func newBot(name string, id string) bot {
	botIDToken := fmt.Sprintf("<@%s>", id)
	regex := regexp.MustCompile(fmt.Sprintf(`\s*^(?:%s|(?i)%s)\s+(.+)$`, botIDToken, name))
	return bot{name: name, id: id, directedMessageRegex: regex}
}

func (b bot) Name() string {
	return b.name
}

func (b bot) ID() interface{} {
	return b.id
}

func (b bot) DirectedMessageRegex() *regexp.Regexp {
	return b.directedMessageRegex
}
