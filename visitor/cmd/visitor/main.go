package main

import (
	`github.com/sshawnta/visitor/pkg/car`
	`github.com/sshawnta/visitor/pkg/frame`
	`github.com/sshawnta/visitor/pkg/model`
	`github.com/sshawnta/visitor/pkg/plane`
)

func main() {
	newCar := car.NewCar(model.OpelCar)
	newPlane := plane.NewPlane(model.AirPlane)

	newCar.FullInfo()
	newPlane.FullInfo()
	newVisitor := frame.NewVisitor(model.AddPriceCar, model.SalePlane64)
	newCar.Accept(newVisitor)
	newPlane.Accept(newVisitor)
}
