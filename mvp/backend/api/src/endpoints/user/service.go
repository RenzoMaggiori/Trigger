package user

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m Model) Add(new User) (*primitive.ObjectID, error) {
	userCollection := m.Mongo.Database(os.Getenv("MONGO_DB")).Collection("user")
	result, err := userCollection.InsertOne(context.TODO(), new)
	if err != nil {
		return nil, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("could not get id of user")
	}

	return &id, nil
}

func (m Model) GetByEmail(email string) (*User, error) {
	userCollection := m.Mongo.Database(os.Getenv("MONGO_DB")).Collection("user")
	var user User
	err := userCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
