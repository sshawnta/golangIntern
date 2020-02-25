package frame

import (
	`github.com/sshawnta/golangIntern/pkg/model`
)

type phone interface {
	Call() (err error)
	SendMessage(text string) (err error)
}

type computer interface {
	SendMail(text string) (err error)
}

// Facade interface Active actions that can be performed on the facade
type Facade interface {
	CallPhone() (err error)
	MessagePhone() (err error)
	MessageComp() (err error)
}

type facade struct {
	phone    phone
	computer computer
}

// CallPhone Make a phone call
func (f *facade) CallPhone() (err error) {
	return f.phone.Call()
}

// MessagePhone Send message from phone
func (f *facade) MessagePhone() (err error) {
	return f.phone.SendMessage(model.Message)
}

// MessageComp Send message from computer
func (f *facade) MessageComp() (err error) {
	return f.computer.SendMail(model.Message)
}

// NewFacade constructor of facade exemplar
// You must fill in information about your computer and phone number, password, user name and message text
func NewFacade(phone phone, computer computer) Facade {
	return &facade{
		phone:    phone,
		computer: computer,
	}
}
