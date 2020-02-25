package phone

import (
	`fmt`
	"testing"

	`github.com/sshawnta/golangIntern/pkg/model`
	"github.com/stretchr/testify/assert"
)

const (
	wrongPass = "12"
	wrongNumb = "893448234234234"
)

//Simple Test for phone
func TestForPhone(t *testing.T) {
	t.Run("Correct testing ForPhone", func(t *testing.T) {
		correctPhone := NewPhone(model.PhoneLock, model.NumberToCallOrMassage, model.PhonePassword)
		assert.Equal(t, nil, correctPhone.Call())
		assert.Equal(t, nil, correctPhone.SendMessage(model.Message))
	})
	t.Run("Wrong testing ForPhone", func(t *testing.T) {
		wrongNumbPhone := NewPhone(model.PhoneLock, wrongNumb, model.PhonePassword)
		err := fmt.Errorf(model.PhoneIncorrectNumb)
		assert.Equal(t, err, wrongNumbPhone.Call())
		assert.Equal(t, err, wrongNumbPhone.SendMessage(model.Message))

		wrongPassPhone := NewPhone(model.PhoneLock, model.NumberToCallOrMassage, wrongPass)
		err = fmt.Errorf(model.PhoneIncorrectPass)
		assert.Equal(t, err, wrongPassPhone.Call())
		assert.Equal(t, err, wrongPassPhone.SendMessage(model.Message))
	})
}
