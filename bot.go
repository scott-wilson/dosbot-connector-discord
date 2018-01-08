package discord

import (
	"fmt"
	"log"
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
	log.Printf("Sending message.\n\tRoom: %#v\n\tMessage: %s\n", room, message)
	_, err := b.session.ChannelMessageSend(room.ID().(string), message)
	return err
}

func (b Bot) SendDirectMessage(room dosbot.Room, user dosbot.User, message string) error {
	log.Printf("Sending direct message.\n\tRoom: %#v\n\tUser: %#v\n\tMessage: %s\n", room, user, message)
	message = fmt.Sprintf("<@%s> %s", user.ID().(string), message)
	_, err := b.session.ChannelMessageSend(room.ID().(string), message)
	return err
}

func (b Bot) SendEmote(room dosbot.Room, message string) error {
	log.Printf("Sending emote.\n\tRoom: %#v\n\tMessage: %s\n", room, message)
	message = fmt.Sprintf("_%s_", message)
	_, err := b.session.ChannelMessageSend(room.ID().(string), message)
	return err
}

func (b Bot) SendPrivateMessage(user dosbot.User, message string) error {
	log.Printf("Sending private message.\n\tUser: %#v\n\tMessage: %s\n", user, message)
	channel, err := b.session.UserChannelCreate(user.ID().(string))

	if err != nil {
		return err
	}

	_, err = b.session.ChannelMessageSend(channel.ID, message)

	return err
}

func (b Bot) SendPrivateEmote(user dosbot.User, message string) error {
	log.Printf("Sending private emote.\n\tUser: %#v\n\tMessage: %s\n", user, message)
	message = fmt.Sprintf("_%s_", message)
	channel, err := b.session.UserChannelCreate(user.ID().(string))

	if err != nil {
		return err
	}

	_, err = b.session.ChannelMessageSend(channel.ID, message)

	return err
}
