package session

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	errUserAlreadyExists error = errors.New("user already exists")
)

func (m Model) Get() ([]SessionModel, error) {
	var sessions []SessionModel
	ctx := context.TODO()
	filter := bson.M{}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

func (m Model) GetById(id primitive.ObjectID) (*SessionModel, error) {
	var session SessionModel
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&session)

	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (m Model) GetByUserId(userId primitive.ObjectID) ([]SessionModel, error) {
	var sessions []SessionModel
	ctx := context.TODO()
	filter := bson.M{"userId": userId}
	cursor, err := m.Collection.Find(ctx, filter)

	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

func (m Model) Add(add *AddSessionModel) (*SessionModel, error) {
	userExists, err := m.GetByUserId(add.UserId)

	if err != nil {
		return nil, err
	}
	for _, user := range userExists {
		if user.ProviderName == add.ProviderName {
			log.Print("user already exists with provider")
			return nil, errUserAlreadyExists
		}
	}
	ctx := context.TODO()

	newSession := SessionModel{
		Id:           primitive.NewObjectID(),
		UserId:       add.UserId,
		ProviderName: add.ProviderName,
		AccessToken:  add.AccessToken,
		RefreshToken: add.RefreshToken,
		Expiry:       add.Expiry,
		IdToken:      add.IdToken,
	}
	_, err = m.Collection.InsertOne(ctx, newSession)

	if err != nil {
		return nil, err
	}
	return &newSession, nil
}

func (m Model) UpdateById(id primitive.ObjectID, update *UpdateSessionModel) (*SessionModel, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	updateData := bson.M{}
	updateBytes, err := bson.Marshal(update)

	if err != nil {
		return nil, err
	}

	bson.Unmarshal(updateBytes, updateData)
	result := m.Collection.FindOneAndUpdate(ctx, filter, updateData)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var updatedSession SessionModel
	if err := result.Decode(&updatedSession); err != nil {
		return nil, err
	}
	return &updatedSession, nil
}

func (m Model) UpdateByUserId(userId primitive.ObjectID, providerName string, update *UpdateSessionModel) (*SessionModel, error) {
	ctx := context.TODO()
	filter := bson.M{"userId": userId, "providerName": providerName}
	updateData := bson.M{}
	updateBytes, err := bson.Marshal(update)

	if err != nil {
		return nil, err
	}

	bson.Unmarshal(updateBytes, updateData)
	result := m.Collection.FindOneAndUpdate(ctx, filter, updateData)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var updatedSession SessionModel
	if err := result.Decode(&updatedSession); err != nil {
		return nil, err
	}
	return &updatedSession, nil
}

func (m Model) DeleteById(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	result, err := m.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (m Model) DeleteByUserId(userId primitive.ObjectID, providerName string) error {
	ctx := context.TODO()
	filter := bson.M{"userId": userId, "providerName": providerName}
	result, err := m.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
