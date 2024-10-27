package sync

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/settings"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) GrantAccess(w http.ResponseWriter, r *http.Request) error {
	redirectUrl := r.URL.Query().Get("redirect")
	access_token := r.Header.Get("Authorization")

	url := base64.URLEncoding.EncodeToString([]byte(redirectUrl))
	token := base64.URLEncoding.EncodeToString([]byte(access_token))
	state := fmt.Sprintf("%s:%s", url, token)

	values := r.URL.Query()
	values.Set("state", state)

	r.URL.RawQuery = values.Encode()
	gothic.BeginAuthHandler(w, r)
	return nil
}

func (m Model) SyncWith(gothUser goth.User, access_token string) error {
	user, _, err := user.GetUserByAccesstokenRequest(access_token)
	if err != nil {
		return err
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

func (m Model) Callback(gothUser goth.User, access_token string) error {
	user, _, err := user.GetUserByAccesstokenRequest(access_token)
	if err != nil {
		return err
	}

	var sync SyncModel
	err = m.Collection.FindOne(context.TODO(), bson.M{"userId": user.Id}).Decode(&sync)
	if err != nil {
		if sync.ProviderName == &gothUser.Provider {
			return m.SyncWith(gothUser, access_token)
		}
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

	log.Println("SYNC ADDED SUCCESSFULLY")

	addSettings := settings.AddSettingsModel{
		UserId:       user.Id,
		ProviderName: &gothUser.Provider,
		Active:       true,
	}

	body, err := json.Marshal(addSettings)
	if err != nil {
		return errors.New("failed to marshal settings")
	}

	res, err := fetch.Fetch(
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
		return errors.New("status code not OK, couldn't add new setting")
	}

	log.Println("SETTINGS ADDED SUCCESSFULLY")

	return nil
}
