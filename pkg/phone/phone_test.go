package phone

import (
	"testing"

	`github.com/sshawnta/golangIntern/pkg/model`
	"github.com/stretchr/testify/assert"
)

const (
	result       =  "Success"
)

//Simple Test for phone
func TestForPhone(t *testing.T) {
	t.Run( "testing ForPhone", func(t *testing.T) {
		forPhone := NewPhone(model.NumberToCallOrMassage, model.PhonePassword)
		assert.Equal(t, result, forPhone.Call())
		assert.Equal(t, result, forPhone.SendMessage(model.Message))
	})
}
