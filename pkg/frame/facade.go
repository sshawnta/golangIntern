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

//make a phone call
func (f *facade)CallPhone() {
		phone.Call(*f.phone)
}

//send message from phone
func (f *facade)MessagePhone() {
		phone.SendMessage(*f.phone, model.Message)
}

//send message from computer
func (f *facade)MessageComp(){
	computer.SendMail(*f.computer,model.Message)
}

func NewFacade(phone phone, computer computer) *facade {
	return &facade{
		phone: &phone,
		computer: &computer,
	}
}
