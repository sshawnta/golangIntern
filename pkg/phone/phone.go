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

//method sending a massage, must to pass a message text
func (p *phone)SendMessage(text string){
	if p.checkNumber() == false{
		return
	}
	if p.isLock{
		p.unlock()
	}
	if !p.isLock{
		p.sending(text)
		p.lock()
	}
}

//method make phone call
func (p *phone)Call(){
	if p.checkNumber() == false{
		return
	}
	if p.isLock{
		p.unlock()
	}
	if !p.isLock {
		fmt.Println("Calling number ", p.number, )
		p.lock()
	}
}

func (p *phone)unlock() bool{
	if p.pass == model.CorrectPhonePassword{
		p.isLock = false
		fmt.Println("Phone unlock")
		return true
	}
	fmt.Println(model.PhoneIncorrectPass)
	return false
}

func (p *phone)lock(){
	fmt.Println("phone lock")
	p.isLock = true
}

func (p *phone)checkNumber() bool{
	if len(p.number) > 11 || len(p.number) < 6{
		fmt.Println(model.PhoneIncorrectNumb)
		return false
	}
	return true
}

func (p *phone)sending(text string){
	fmt.Println("Send Message to number ",p.number)
}

//fabric of phone exemplar
func NewPhone(numb string, upass string) *phone{
	return &phone{
		true,
		numb,
		upass,
	}
}
