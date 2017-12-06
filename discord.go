package discord

import (
	"github.com/spf13/viper"

	"github.com/bwmarrin/discordgo"
	"github.com/scott-wilson/dosbot"
)

func DiscordConnector(toActions chan<- dosbot.Event, toChannel <-chan dosbot.Event) func() error {
	// Register
	discord, err := discordgo.New("Bot " + viper.GetString("DISCORD_TOKEN"))

	if err != nil {
		panic(err)
	}

	user, err := discord.User("@me")

	if err != nil {
		panic(err)
	}

	bot := newBot(user.Username, user.ID, discord)

	// Input from Discord
	discord.AddHandler(handleMessageCreate(bot, toActions))

	// Output to Discord
	go func() {
		for event := range toChannel {
			err := event.Error()

			if err != nil {
				panic(err)
			}
		}
	}()

	// Open connection to Discord
	if err := discord.Open(); err != nil {
		panic(err)
	}

	// Send close function for cleanup.
	return discord.Close
}

func handleMessageCreate(bot dosbot.Bot, toActions chan<- dosbot.Event) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, message *discordgo.MessageCreate) {
		if session.State.User.ID == message.Author.ID {
			return
		}

		channel, err := session.Channel(message.ChannelID)

		if err != nil {
			panic(err)
		}

		u := newUser(message.Author.Username, message.Author.ID)
		r := newRoom(channel.Name, channel.ID)

		dosbot.EmitMessageActions(message.Content, u, r, bot, toActions)
	}
}

func init() {
	viper.SetEnvPrefix("DOSBOT")
	viper.BindEnv("DISCORD_TOKEN")
}
