package phone

import (
	`fmt`
	"testing"

	"github.com/stretchr/testify/assert"

	`github.com/sshawnta/golangIntern/pkg/model`
)

const (
	wrongPass = "12"
	wrongNumb = "893448234234234"
)

// Simple Test for phone
func TestForPhone(t *testing.T) {
	t.Run("Correct testing ForPhone", func(t *testing.T) {
		correctPhone := NewPhone(model.PhoneLock, model.PhonePassword, model.MaxNumberLenght, model.MinNumberLenght)
		assert.Equal(t, nil, correctPhone.Call(model.NumberToCallOrMassage))
		assert.Equal(t, nil, correctPhone.SendMessage(model.NumberToCallOrMassage, model.Message))
	})
	t.Run("Wrong testing ForPhone", func(t *testing.T) {
		wrongNumbPhone := NewPhone(model.PhoneLock, model.PhonePassword, model.MaxNumberLenght, model.MinNumberLenght)
		err := fmt.Errorf(model.PhoneIncorrectNumb)
		assert.Equal(t, err, wrongNumbPhone.Call(wrongNumb))
		assert.Equal(t, err, wrongNumbPhone.SendMessage(wrongNumb, model.Message))

		wrongPassPhone := NewPhone(model.PhoneLock, wrongPass, model.MaxNumberLenght, model.MinNumberLenght)
		err = fmt.Errorf(model.PhoneIncorrectPass)
		assert.Equal(t, err, wrongPassPhone.Call(model.NumberToCallOrMassage))
		assert.Equal(t, err, wrongPassPhone.SendMessage(model.NumberToCallOrMassage, model.Message))
	})
}
