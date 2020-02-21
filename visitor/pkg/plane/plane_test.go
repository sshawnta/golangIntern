package plane

import (
	`fmt`
	`testing`

	"github.com/sshawnta/golangIntern/visitor/pkg/model"
	`github.com/stretchr/testify/assert`
)

const (
	correctModel = "AirBus"
	correctSale  = 5
	correctPrice = 1000.0
	wrongModel   = "asd"
)

//Simple tests for computer
func TestForPlane(t *testing.T) {
	t.Run("Test for Plane", func(t *testing.T) {
		correctPlane := NewPlane(correctModel)
		assert.Equal(t, correctModel, correctPlane.FullInfo())

		price, err := correctPlane.Price()
		assert.Equal(t, correctPrice, price)
		assert.Equal(t, nil, err)

		wrongCar := NewPlane(wrongModel)
		assert.Equal(t, wrongModel, wrongCar.FullInfo())

		price, err = wrongCar.Price()
		err2 := fmt.Errorf(model.NotFoundModel)
		assert.Equal(t, price, 0.0)
		assert.Equal(t, err, err2)
	})
}
