package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id       primitive.ObjectID `json:"id" bson:"_Id"`
	Email    string             `json:"email" bson:"email"`
	Password *string            `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}
