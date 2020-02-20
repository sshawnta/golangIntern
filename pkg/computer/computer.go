package computer

import (
	"fmt"

	"github.com/sshawnta/golangIntern/pkg/model"
)

//Active actions that can be performed on the computer
type Computer interface {
	SendMail(text string) string
}

type computer struct {
	isPower bool
	number  string
	user    string
	pass    string
}

//method sending a massage, must to pass a message text
func (c *computer) SendMail(text string) string {
	if c.isPower == false {
		c.powerOn()
	}
	err := c.login()
	if err != nil {
		return model.FailComplete
	}
	c.sending(text)
	c.powerOff()
	return model.SuccessComplete
}

func (c *computer) login() error {
	var err error
	if c.user == model.CorrectCompUser && c.pass == model.CorrectCompPass {
		fmt.Println("Login")
		return nil
	}
	fmt.Println(model.CompIncorrectLogOrPass)
	err = fmt.Errorf(model.PhoneIncorrectPass)
	return err
}

func (c *computer) powerOn() {
	if c.isPower == false {
		c.isPower = true
		fmt.Println("PowerOn")
	}
}

func (c *computer) powerOff() {
	if c.isPower == true {
		c.isPower = false
		fmt.Println("PowerOff")
	}
}

func (c *computer) sending(text string) {
	fmt.Println("Send Message", c.number)
}

//constructor of computer exemplar.
func NewComputer(num string, name string, pass string) Computer {
	return &computer{
		isPower: false,
		number:  num,
		user:    name,
		pass:    pass,
	}
}
