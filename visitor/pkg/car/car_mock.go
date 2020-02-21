package car

import (
	`github.com/stretchr/testify/mock`
)

//Mock of car
type Mock struct {
	mock.Mock
}

//Signature of Accept car method
func (m *Mock) Accept(v Visitor) int {
	args := m.Called(v)
	return args.Get(0).(int)
}
