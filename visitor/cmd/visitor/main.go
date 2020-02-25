package main

import (
	"github.com/sshawnta/golangIntern/visitor/pkg/car"
	"github.com/sshawnta/golangIntern/visitor/pkg/frame"
	"github.com/sshawnta/golangIntern/visitor/pkg/model"
	"github.com/sshawnta/golangIntern/visitor/pkg/plane"
)

func main() {
	newCar := car.NewCar(model.OpelCar)
	newPlane := plane.NewPlane(model.AirPlane)

	newCar.FullInfo()
	newPlane.FullInfo()
	newVisitor := frame.NewVisitor(newCar, newPlane,model.AddPriceCar, model.SalePlane64)
	newCar.Accept(newVisitor)
	newPlane.Accept(newVisitor)
}
