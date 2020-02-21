package plane

import (
	`github.com/stretchr/testify/mock`
)

//Mock of plane
type Mock struct {
	mock.Mock
}

//Signature of Accept car method
func (m *Mock) Accept(v Visitor) float64 {
	args := m.Called(v)
	return args.Get(0).(float64)
}
