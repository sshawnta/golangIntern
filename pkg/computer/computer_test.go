package computer

import (
	`fmt`
	"testing"

	`github.com/sshawnta/golangIntern/pkg/model`
	"github.com/stretchr/testify/assert"
)

const (
	wrongPass = "12"
	wrongUser = "us"
)

//Simple tests for computer
func TestForPhone(t *testing.T) {

	t.Run("correct testing ForComputer", func(t *testing.T) {
		correctComputer := NewComputer(model.PowerOfComputer, model.UserComp, model.PassComp, model.CorrectCompUser, model.CorrectCompPass)
		assert.Equal(t, nil, correctComputer.SendMail(model.NumberToCallOrMassage, model.Message))

	})
	t.Run("wrong testing for Computer", func(t *testing.T) {

		wrongPassComputer := NewComputer(model.PowerOfComputer, model.UserComp, wrongPass,model.CorrectCompUser, model.CorrectCompUser)
		err := fmt.Errorf(model.CompIncorrectLogOrPass)
		assert.Equal(t, err, wrongPassComputer.SendMail(model.NumberToCallOrMassage, model.Message))

		wrongUserComputer := NewComputer(model.PowerOfComputer, wrongUser, model.PassComp,model.CorrectCompUser, model.CorrectCompUser)
		assert.Equal(t, err, wrongUserComputer.SendMail(model.NumberToCallOrMassage, model.Message))
	})

}
