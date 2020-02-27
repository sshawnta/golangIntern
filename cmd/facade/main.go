package main

import (
	`log`

	"github.com/pkg/errors"

	"github.com/sshawnta/golangIntern/pkg/computer"
	"github.com/sshawnta/golangIntern/pkg/frame"
	`github.com/sshawnta/golangIntern/pkg/log_decorate`
	"github.com/sshawnta/golangIntern/pkg/model"
	"github.com/sshawnta/golangIntern/pkg/phone"
)

func main() {
	newPhone := phone.NewPhone(model.PhoneLock, model.PhonePassword, model.MaxNumberLenght, model.MinNumberLenght)
	newComputer := computer.NewComputer(model.PowerOfComputer, model.UserComp, model.PassComp, model.CorrectCompUser, model.CorrectCompPass)
	newFacade := frame.NewFacade(newPhone, newComputer)

	callPhone := log_decorate.LogCallPhone(newFacade.CallPhone)
	callPhone(model.NumberToCallOrMassage)

	sendMessage := log_decorate.LogMailCompOrPhone(newFacade.MessagePhone)
	sendMessage(model.NumberToCallOrMassage, model.Message)

	sendMessage = log_decorate.LogMailCompOrPhone(newFacade.MessageComp)
	sendMessage(model.NumberToCallOrMassage, model.Message)


	err := newFacade.CallPhone(model.NumberToCallOrMassage)
	if err != nil {
		err = errors.Wrap(err, "CallPhone with Error")
		log.Println(err)
	}
	err = newFacade.MessagePhone(model.NumberToCallOrMassage, model.Message)
	if err != nil {
		err = errors.Wrap(err, "MessagePhone with Error")
		log.Println(err)
	}
	err = newFacade.MessageComp(model.NumberToCallOrMassage, model.Message)
	if err != nil {
		err = errors.Wrap(err, "MailComp with Error")
		log.Println(err)
	}
}