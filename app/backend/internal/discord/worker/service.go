package worker

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) Me(token string) error {
	user, status, err := user.GetUserByAccesstokenRequest(token)
	if err != nil {
		return err
	}

	fmt.Printf("mail: %s, ID: %s\n", user.Email, user.Id)
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf(userEndpoint),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
}

func (m Model) Guilds() error {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return errors.New("error creating Discord session")
	}

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer dg.Close()


	// Fetch the guilds the bot is in.
	guilds, err := dg.Guilds()
	if err != nil {
		log.Fatalf("Error fetching guilds: %v", err)
	}

	// Print out each guild's name and ID.
	// guilds := discord.Guilds()
	fmt.Println("Guilds the bot is in:")
	for _, guild := range guilds {
		fmt.Printf("Name: %s, ID: %s\n", guild.Name, guild.ID)
	}

	return nil
}

func (m Model) GuildChannels(guildID string) error {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return errors.New("error creating Discord session")
	}

	defer discord.Close()

	channels, err := discord.GuildChannels(guildID)

	if err != nil {
		return err
	}

	for _, channel := range channels {
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		fmt.Printf("Channel ID: %s, Name: %s\n", channel.ID, channel.Name)
	}

	return nil

}