package gmail

import "errors"

var _ Gmail = Model{}

func (m Model) Register() error {
	return errors.New("Not implemented")
}

func (m Model) Webhook() error {
	return errors.New("Not implemented")
}

func (m Model) Send(email Email) error {
	return errors.New("Not implemented")
}
