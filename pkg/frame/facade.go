package frame

import ""

type Facade interface {
	Call()
	SendMessage()
	MessageComp()
}

func CallPhone() {
	phone := NewPhone("12313", "1234")
	if phone.Unlock() == true {
		phone.Call()
		phone.Lock()
	}
}

func MessagePhone(){
	phone := NewPhone("12341", "1233")
	if phone.Unlock() == true {
		phone.SendMessage("text")
		phone.Lock()
	}

	func MessageComp(){
		comp := Newcomputer("1234567","user","pass")
		comp.SendTeleg("text")
	}
}

