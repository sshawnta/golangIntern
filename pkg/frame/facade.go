package frame

import (
	"github.com/sshawnta/golangIntern/pkg/computer"
	"github.com/sshawnta/golangIntern/pkg/phone"
)


//facade functional
type Facade interface {
	Call()
	SendMessage()
	MessageComp()
}

//make a phone call
func CallPhone() {
	phone := phone.NewPhone("12313", "1234")
	if phone.Unlock() == true {
		phone.Call()
		phone.Lock()
	}
}

//send message from phone
func MessagePhone() {
	phone := phone.NewPhone("12341", "1233")
	if phone.Unlock() == true {
		phone.SendMessage("text")
		phone.Lock()
	}
}

//send message from computer
func MessageComp(){
	comp := computer.NewComputer("1234567","user","pass")
	comp.SendTeleg("text")
}

