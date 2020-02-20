package phone

import (
	"fmt"
	`log`

	"github.com/sshawnta/golangIntern/pkg/model"
)

//Active actions that can be performed on the phone
type Phone interface {
	Call() string
	SendMessage(text string) string
}

type phone struct {
	isLock bool
	number string
	pass   string
}

//method sending a massage, must to pass a message text
func (p *phone) SendMessage(text string) string {
	err := p.checkNumber()

	if err != nil {
		log.Fatal(err)
	}
	if p.isLock {
		p.unlock()
	}
	if !p.isLock {
		p.sending(text)
		p.lock()
		return model.SuccessComplete

	}
	return model.FailComplete
}

//method make phone call
func (p *phone) Call() string {
	err := p.checkNumber()
	if err != nil {
		fmt.Println(err)
		return model.FailComplete
	}
	if p.isLock {
		p.unlock()
	}
	if !p.isLock {
		fmt.Println("Calling number ", p.number, )
		p.lock()
		return model.SuccessComplete
	}
	return model.FailComplete
}

func (p *phone) unlock() bool {
	if p.pass == model.CorrectPhonePassword {
		p.isLock = false
		fmt.Println("Phone unlock")
		return true
	}
	fmt.Println(model.PhoneIncorrectPass)
	return false
}

func (p *phone) lock() {
	fmt.Println("phone lock")
	p.isLock = true
}

func (p *phone) checkNumber() error {
	var err error
	if len(p.number) > 11 || len(p.number) < 6 {
		err = fmt.Errorf(model.PhoneIncorrectNumb)
		return err
	}
	return err
}

func (p *phone) sending(text string) {
	fmt.Println("Send Message to number ", p.number)
}

//fabric of phone exemplar
func NewPhone(numb string, upass string) Phone {
	return &phone{
		true,
		numb,
		upass,
	}
}
