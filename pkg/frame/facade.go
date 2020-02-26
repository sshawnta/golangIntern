package frame

type phoneFacade interface {
	Call(number string) (err error)
	SendMessage(number string, text string) (err error)
}

type computerFacade interface {
	SendMail(number string, text string) (err error)
}

// Facade interface Active actions that can be performed on the facade
type Facade interface {
	CallPhone(number string) (err error)
	MessagePhone(number string, message string) (err error)
	MessageComp(number string, message string) (err error)
}

type facade struct {
	phone    phoneFacade
	computer computerFacade
}

// CallPhone Make a phone call
func (f *facade) CallPhone(number string) (err error) {
	return f.phone.Call(number)
}

// MessagePhone Send message from phone
func (f *facade) MessagePhone(number string, message string) (err error) {
	return f.phone.SendMessage(number, message)
}

// MessageComp Send message from computer
func (f *facade) MessageComp(number string, message string) (err error) {
	return f.computer.SendMail(number, message)
}

// NewFacade constructor of facade exemplar
// You must fill in information about your computer and phone number, password, user name and message text
func NewFacade(phone phoneFacade, computer computerFacade) Facade {
	return &facade{
		phone:    phone,
		computer: computer,
	}
}
