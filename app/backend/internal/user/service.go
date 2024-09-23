package user

import "go.mongodb.org/mongo-driver/bson/primitive"

func (m Model) Get() ([]UserModel, error) {
	return nil, nil
}

func (m Model) GetById(id primitive.ObjectID) (*UserModel, error) {
	return nil, nil
}

func (m Model) GetByEmail(email string) (*UserModel, error) {
	return nil, nil
}

func (m Model) Add(add *AddUserModel) (*UserModel, error) {
	return nil, nil
}

func (m Model) UpdateById(id primitive.ObjectID, update *UpdateUserModel) (*UserModel, error) {
	return nil, nil
}

func (m Model) UpdateByEmail(email string, update *UpdateUserModel) (*UserModel, error) {
	return nil, nil
}

func (m Model) DeleteById(id primitive.ObjectID) error {
	return nil
}

func (m Model) DeleteByEmail(email string) error {
	return nil
}
