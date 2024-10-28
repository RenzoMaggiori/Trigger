package action

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/pkg/errors"
)

func (m Model) Get() ([]ActionModel, error) {
	actions := make([]ActionModel, 0)
	ctx := context.TODO()
	filter := bson.M{}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &actions); err != nil {
		return nil, err
	}
	return actions, nil
}

func (m Model) GetById(id primitive.ObjectID) (*ActionModel, error) {
	var action ActionModel
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&action)

	if err != nil {
		return nil, errors.ErrActionNotFound
	}
	return &action, nil
}

func (m Model) GetByProvider(provider string) ([]ActionModel, error) {
	actions := make([]ActionModel, 0)
	ctx := context.TODO()
	filter := bson.M{"provider": provider}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &actions); err != nil {
		return nil, err
	}
	return actions, nil
}

func (m Model) GetByActionName(action string) (*ActionModel, error) {
	var actions ActionModel
	ctx := context.TODO()
	filter := bson.M{"action": action}
	err := m.Collection.FindOne(ctx, filter).Decode(&actions)

	if err != nil {
		return nil, errors.ErrActionNotFound
	}
	return &actions, nil
}

func (m Model) Add(add *AddActionModel) (*ActionModel, error) {
	ctx := context.TODO()

	if add.Type != "trigger" && add.Type != "reaction" {
		return nil, errors.ErrActionTypeNone
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
