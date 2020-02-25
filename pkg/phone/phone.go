package phone

import (
	"fmt"

	"github.com/sshawnta/golangIntern/pkg/model"
)

// Phone Active actions that can be performed on the phone
type Phone interface {
	Call() string
	SendMessage(text string) string
}

type phone struct {
	isLock bool
	number string
	pass   string
}

// SendMessage method sending a massage, must to pass a message text
func (p *phone) SendMessage(text string) string {
	err := p.checkNumber()
	if err != nil {
		return model.FailComplete
	}
	if p.isLock {
		err = p.unlock()
	}
	if err != nil {
		return model.FailComplete
	}
	p.sending(text)
	p.lock()
	return model.SuccessComplete
}

// Call method make phone call
func (p *phone) Call() string {
	err := p.checkNumber()
	if err != nil {
		fmt.Println(err)
		return model.FailComplete
	}
	if p.isLock {
		err = p.unlock()
	}
	if err != nil {
		return model.FailComplete
	}
	fmt.Println("Calling number ", p.number, )
	p.lock()
	return model.SuccessComplete
}

func (p *phone) unlock() (err error) {
	if p.pass == model.CorrectPhonePassword {
		p.isLock = false
		fmt.Println("Phone unlock")
		return nil
	}
	fmt.Println(model.PhoneIncorrectPass)
	//err =
	return fmt.Errorf(model.PhoneIncorrectPass)
}

func (p *phone) lock() {
	fmt.Println("phone lock")
	p.isLock = true
}

func (p *phone) checkNumber() (err error) {
	if len(p.number) > 11 || len(p.number) < 6 {
		err = fmt.Errorf(model.PhoneIncorrectNumb)
		return err
	}
	return err
}

func (p *phone) sending(text string) {
	fmt.Println("Send Message to number ", p.number, " ", text)
}

// NewPhone fabric of phone exemplar
func NewPhone(lock bool, numb string, upass string) Phone {
	return &phone{
		lock,
		numb,
		upass,
	}
}
