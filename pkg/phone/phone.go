package phone

import (
	"fmt"

	"github.com/sshawnta/golangIntern/pkg/model"
)

// Phone Active actions that can be performed on the phone
type Phone interface {
	Call(number string) (err error)
	SendMessage(number string, text string) (err error)
}

type phone struct {
	isLock bool
	pass   string
	maxLen int
	minLen int
}

// SendMessage method sending a massage, must to pass a message text
func (p *phone) SendMessage(number string, text string) (err error) {
	err = p.checkNumber(number)
	if err != nil {
		return
	}
	if p.isLock {
		err = p.unlock()
	}
	if err != nil {
		return
	}
	p.sending(number, text)
	p.lock()
	return
}

// Call method make phone call
func (p *phone) Call(number string) (err error) {
	err = p.checkNumber(number)
	if err != nil {
		return
	}
	if p.isLock {
		err = p.unlock()
	}
	if err != nil {
		return
	}
	fmt.Println("Calling number", number, )
	p.lock()
	return
}

func (p *phone) unlock() (err error) {
	if p.pass == model.CorrectPhonePassword {
		p.isLock = false
		fmt.Println("Phone unlock")
		return
	}
	fmt.Println(model.PhoneIncorrectPass)
	return fmt.Errorf(model.PhoneIncorrectPass)
}

func (p *phone) lock() {
	fmt.Println("phone lock")
	p.isLock = true
}

func (p *phone) checkNumber(number string) (err error) {
	if len(number) > p.maxLen || len(number) < p.minLen {
		return fmt.Errorf(model.PhoneIncorrectNumb)
	}
	return
}

func (p *phone) sending(number string, text string) {
	fmt.Println("Send Message to number", number, text)
}

// NewPhone creates new phone implementation of the Phone interface
func NewPhone(isLock bool, pass string, maxLen int, minLen int) Phone {
	return &phone{
		isLock: isLock,
		pass:   pass,
		maxLen: maxLen,
		minLen: minLen,
	}
}
