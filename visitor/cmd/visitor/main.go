package main

import (
	`fmt`

	"github.com/sshawnta/golangIntern/visitor/pkg/car"
	`github.com/sshawnta/golangIntern/visitor/pkg/frame`

	//"github.com/sshawnta/golangIntern/visitor/pkg/frame"
	"github.com/sshawnta/golangIntern/visitor/pkg/model"
)

var (
	fullInfoCar = map[string]int{
		model.OpelCar:  model.OpelPrice,
		model.MazdaCar: model.MazdaPrice,
		model.BMWCar:   model.BMWPrice,
	}
)

func main() {

	newCar := car.NewCar(model.OpelCar, fullInfoCar)
	//newPlane := plane.NewPlane(model.AirPlane)
	_, _, err := newCar.FullInfo()
	if err != nil {
		fmt.Println(err)
	}
	//newPlane.FullInfo()
	newVisitor := frame.NewVisitor(newCar, model.AddPriceCar, model.SalePlane64)
	newCar.Accept(newVisitor)
	_, _, err = newCar.FullInfo()
	//newPlane.Accept(newVisitor)
}
