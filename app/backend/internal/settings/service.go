package settings

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m Model) GetById(id primitive.ObjectID) (*SettingsModel, error) {
	var sync SettingsModel
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&sync)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &sync, nil
}

func (m Model) GetByUserId(userId primitive.ObjectID) ([]SettingsModel, error) {
	var settings []SettingsModel
	ctx := context.TODO()
	filter := bson.M{"userId": userId}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &settings); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return settings, nil
}

func (m Model) Add(addSettings *AddSettingsModel) (error) {
	newSeetings := SettingsModel{
		Id:           primitive.NewObjectID(),
		UserId:       addSettings.UserId,
		ProviderName: addSettings.ProviderName,
		AccessToken:  addSettings.AccessToken,
		Active:       addSettings.Active,
	}

	_, err := m.Collection.InsertOne(context.TODO(), newSeetings)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
