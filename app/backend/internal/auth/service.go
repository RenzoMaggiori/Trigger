package auth

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/hash"
	"trigger.com/trigger/pkg/jwt"
)

var (
	errCredentialsNotFound error = errors.New("could not get credentials from context")
)

func (m Model) Login(ctx context.Context) (string, error) {
	credentials, ok := ctx.Value(CredentialsCtxKey).(CredentialsModel)
	if !ok {
		return "", errCredentialsNotFound
	}

	var user user.UserModel
	filter := bson.M{"email": credentials.Email}
	err := m.DB.Collection("user").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return "", err
	}

	err = hash.VerifyPassword(*user.Password, credentials.Password)
	if err != nil {
		return "", err
	}

	token, err := jwt.Create(
		map[string]string{
			"email": credentials.Email,
		},
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		return "", err
	}
	return token, nil
}
