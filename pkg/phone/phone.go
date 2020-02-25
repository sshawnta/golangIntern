package phone

import (
	"fmt"

	"github.com/sshawnta/golangIntern/pkg/model"
)

// Phone Active actions that can be performed on the phone
type Phone interface {
	Call() (err error)
	SendMessage(text string) (err error)
}

type phone struct {
	isLock bool
	number string
	pass   string
}

// SendMessage method sending a massage, must to pass a message text
func (p *phone) SendMessage(text string) (err error) {
	err = p.checkNumber()
	if err != nil {
		return err
	}
	if p.isLock {
		err = p.unlock()
	}
	if err != nil {
		return err
	}
	p.sending(text)
	p.lock()
	return
}

// Call method make phone call
func (p *phone) Call() (err error) {
	err = p.checkNumber()
	if err != nil {
		return err
	}
	if p.isLock {
		err = p.unlock()
	}
	if err != nil {
		return err
	}
	fmt.Println("Calling number ", p.number, )
	p.lock()
	return
}

func (p *phone) unlock() (err error) {
	if p.pass == model.CorrectPhonePassword {
		p.isLock = false
		fmt.Println("Phone unlock")
		return nil
	}
	fmt.Println(model.PhoneIncorrectPass)
	return fmt.Errorf(model.PhoneIncorrectPass)
}

func (p *phone) lock() {
	fmt.Println("phone lock")
	p.isLock = true
}

func (p *phone) checkNumber() (err error) {
	if len(p.number) > model.MaxNumberLenght || len(p.number) < model.MinNumberLenght{
		return fmt.Errorf(model.PhoneIncorrectNumb)
	}
	return
}

func (p *phone) sending(text string) {
	fmt.Println("Send Message to number ", p.number, " ", text)
}

// NewPhone creates new phone implementation of the Phone interface
func NewPhone(isLock bool, number string, pass string) Phone {
	return &phone{
		isLock: isLock,
		number: number,
		pass:   pass,
	}
}
