package frame

import (
	`testing`

	`github.com/sshawnta/golangIntern/pkg/computer`
	`github.com/sshawnta/golangIntern/pkg/model`
	`github.com/sshawnta/golangIntern/pkg/phone`
	`github.com/stretchr/testify/assert`
)

const (
	methodCallPhone    = "Call"
	methodMessagePhone = "SendMessage"
	methodMessageComp  = "SendMail"
)

//Facade testing
func TestAllUse(t *testing.T) {
	t.Run("Test facade", func(t *testing.T) {
		err := *new(error)
		comparisonMock := new(phone.Mock)
		comparisonMock.On(methodCallPhone, model.NumberToCallOrMassage).Return(error(err)).Once()
		comparisonMock.On(methodMessagePhone, model.NumberToCallOrMassage, model.Message).Return(error(err)).Once()

		computerMock := new(computer.Mock)
		computerMock.On(methodMessageComp, model.NumberToCallOrMassage, model.Message).Return(error(err)).Once()
		myFacade := NewFacade(comparisonMock, computerMock)

		res := myFacade.CallPhone(model.NumberToCallOrMassage)
		assert.Equal(t, err, res)
		res = myFacade.MessagePhone(model.NumberToCallOrMassage, model.Message)
		assert.Equal(t, err, res)
		res = myFacade.MessageComp(model.NumberToCallOrMassage, model.Message)
		assert.Equal(t, err, res)

	})

}
