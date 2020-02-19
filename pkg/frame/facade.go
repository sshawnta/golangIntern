package frame

import "github.com/sshawnta/golangIntern/pkg/model"

type phone interface {
	Call()
	SendMessage(text string)
}

type computer interface {
	SendMail(text string)
}

type facade struct {
	phone *phone
	computer *computer
}

//Make a phone call
func (f *facade)CallPhone() {
		phone.Call(*f.phone)
}

//Send message from phone
func (f *facade)MessagePhone() {
		phone.SendMessage(*f.phone, model.Message)
}

//Send message from computer
func (f *facade)MessageComp(){
	computer.SendMail(*f.computer,model.Message)
}

//Fabric of facade exemplar
//You must fill in information about your computer and phone number, password, user name and message text
func NewFacade(phone phone, computer computer) *facade {
	return &facade{
		phone: &phone,
		computer: &computer,
	}
}
