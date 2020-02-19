package phone

import (
	"fmt"
	"github.com/sshawnta/golangIntern/pkg/model"
)

//Active actions that can be performed on the phone
type Phone interface {
	Call()
	SendMessage(text string)
}

type phone struct {
	isLock bool
	number string
	pass string
}

func NewPhone(number1 string, upass string) *phone{
	return &phone{
		true,
		number1,
		upass,
	}
}

func (s *phone)Unlock() bool{
	if s.pass == model.CorrectPhonePassword{
		s.isLock = false
		fmt.Println("Phone unlock")
		return true
	}
	fmt.Println("Incorrect password")
	return false
}

func (s *phone)Lock(){
	fmt.Println("phone lock")
	s.isLock = true
}

func (s *phone)Call(){
	if s.checkNumber() > 0{
		return
	}
	if s.isLock{
		s.Unlock()
	}
	fmt.Println("Calling number ", s.number,)
	s.Lock()
}

func (s *phone)SendMessage(text string){
	if s.checkNumber() > 0{
		return
	}
	s.Unlock()
	s.sending(text)
	s.Lock()
}

func (s *phone)checkNumber() int{
	if len(s.number) > 11 || len(s.number) < 6{
		fmt.Println("Incorrect number")
		return 1
	}
	return 0
}

//метод отправки сообщения
func (s *phone)sending(text string){
	fmt.Println("Send Message to number ",s.number)
}
