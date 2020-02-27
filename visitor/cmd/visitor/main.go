package main

import (
	`fmt`
	`log`

	"github.com/sshawnta/golangIntern/visitor/pkg/car"
	`github.com/sshawnta/golangIntern/visitor/pkg/frame`
	`github.com/sshawnta/golangIntern/visitor/pkg/plane`

	// "github.com/sshawnta/golangIntern/visitor/pkg/frame"
	"github.com/sshawnta/golangIntern/visitor/pkg/model"
)

var (
	fullInfoCar = map[string]int{
		model.OpelCar:  model.OpelPrice,
		model.MazdaCar: model.MazdaPrice,
		model.BMWCar:   model.BMWPrice,
	}

	fullInfoPlane = map[string]float64{
		model.BoeingPlane: model.BoeingPrice,
		model.AirBusPlane:    model.AirBusPrice,
		model.IlPlane:     model.IlPrice,
	}
)

func main() {

	newCar := car.NewCar(model.OpelCar, fullInfoCar)
	newPlane := plane.NewPlane(model.AirBusPlane,fullInfoPlane)
	newVisitor := frame.NewVisitor(model.AddPriceCar, model.SalePlane64)
	newCar.FullInfo()
	_,err := newCar.Accept(newVisitor)
	if err != nil{
		log.Println(err, "car")
	}
	newCar.FullInfo()
	fmt.Println("______________")
	newPlane.FullInfo()
	_, err = newPlane.Accept(newVisitor)
	if err != nil{
		log.Println(err,"plane")
	}
	newPlane.FullInfo()
	//newPlane.Accept(newVisitor)
}
