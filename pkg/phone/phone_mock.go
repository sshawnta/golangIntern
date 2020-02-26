package phone

import (
	"github.com/stretchr/testify/mock"
)

//Mock of phone
type Mock struct {
	mock.Mock
}

//Mock sendMessage method of phone
func (m *Mock) SendMessage(number string, text string) (err error) {
	args := m.Called(number, text)
	return args.Error(0)
}

//Mock call method of phone
func (m *Mock) Call(number string) (err error) {
	args := m.Called(number)
	return args.Error(0)
}
