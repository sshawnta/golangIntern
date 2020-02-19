package phone

import (
	"github.com/stretchr/testify/mock"
)

//Mock of phone
type Mock struct {
	mock.Mock
}

//Mock sendMessage method of phone
func (m *Mock) SendMessage(text string) string {
	args := m.Called(text)
	return args.Get(0).(string)
}

//Mock call method of phone
func (m *Mock) Call() string {
	args := m.Called()
	return args.Get(0).(string)
}
