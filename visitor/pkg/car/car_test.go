package car

import (
	`fmt`
	`testing`

	"github.com/sshawnta/golangIntern/visitor/pkg/model"
	`github.com/stretchr/testify/assert`
)

const (
	correctModel = "Opel"
	correctPrice = 1000
	wrongModel   = "asd"
)

//Simple tests for computer
func TestForCar(t *testing.T) {
	t.Run("Test for Car", func(t *testing.T) {
		correctCar := NewCar(correctModel)
		assert.Equal(t, correctModel, correctCar.FullInfo())

		price, err := correctCar.Price()
		assert.Equal(t, correctPrice, price)
		assert.Equal(t, nil, err)

		wrongCar := NewCar(wrongModel)
		assert.Equal(t, wrongModel, wrongCar.FullInfo())

		price, err = wrongCar.Price()
		err2 := fmt.Errorf(model.NotFoundModel)
		assert.Equal(t, price, 0)
		assert.Equal(t, err, err2)
	})
}
