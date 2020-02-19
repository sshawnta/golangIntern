package main

import (
	"github.com/sshawnta/golangIntern/pkg/computer"
	"github.com/sshawnta/golangIntern/pkg/frame"
	"github.com/sshawnta/golangIntern/pkg/model"
	"github.com/sshawnta/golangIntern/pkg/phone"
)

func main() {
	ph := phone.NewPhone(model.NumberToCallOrMassage, model.PhonePassword)
	comp := computer.NewComputer(model.NumberToCallOrMassage, model.UserComp, model.PassComp)
	facade := frame.NewFacade(ph,comp)
	facade.CallPhone()
	facade.MessagePhone()
	facade.MessageComp()
}
