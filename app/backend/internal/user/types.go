package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollectionCtxKey = "UserCollectionCtxKey"

type UserModel struct {
	Id       primitive.ObjectID `json:"id" bson:"_Id"`
	Email    string             `json:"email" bson:"email"`
	Password *string            `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}

type AddUserModel struct {
	Email    string  `json:"email" bson:"email"`
	Password *string `json:"password" bson:"password"`
	Role     string  `json:"role" bson:"role"`
}

type UpdateUserModel struct {
	Password *string `json:"password" bson:"password"`
}

type Service interface {
	Get() ([]UserModel, error)
	GetById(primitive.ObjectID) (*UserModel, error)
	GetByEmail(string) (*UserModel, error)
	Add(*AddUserModel) (*UserModel, error)
	UpdateById(primitive.ObjectID, *UpdateUserModel) (*UserModel, error)
	UpdateByEmail(string, *UpdateUserModel) (*UserModel, error)
	DeleteById(primitive.ObjectID) error
	DeleteByEmail(string) error
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}
