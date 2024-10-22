package sync

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/markbates/goth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/settings"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) SyncWith(gothUser goth.User) error {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/user/email/%s", os.Getenv("USER_SERVICE_BASE_URL"), gothUser.Email),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)

	if err != nil {
		return errors.New("failed to fetch user")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("status code not OK, couldn't fetch user")
	}

	user, err := decode.Json[user.UserModel](res.Body)
	if err != nil {
		return errors.New("failed to decode user")
	}

	ctx := context.TODO()
	filter := bson.M{"userId": user.Id}
	update := UpdateSyncModel{
		AccessToken:  &gothUser.AccessToken,
		RefreshToken: &gothUser.RefreshToken,
		Expiry:       &gothUser.ExpiresAt,
		IdToken:      &gothUser.IDToken,
	}
	updateData := bson.M{"$set": update}
	_, err = m.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return errors.New("failed to find sync")
	}

	var updatedSync SyncModel
	err = m.Collection.FindOne(ctx, filter).Decode(&updatedSync)

	if err != nil {
		return errors.New("sync not found")
	}

	return nil
}

func (m Model) Callback(gothUser goth.User) error {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/user/email/%s", os.Getenv("USER_SERVICE_BASE_URL"), gothUser.Email),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)

	if err != nil {
		return errors.New("failed to fetch user")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("status code not OK, couldn't fetch user")
	}

	user, err := decode.Json[user.UserModel](res.Body)
	if err != nil {
		return errors.New("failed to decode user")
	}

	syncExists := m.Collection.FindOne(context.TODO(), bson.M{"userId": user.Id})
	if syncExists.Err() == nil {
		return m.SyncWith(gothUser)
	}

	newSync := SyncModel{
		Id:           primitive.NewObjectID(),
		UserId:       user.Id,
		ProviderName: &gothUser.Provider,
		AccessToken:  gothUser.AccessToken,
		RefreshToken: &gothUser.RefreshToken,
		Expiry:       gothUser.ExpiresAt,
		IdToken:      &gothUser.IDToken,
	}

	ctx := context.TODO()
	_, err = m.Collection.InsertOne(ctx, newSync)
	if err != nil {
		return errors.New("failed to insert sync")
	}

	addSettings := settings.AddSettingsModel{
		UserId:       user.Id,
		ProviderName: &gothUser.Provider,
		AccessToken:  gothUser.AccessToken,
		Active:       true,
	}

	body, err := json.Marshal(addSettings)
	if err != nil {
		return errors.New("failed to marshal settings")
	}

	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/settings/add", os.Getenv("SETTINGS_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)

	if err != nil {
		return errors.New("failed to add new settings")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return errors.New("status code not OK, couldn't add new settings")
	}

	return nil
}
