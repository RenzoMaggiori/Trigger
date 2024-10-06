package action

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m Model) Get() ([]ActionModel, error) {
	var sessions []ActionModel
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

func (m Model) GetById(id primitive.ObjectID) (*ActionModel, error) {
	var session ActionModel
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&session)

	if err != nil {
		return nil, errActionNotFound
	}
	return &session, nil
}

func (m Model) GetByProvider(provider string) ([]ActionModel, error) {
	var sessions []ActionModel
	ctx := context.TODO()
	filter := bson.M{"provider": provider}
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

func (m Model) GetByAction(action string) (*ActionModel, error) {
	var session ActionModel
	ctx := context.TODO()
	filter := bson.M{"action": action}
	err := m.Collection.FindOne(ctx, filter).Decode(&session)

	if err != nil {
		return nil, errActionNotFound
	}
	return &session, nil
}

func (m Model) Add(add *AddActionModel) (*ActionModel, error) {
	ctx := context.TODO()

	if add.Type != "trigger" && add.Type != "reaction" {
		return nil, errBadActionType
	}
	newAction := ActionModel{
		Id:       primitive.NewObjectID(),
		Input:    add.Input,
		Output:   add.Output,
		Provider: add.Provider,
		Type:     add.Type,
		Action:   add.Action,
	}
	_, err := m.Collection.InsertOne(ctx, newAction)

	if err != nil {
		return nil, err
	}
	return &newAction, nil
}
