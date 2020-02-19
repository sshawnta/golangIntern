package computer

import (
	"testing"

	`github.com/sshawnta/golangIntern/pkg/model`
	"github.com/stretchr/testify/assert"
)

const (
	testForComputer = "testing ForComputer"
	result          = "Success"
)

//Simple tests for computer
func TestForPhone(t *testing.T) {
	t.Run(testForComputer, func(t *testing.T) {
		forComputer := NewComputer(model.NumberToCallOrMassage, model.UserComp, model.PassComp)
		assert.Equal(t, result, forComputer.SendMail(model.Message))
	})
}
