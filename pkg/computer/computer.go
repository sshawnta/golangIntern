package computer

import (
	"fmt"
	"github.com/sshawnta/golangIntern/pkg/model"
)

//Active actions that can be performed on the computer
type Computer interface {
	SendMail(text string)
}

type computer struct {
	isPower bool
	number  string
	user    string
	pass    string
}

func NewComputer(num string, name string, pass string) *computer {
	return &computer{
		isPower: false,
		number:  num,
		user:    name,
		pass:    pass,
	}
}

func (c *computer) SendMail(text string) {
	if c.isPower == false {
		c.powerOn()
	}
	browser()
	if c.login() == true {
		c.sending(text)
	}
	c.powerOff()
}

func (c *computer) login() bool {
	if c.user == model.CorrectCompUser && c.pass == model.CorrectCompPass {
		fmt.Println("Login")
		return true
	} else {
		fmt.Println("Incorrect login or password")
	}
	return false
}

func browser() {
	fmt.Println("Browser is open")
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
	//todo with text
	fmt.Println("Messange send", c.number)
}
