package session

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		return nil, fmt.Errorf("%w: %v", errBadSessionID, err)
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
		return nil, fmt.Errorf("%w: %v", errBadUserID, err)
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
			return nil, fmt.Errorf("%w: %v", errUserAlreadyExists, add.ProviderName)
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
		return nil, fmt.Errorf("%w: %v", errInsertSession, err)
	}
	return &newSession, nil
}

func (m Model) UpdateById(id primitive.ObjectID, update *UpdateSessionModel) (*SessionModel, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	updateData := bson.M{"$set": update}

	_, err := m.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, err
	}

	var updatedSession SessionModel
	err = m.Collection.FindOne(ctx, filter).Decode(&updatedSession)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errSessionNotFound, err)
	}
	return &updatedSession, nil
}

func (m Model) DeleteById(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	result, err := m.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return fmt.Errorf("%w: %v", errSessionNotFound, err)
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
		return fmt.Errorf("%w: %v", errSessionNotFound, err)
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (m Model) GetByToken(token string) (*SessionModel, error) {
	var session SessionModel
	ctx := context.TODO()
	filter := bson.M{"accessToken": token}
	err := m.Collection.FindOne(ctx, filter).Decode(&session)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errSessionNotFound, err)
	}
	return &session, nil
}
