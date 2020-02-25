package frame

import (
	`testing`

	computer2 `github.com/sshawnta/golangIntern/pkg/computer`
	`github.com/sshawnta/golangIntern/pkg/model`
	phone2 `github.com/sshawnta/golangIntern/pkg/phone`
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
		comparisonMock := new(phone2.Mock)
		comparisonMock.On(methodCallPhone).Return(error(err)).Once()
		comparisonMock.On(methodMessagePhone, model.Message).Return(error(err)).Once()

		computerMock := new(computer2.Mock)
		computerMock.On(methodMessageComp, model.Message).Return(error(err)).Once()
		myFacade := NewFacade(comparisonMock, computerMock)

		res := myFacade.CallPhone()
		assert.Equal(t, err, res)
		res = myFacade.MessagePhone()
		assert.Equal(t, err, res)
		res = myFacade.MessageComp()
		assert.Equal(t, err, res)

	})
}
