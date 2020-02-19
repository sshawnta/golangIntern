package main

import (
	"github.com/sshawnta/golangIntern/pkg/computer"
	"github.com/sshawnta/golangIntern/pkg/frame"
	"github.com/sshawnta/golangIntern/pkg/model"
	"github.com/sshawnta/golangIntern/pkg/phone"
)

func main() {
	newPhone := phone.NewPhone(model.NumberToCallOrMassage, model.PhonePassword)
	newComputer := computer.NewComputer(model.NumberToCallOrMassage, model.UserComp, model.PassComp)
	newFacade := frame.NewFacade(newPhone, newComputer)
	newFacade.CallPhone()
	newFacade.MessagePhone()
	newFacade.MessageComp()
}
