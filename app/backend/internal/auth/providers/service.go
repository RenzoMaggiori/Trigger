package providers

import (
	"context"
	"errors"

	"github.com/markbates/goth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/user"
)

var (
	errCredentialsNotFound         error = errors.New("could not get credentials from context")
	errAuthorizationHeaderNotFound error = errors.New("could not get authorization header")
	errAuthorizationTypeNone       error = errors.New("could not decypher auth type")
	errTokenNotFound               error = errors.New("could not find token in authorization header")
	errAuthTypeUndefined           error = errors.New("auth type is undefined")
)

func (m Model) Login(ctx context.Context) (string, error) {
	return "", nil
}

func (m Model) Callback(gothUser goth.User) (string, error) {
	user := user.AddUserModel{
		Email:    gothUser.Email,
		Password: nil,
	}

	return "", nil
}

func (m Model) Logout(ctx context.Context) (string, error) {

	return "", nil
}

func (m Model) Get() ([]SessionModel, error) {
	return nil, nil
}

func (m Model) GetById(primitive.ObjectID) (*SessionModel, error) {
	return nil, nil
}

func (m Model) GetByUserId(primitive.ObjectID) (*SessionModel, error) {
	return nil, nil
}

func (m Model) Add(*AddSessionModel) (*SessionModel, error) {
	return nil, nil
}

func (m Model) UpdateById(primitive.ObjectID, *UpdateSessionModel) (*SessionModel, error) {
	return nil, nil
}

func (m Model) UpdateByUserId(primitive.ObjectID, *UpdateSessionModel) (*SessionModel, error) {
	return nil, nil
}

func (m Model) DeleteById(primitive.ObjectID) error {
	return nil
}

func (m Model) DeleteByUserId(primitive.ObjectID) error {
	return nil
}
