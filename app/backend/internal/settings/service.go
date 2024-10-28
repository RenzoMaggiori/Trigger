package settings

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m Model) GetById(id primitive.ObjectID) (*SettingsResponseModel, error) {
	var sync SettingsModel
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&sync)

	if err != nil {
		return nil, fmt.Errorf("%s %v", "could not find sync", err)
	}

	var response SettingsResponseModel
	response.ProviderName = sync.ProviderName
	response.Active = sync.Active

	return &response, nil
}

func (m Model) GetByUserId(userId primitive.ObjectID) ([]SettingsResponseModel, error) {
	settings := make([]SettingsModel, 0)
	ctx := context.TODO()
	filter := bson.M{"userId": userId}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("%s %v", "could not find user", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &settings); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	response := make([]SettingsResponseModel, 0)
	for _, setting := range settings {
		response = append(response, SettingsResponseModel{
			ProviderName: setting.ProviderName,
			Active:       setting.Active,
		})
	}

	return response, nil
}

func (m Model) Add(addSettings *AddSettingsModel) error {
	newSeetings := SettingsModel{
		Id:           primitive.NewObjectID(),
		UserId:       addSettings.UserId,
		ProviderName: addSettings.ProviderName,
		// AccessToken:  addSettings.AccessToken,
		Active: addSettings.Active,
	}

	_, err := m.Collection.InsertOne(context.TODO(), newSeetings)
	if err != nil {
		return fmt.Errorf("%s %v", "could not add new setting", err)
	}

	return nil
}
