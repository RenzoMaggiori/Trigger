package gmail

type Email struct {
}

type Gmail interface {
	Register() error
	Webhook() error
	Send(Email) error
}

type Handler struct {
	Gmail
}

type Model struct {
}
