package frame

import (
	`github.com/sshawnta/golangIntern/pkg/model`
)

type phone interface {
	Call() string
	SendMessage(text string) string
}

type computer interface {
	SendMail(text string) string
}

// Facade interface Active actions that can be performed on the facade
type Facade interface {
	CallPhone() string
	MessagePhone() string
	MessageComp() string
}
type facade struct {
	phone    phone
	computer computer
}

// CallPhone Make a phone call
func (f *facade) CallPhone() string {
	return f.phone.Call()
}

// MessagePhone Send message from phone
func (f *facade) MessagePhone() string {
	return f.phone.SendMessage(model.Message)
}

// MessageComp Send message from computer
func (f *facade) MessageComp() string {
	return f.computer.SendMail(model.Message)
}

// NewFacade Fabric of facade exemplar
// You must fill in information about your computer and phone number, password, user name and message text
func NewFacade(phone phone, computer computer) Facade {
	return &facade{
		phone:    phone,
		computer: computer,
	}
}
