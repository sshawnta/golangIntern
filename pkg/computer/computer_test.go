package computer

import (
	"testing"

	`github.com/sshawnta/golangIntern/pkg/model`
	"github.com/stretchr/testify/assert"
)

const (
	result          = "Success"
)

//Simple tests for computer
func TestForPhone(t *testing.T) {
	t.Run("Testing ForComputer", func(t *testing.T) {
		forComputer := NewComputer(model.NumberToCallOrMassage, model.UserComp, model.PassComp)
		assert.Equal(t, result, forComputer.SendMail(model.Message))
	})
}
