package computer

import (
	"github.com/stretchr/testify/mock"
)

//Mock of phone
type Mock struct {
	mock.Mock
}

//Signature of SendMail computer method
func (m *Mock) SendMail(text string) string {
	args := m.Called(text)
	return args.Get(0).(string)
}
