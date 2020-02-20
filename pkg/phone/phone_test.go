package phone

import (
	"testing"

	`github.com/sshawnta/golangIntern/pkg/model`
	"github.com/stretchr/testify/assert"
)

const (
	successResult = "Success"
	failResult    = "Fail"
	wrongPass     = "12"
	wrongNumb     = "893448234234234"
)

//Simple Test for phone
func TestForPhone(t *testing.T) {
	t.Run("testing ForPhone", func(t *testing.T) {
		correctPhone := NewPhone(model.PhoneLock, model.NumberToCallOrMassage, model.PhonePassword)
		assert.Equal(t, successResult, correctPhone.Call())
		assert.Equal(t, successResult, correctPhone.SendMessage(model.Message))

		wrongPassPhone := NewPhone(model.PhoneLock, model.NumberToCallOrMassage, wrongPass)
		assert.Equal(t, failResult, wrongPassPhone.SendMessage(model.Message))
		assert.Equal(t, failResult, wrongPassPhone.Call())

		wrongNumbPhone := NewPhone(model.PhoneLock, wrongNumb, model.PhonePassword)
		assert.Equal(t, failResult, wrongNumbPhone.SendMessage(model.Message))
		assert.Equal(t, failResult, wrongNumbPhone.Call())

		wrongPassPhoneUnlock := NewPhone(model.PhoneUnlock, model.NumberToCallOrMassage, wrongPass)
		assert.Equal(t, successResult, wrongPassPhoneUnlock.SendMessage(model.Message))
		assert.Equal(t, failResult, wrongPassPhoneUnlock.Call())
	})
}
