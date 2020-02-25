package phone

import (
	"github.com/stretchr/testify/mock"
)

//Mock of phone
type Mock struct {
	mock.Mock
}

//Mock sendMessage method of phone
func (m *Mock) SendMessage(text string) (err error) {
	args := m.Called(text)
	return args.Error(0)
}

//Mock call method of phone
func (m *Mock) Call()  (err error) {
	args := m.Called()
	return args.Error(0)
}
