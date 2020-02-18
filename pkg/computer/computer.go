package computer

import "fmt"

type computerPrivate interface {
	powerOn()
	powerOff()
	sending(text string)
	browser()
	login() bool
}

type ComputerPublic interface {
	SendTeleg(text string)
}

type computer struct {
	ispower bool
	number string
	user string
	pass string
}

func NewComputer(num string, name string, pass string) *computer{
	return &computer{
		ispower: false,
		number:  num,
		user:    name,
		pass:    pass,
	}
}

func (c *computer)SendTeleg(text string){
	if c.ispower == false{
		c.powerOn()
	}
	browser()
	if c.login() == true {
		c.sending(text)
	}
	c.powerOff()
}

func (c *computer) login() bool{
	if c.user == "user" && c.pass == "pass"{
		fmt.Println("Login")
		return true
	}
	return false
}

func browser(){
	fmt.Println("Browser is open")
}

func (c *computer)powerOn(){
	if c.ispower == false{
		c.ispower = true
		fmt.Print("PowerOn")
	}
}

func (c *computer)powerOff(){
	if c.ispower == true{
		c.ispower = false
		fmt.Println("PowerOff")
	}
}

func (c *computer)sending(text string){
	//todo with text
	fmt.Println("Messange send", c.number)
}
