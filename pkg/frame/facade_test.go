package frame

import (
	`testing`

	computer2 `github.com/sshawnta/golangIntern/pkg/computer`
	`github.com/sshawnta/golangIntern/pkg/model`
	phone2 `github.com/sshawnta/golangIntern/pkg/phone`
	`github.com/stretchr/testify/assert`
)

const (
	testName           = "Test facade"
	methodCallPhone    = "Call"
	methodMessagePhone = "SendMessage"
	methodMessageComp  = "SendMail"

	Result = "Success"
)

//Facade testing
func TestAllUse(t *testing.T) {
	t.Run(testName, func(t *testing.T) {
		comparisonMock := new(phone2.Mock)
		comparisonMock.On(methodCallPhone).Return(string(Result)).Once()
		comparisonMock.On(methodMessagePhone, model.Message).Return(string(Result)).Once()

		computerMock := new(computer2.Mock)
		computerMock.On(methodMessageComp, model.Message).Return(string(Result)).Once()

		myFacade := NewFacade(comparisonMock, computerMock)
		res := myFacade.CallPhone()
		assert.Equal(t, Result, res)
		res = myFacade.MessagePhone()
		assert.Equal(t, Result, res)
		res = myFacade.MessageComp()
		assert.Equal(t, Result, res)
	})
}
