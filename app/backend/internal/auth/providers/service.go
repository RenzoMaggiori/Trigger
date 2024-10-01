package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/goth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

var (
	errCredentialsNotFound         error = errors.New("could not get credentials from context")
	errAuthorizationHeaderNotFound error = errors.New("could not get authorization header")
	errAuthorizationTypeNone       error = errors.New("could not decypher auth type")
	errTokenNotFound               error = errors.New("could not find token in authorization header")
	errAuthTypeUndefined           error = errors.New("auth type is undefined")
)

func (m Model) Login(ctx context.Context) (string, error) {
	return "", nil
}

func (m Model) Callback(gothUser goth.User) (string, error) {
	addUser := user.AddUserModel{
		Email:    gothUser.Email,
		Password: nil,
	}

	body, err := json.Marshal(addUser)
	if err != nil {
		return "", err
	}
	log.Println("Registering user")

	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/user/add", os.Getenv("USER_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("invalid status code, received %s\n", res.Status)
		return "", errors.New("unable to create user")
	}
	fmt.Printf("%+v\n", res.Body)
	user, err := decode.Json[user.UserModel](res.Body)
	if err != nil {
		log.Println("Decode Json: ", err)
		return "", err
	}
	newSession := SessionModel{
		Id:           primitive.NewObjectID(),
		UserId:       user.Id,
		ProviderName: &gothUser.Provider,
		AccessToken:  gothUser.AccessToken,
		RefreshToken: &gothUser.RefreshToken,
		Expiry:       gothUser.ExpiresAt,
		IdToken:      &gothUser.IDToken,
	}
	_, err = m.DB.Collection("session").InsertOne(context.TODO(), newSession)
	fmt.Println("newSession", newSession)
	if err != nil {
		return "", err
	}
	return gothUser.AccessToken, nil
}

func (m Model) Logout(ctx context.Context) (string, error) {
	accessToken, ok := ctx.Value(CredentialsCtxKey).(string)

	if !ok {
		return "", errCredentialsNotFound
	}
	filter := bson.M{"accessToken": accessToken}
	_, err := m.DB.Collection("session").DeleteOne(ctx, filter)
	if err != nil {
		log.Println("could not delete session")
		return "", err
	}
	return "", nil
}
