package computer

import (
	"testing"

	`github.com/sshawnta/golangIntern/pkg/model`
	"github.com/stretchr/testify/assert"
)

const (
	success   = "Success"
	fail      = "Fail"
	wrongPass = "12"
	wrongUser = "us"
)

//Simple tests for computer
func TestForPhone(t *testing.T) {
	t.Run("Testing ForComputer", func(t *testing.T) {
		correctComputer := NewComputer(model.NumberToCallOrMassage, model.UserComp, model.PassComp)
		assert.Equal(t, success, correctComputer.SendMail(model.Message))

		wrongPassComputer := NewComputer(model.NumberToCallOrMassage, model.UserComp, wrongPass)
		assert.Equal(t, fail, wrongPassComputer.SendMail(model.Message))

		wrongUserComputer := NewComputer(model.NumberToCallOrMassage, wrongUser, model.PassComp)
		assert.Equal(t, fail, wrongUserComputer.SendMail(model.Message))
	})
}
