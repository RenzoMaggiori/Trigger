package worker

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) Me(token string) (*Me, error) {
	user, _, err := user.GetUserByAccesstokenRequest(token)
	if err != nil {
		return nil, err
	}

	sync, _, err := sync.GetSyncAccessTokenRequest(token, user.Id.Hex(), "discord")
	log.Println("sync_access_token", sync.AccessToken)
	if err != nil {
		return nil, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			userEndpoint,
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", sync.AccessToken),
			},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDiscordMe, err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %v", errors.ErrDiscordMe, res.StatusCode)
	}

	discord_me, err := decode.Json[DiscordMe](res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDecodeData, err)
	}

	log.Println("id", discord_me.Id)
	log.Println("username", discord_me.Username)
	log.Println("email", discord_me.Email)

	me := Me{
		DiscordId:       discord_me.Id,
		Username: discord_me.Username,
		Email:    discord_me.Email,
	}

	log.Println("me", me)

	return &me, nil
}

func (m Model) GuildChannels(guildID string) ([]Channel, error) {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrCreateDiscordSession, err)
	}

	defer discord.Close()

	channels, err := discord.GuildChannels(guildID)

	if err != nil {
		return nil, err
	}

	var response []Channel
	for _, ch := range channels {
		if ch.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		response = append(response, Channel{
			Id:   ch.ID,
			Name: ch.Name,
		})
	}

	return response, nil
}