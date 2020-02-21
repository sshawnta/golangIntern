package frame

import (
	`testing`

	`github.com/sshawnta/golangIntern/visitor/pkg/car`
	`github.com/sshawnta/golangIntern/visitor/pkg/model`
	`github.com/sshawnta/golangIntern/visitor/pkg/plane`
	`github.com/stretchr/testify/assert`
)

const (
	correctResultCar     = 1200
	correctresultPlane   = 950.0
	correctResultWrong   = 0
	correctResultWrong64 = 0.0
	wrongPlane           = "qwe"
	wrongCar             = "dasd"
)

//Simple tests for computer
func TestForVisitor(t *testing.T) {
	t.Run("Test for Visitor", func(t *testing.T) {
		newVisitor := NewVisitor(model.AddPriceCar, model.SalePlane64)
		carCheck := car.NewCar(model.OpelCar)
		planeCheck := plane.NewPlane(model.AirPlane)
		assert.Equal(t, correctResultCar, newVisitor.VisitCar(carCheck))
		assert.Equal(t, correctresultPlane, newVisitor.VisitPlane(planeCheck))

		//Wrong
		carCheck = car.NewCar(wrongCar)
		planeCheck = plane.NewPlane(wrongPlane)
		assert.Equal(t, correctResultWrong, newVisitor.VisitCar(carCheck))
		assert.Equal(t, correctResultWrong64, newVisitor.VisitPlane(planeCheck))

	})
}
