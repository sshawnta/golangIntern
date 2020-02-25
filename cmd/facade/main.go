package main

import (
	"github.com/sshawnta/golangIntern/pkg/computer"
	"github.com/sshawnta/golangIntern/pkg/frame"
	`github.com/sshawnta/golangIntern/pkg/logger`
	"github.com/sshawnta/golangIntern/pkg/model"
	"github.com/sshawnta/golangIntern/pkg/phone"
)

func main() {
	newPhone := phone.NewPhone(model.PhoneLock, model.NumberToCallOrMassage, model.PhonePassword)
	newComputer := computer.NewComputer(model.PowerOfComputer,model.NumberToCallOrMassage, model.UserComp, model.PassComp)
	newFacade := frame.NewFacade(newPhone, newComputer)

	err := newFacade.CallPhone()
	if err !=nil{
		logger.LogErrors(err)
	}
	err = newFacade.MessagePhone()
	if err != nil{
		logger.LogErrors(err)
	}
	err = newFacade.MessageComp()
	if err != nil{
		logger.LogErrors(err)
	}
}
