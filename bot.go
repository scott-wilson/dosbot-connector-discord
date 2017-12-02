package discord

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	name                 string
	id                   string
	directedMessageRegex *regexp.Regexp
	session              *discordgo.Session
}

func newBot(name string, id string, session *discordgo.Session) Bot {
	botIDToken := fmt.Sprintf("<@%s>", id)
	regex := regexp.MustCompile(fmt.Sprintf(`\s*^(?:%s|(?i)%s)\s+(.+)$`, botIDToken, name))
	return Bot{name: name, id: id, directedMessageRegex: regex, session: session}
}

func (b Bot) Name() string {
	return b.name
}

func (b Bot) ID() interface{} {
	return b.id
}

func (b Bot) DirectedMessageRegex() *regexp.Regexp {
	return b.directedMessageRegex
}

func (b Bot) Session() *discordgo.Session {
	return b.session
}
