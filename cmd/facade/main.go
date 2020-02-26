package main

import (
	"github.com/sshawnta/golangIntern/pkg/computer"
	"github.com/sshawnta/golangIntern/pkg/frame"
	`github.com/sshawnta/golangIntern/pkg/logger`
	"github.com/sshawnta/golangIntern/pkg/model"
	"github.com/sshawnta/golangIntern/pkg/phone"
)

func main() {
	newPhone := phone.NewPhone(model.PhoneLock, model.PhonePassword)
	newComputer := computer.NewComputer(model.PowerOfComputer, model.UserComp, model.PassComp, model.CorrectCompUser, model.CorrectCompPass)
	newFacade := frame.NewFacade(newPhone, newComputer)

	err := newFacade.CallPhone(model.NumberToCallOrMassage)
	if err != nil {
		logger.LogErrors(err)
	}
	err = newFacade.MessagePhone(model.NumberToCallOrMassage, model.Message)
	if err != nil {
		logger.LogErrors(err)
	}
	err = newFacade.MessageComp(model.NumberToCallOrMassage, model.Message)
	if err != nil {
		logger.LogErrors(err)
	}
}
