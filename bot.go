package discord

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/scott-wilson/dosbot"
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

func (b Bot) SendMessage(room dosbot.Room, message string) error {
	_, err := b.session.ChannelMessageSend(room.ID().(string), message)
	return err
}

func (b Bot) SendDirectMessage(room dosbot.Room, user dosbot.User, message string) error {
	message = fmt.Sprintf("<@%s> %s", user.ID().(string), message)
	_, err := b.session.ChannelMessageSend(room.ID().(string), message)
	return err
}

func (b Bot) SendEmote(room dosbot.Room, message string) error {
	message = fmt.Sprintf("_%s_", message)
	_, err := b.session.ChannelMessageSend(room.ID().(string), message)
	return err
}

func (b Bot) SendPrivateMessage(user dosbot.User, message string) error {
	channel, err := b.session.UserChannelCreate(user.ID().(string))

	if err != nil {
		return err
	}

	_, err = b.session.ChannelMessageSend(channel.ID, message)

	return err
}

func (b Bot) SendPrivateEmote(user dosbot.User, message string) error {
	message = fmt.Sprintf("_%s_", message)
	channel, err := b.session.UserChannelCreate(user.ID().(string))

	if err != nil {
		return err
	}

	_, err = b.session.ChannelMessageSend(channel.ID, message)

	return err
}
