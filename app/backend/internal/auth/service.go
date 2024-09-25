package auth

import (
	"context"
	"errors"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/hash"
	"trigger.com/trigger/pkg/jwt"
)

var (
	errCredentialsNotFound         error = errors.New("could not get credentials from context")
	errAuthorizationHeaderNotFound error = errors.New("could not get authorization header")
	errAuthorizationTypeNone       error = errors.New("could not decypher auth type")
	errTokenNotFound               error = errors.New("could not find token in authorization header")
	errAuthTypeUndefined           error = errors.New("auth type is undefined")
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

func (m Model) GetToken(authorizationHeader string) (string, error) {
	if authorizationHeader == "" {
		return "", errAuthorizationHeaderNotFound
	}

	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		parts := strings.Split(authorizationHeader, " ")
		if len(parts) < 2 || parts[0] != "Bearer" {
			return "", errTokenNotFound
		}
		m.authType = Credentials
		return parts[1], nil
	}
	// TODO: check for oauth token
	return "", errAuthorizationTypeNone
}

func (m Model) VerifyToken(token string) error {
	switch m.authType {
	case Credentials:
		return jwt.Verify(token, os.Getenv("TOKEN_SECRET"))
	case OAuth:
		// TODO: verify oauth2 token
		return errAuthTypeUndefined
	default:
		return errAuthTypeUndefined
	}
}
