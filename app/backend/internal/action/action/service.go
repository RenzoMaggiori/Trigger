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

func (m Model) GetByActionType(actionType string) ([]ActionModel, error) {
	var sessions []ActionModel
	ctx := context.TODO()
	filter := bson.M{"action_type": actionType}
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

func (m Model) Add(add *AddActionModel) (*ActionModel, error) {
	ctx := context.TODO()

	newAction := ActionModel{
		Id:         primitive.NewObjectID(),
		Input:      add.Input,
		Output:     add.Output,
		ActionType: add.ActionType,
	}
	_, err := m.Collection.InsertOne(ctx, newAction)

	if err != nil {
		return nil, err
	}
	return &newAction, nil
}
