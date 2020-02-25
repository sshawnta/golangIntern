package computer

import (
	"fmt"

	"github.com/sshawnta/golangIntern/pkg/model"
)

// Computer active actions that can be performed on the computer
type Computer interface {
	SendMail(number string, text string) (err error)
}

type computer struct {
	isPower bool
	user    string
	pass    string
}

// SendMail method sending a massage, must to pass a message text
func (c *computer) SendMail(number string, text string) (err error) {
	if !c.isPower {
		c.powerOn()
	}
	err = c.login()
	if err != nil {
		return err
	}
	c.sending(number, text)
	c.powerOff()
	return
}

func (c *computer) login() (err error) {
	if c.user == model.CorrectCompUser && c.pass == model.CorrectCompPass {
		fmt.Println("Login")
		return nil
	}
	err = fmt.Errorf(model.CompIncorrectLogOrPass)
	return err
}

func (c *computer) powerOn() {
	if !c.isPower {
		c.isPower = true
		fmt.Println("PowerOn")
	}
}

func (c *computer) powerOff() {
	if c.isPower {
		c.isPower = false
		fmt.Println("PowerOff")
	}
}

func (c *computer) sending(number string, text string) {
	fmt.Println("Send Message", number, text)
}

// NewComputer creates new computer implementation of the computer interface
func NewComputer(isPower bool, user string, pass string) Computer {
	return &computer{
		isPower: isPower,
		user:    user,
		pass:    pass,
	}
}
