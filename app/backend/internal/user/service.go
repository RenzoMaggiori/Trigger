package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/hash"
)

func (m Model) Get() ([]UserModel, error) {
	var users []UserModel
	ctx := context.TODO()
	filter := bson.M{}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (m Model) GetById(id primitive.ObjectID) (*UserModel, error) {
	var user UserModel
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, errUserNotFound
	}
	return &user, nil
}

func (m Model) GetByEmail(email string) (*UserModel, error) {
	var user UserModel
	ctx := context.TODO()
	filter := bson.M{"email": email}
	err := m.Collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m Model) Add(add *AddUserModel) (*UserModel, error) {
	userExists, err := m.GetByEmail(add.Email)
	if userExists != nil {
		return nil, errUserAlreadyExists
	}

	ctx := context.TODO()
	var hashedPassword string = ""
	if add.Password != nil {
		hashedPassword, err = hash.Password(*add.Password)
		if err != nil {
			return nil, err
		}
	}

	newUser := UserModel{
		Id:       primitive.NewObjectID(),
		Email:    add.Email,
		Password: &hashedPassword,
		Role:     "default",
	}
	_, err = m.Collection.InsertOne(ctx, newUser)

	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (m Model) UpdateById(id primitive.ObjectID, update *UpdateUserModel) (*UserModel, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	updateData := bson.M{"$set": update}

	_, err := m.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, err
	}

	var updatedUser UserModel
	err = m.Collection.FindOne(ctx, filter).Decode(&updatedUser)

	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (m Model) UpdateByEmail(email string, update *UpdateUserModel) (*UserModel, error) {
	ctx := context.TODO()
	filter := bson.M{"email": email}
	updateData := bson.M{"$set": update}

	_, err := m.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, err
	}

	var updatedUser UserModel
	err = m.Collection.FindOne(ctx, filter).Decode(&updatedUser)

	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
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

func (m Model) DeleteByEmail(email string) error {
	ctx := context.TODO()
	filter := bson.M{"email": email}
	result, err := m.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
