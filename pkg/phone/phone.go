package phone

import (
	"fmt"
)

type phonePrivate interface {
	sending(text string)
	checkNumber() int
}

//main phone functional
type PhonePublic interface {
	NewStatement(number1 string, upass string) *phone
	Unlock() bool
	Lock()
	Call()
	SendMessage(text string)
}

type phone struct {
	islock bool
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
	if s.pass == "1234"{
		s.islock = false
		fmt.Println("Phone unlock")
		return true
	}
	fmt.Println("Incorrect password")
	return false
}

func (s *phone)Lock(){
	fmt.Println("phone lock")
	s.islock = true
}

func (s *phone)Call(){
	if s.checkNumber() > 0{
		return
	}
	if s.islock{
		s.Unlock()
	}
	fmt.Println("Calling number ", s.number,)
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
	//TODO with text
	fmt.Println("Send Message to number ",s.number)
}
