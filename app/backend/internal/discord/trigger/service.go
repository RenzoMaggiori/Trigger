package trigger

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"

	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/middleware"
)

func (m *Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}
	sync, _, err := sync.GetSyncAccessTokenRequest(accessToken, session.UserId.Hex(), "discord")
	if err != nil {
		return err
	}

	userId := session.UserId.Hex()

	channel_id := actionNode.Input["channel_id"]

	actionId := actionNode.ActionId.Hex()

	discord_me, err := m.GetMe(sync.AccessToken)
	if err != nil {
		return err
	}

	existingSession, _ := m.GetSessionByUserId(session.UserId.Hex())
	if existingSession != nil {
		if existingSession.ChannelId == channel_id {
			log.Println("Session already exists for this channel")
			return nil
		}
	}

	newSession := &DiscordSessionModel{
		UserId:    userId,
		ChannelId: channel_id,
		ActionId:  actionId,
		Token:     session.AccessToken,
		DiscordData: discord_me,
	}
	err = m.AddSession(newSession)
	if err != nil {
		log.Printf("Error adding session [%s]: %v", newSession.ChannelId, err)
		return err
	}
    return nil
}

func (m *Model) Webhook(ctx context.Context) error {
	token, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	event, ok := ctx.Value(DiscordEventCtxKey).(ActionBody)
	if !ok {
		return errors.ErrEventCtx
	}

	user, _, err := user.GetUserByAccesstokenRequest(token)
	if err != nil {
		return err
	}

	action, _, err := action.GetByActionNameRequest(token, event.Type)
	if err != nil {
		return err
	}

	switch action.Action {
	case "watch_channel_message":
		data, ok := event.Data.(map[string]interface{})
		if !ok {
			return errors.ErrBadWebhookData
		}
		var msgInfo MsgInfo
		if err := mapstructure.Decode(data, &msgInfo); err != nil {
			return err
		}

		_, err := workspace.ActionCompletedRequest(token, workspace.ActionCompletedModel{
			UserId:   user.Id,
			ActionId: action.Id,
			Output: map[string]string{
                "content": msgInfo.Content,
				"author_id": msgInfo.AuthoId,
				"author_username":  msgInfo.AuthoUsername,
            },
			// NodeId:   actionNode.Id,
		})
		return err
	}
	return nil
}

func (m *Model) Stop(ctx context.Context) error {
    return nil
}

func (m *Model) GetMe(syncDiscordToken string) (*discordgo.User, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			userEndpoint,
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", syncDiscordToken),
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

	discord_me, err := decode.Json[discordgo.User](res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDecodeData, err)
	}

	return &discord_me, nil
}

func (m *Model) AddSession(data *DiscordSessionModel) error {
	ctx := context.TODO()
	newSync := DiscordSessionModel{
		UserId:    data.UserId,
		ChannelId: data.ChannelId,
		ActionId:  data.ActionId,
		Token: 	data.Token,
		DiscordData: data.DiscordData,
	}

	_, err := m.Collection.InsertOne(ctx, newSync)
	if err != nil {
		return errors.ErrAddDiscordSession
	}

	log.Printf("Discord session created for user %s...\n", data.UserId)

	return nil
}

func (m *Model) GetSessionByUserId(userId string) (*DiscordSessionModel, error) {
	var discordSession DiscordSessionModel
	err := m.Collection.FindOne(context.TODO(), bson.M{"user_id": userId}).Decode(&discordSession)
	if err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	return &discordSession, nil
}

func (m *Model) GetAllSessions() ([]DiscordSessionModel, error) {
	ctx := context.TODO()
	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	var discordSessions []DiscordSessionModel
	if err = cursor.All(ctx, &discordSessions); err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	return discordSessions, nil
}

func (m *Model) DeleteSession(userId string) error {
	ctx := context.TODO()
	filter := bson.M{"user_id": userId}

	_, err := m.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.ErrDeleteDiscordSession
	}

	log.Printf("Discord session deleted for user %s...\n", userId)

	return nil
}