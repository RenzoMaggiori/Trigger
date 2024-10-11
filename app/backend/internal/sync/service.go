package sync

import (
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
		log.Println("failed to fetch user")
		return errors.New("failed to fetch user")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("failed to fetch user 2")
		return errors.New("failed to fetch user")
	}

	user, err := decode.Json[user.UserModel](res.Body)
	if err != nil {
		log.Println("failed to decode user")
		return errors.New("failed to decode user")
	}

	ctx := context.TODO()
	filter := bson.M{"userId": user.Id}
	var sync SyncModel
	err = m.Collection.FindOne(ctx, filter).Decode(&sync)

	if err != nil {
		log.Println("failed to find sync")
		return errors.New("failed to find sync")
	}

	newSync := UpdateSyncModel{
		AccessToken:  &gothUser.AccessToken,
		RefreshToken: &gothUser.RefreshToken,
		Expiry:       &gothUser.ExpiresAt,
		IdToken:      &gothUser.IDToken,
	}

	body, err := json.Marshal(newSync)

	if err != nil {
		return errors.New("failed to marshal sync")
	}

	_, err = m.Collection.UpdateByID(context.TODO(), gothUser.UserID, bson.M{"$set": body})
	if err != nil {
		return errors.New("failed to update sync")
	}

	log.Println("UpdateByID OK")

	return nil
}

func (m Model) Callback(gothUser goth.User) error {
	// TODO: get user propperly
	//* as we need the user email to get the user id and asign it to the sync

	// res, err := fetch.Fetch(
	// 	http.DefaultClient,
	// 	fetch.NewFetchRequest(
	// 		http.MethodGet,
	// 		fmt.Sprintf("%s/api/user/email/%s", os.Getenv("USER_SERVICE_BASE_URL"), gothUser.Email),
	// 		nil,
	// 		map[string]string{
	// 			"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
	// 		},
	// 	),
	// )

	// if err != nil {
	// 	log.Println("failed to fetch user")
	// 	return errors.New("failed to fetch user")
	// }

	// defer res.Body.Close()
	// if res.StatusCode != http.StatusOK {
	// 	log.Println("failed to fetch user")
	// 	return errors.New("failed to fetch user")
	// }

	// user, err := decode.Json[user.UserModel](res.Body)
	// if err != nil {
	// 	log.Println("failed to decode user")
	// 	return errors.New("failed to decode user")
	// }

	newSync := SyncModel{
		Id:           primitive.NewObjectID(),
		UserId:       primitive.NewObjectID(),
		// UserId:       user.Id,
		ProviderName: &gothUser.Provider,
		AccessToken:  gothUser.AccessToken,
		RefreshToken: &gothUser.RefreshToken,
		Expiry:       gothUser.ExpiresAt,
		IdToken:      &gothUser.IDToken,
	}
	log.Println("user id for new sync: ", newSync.UserId)
	
	ctx := context.TODO()

	_, err := m.Collection.InsertOne(ctx, newSync)
	if err != nil {
		return errors.New("failed to insert sync")
	}

	log.Println("InsertOne OK")

	return nil
}
