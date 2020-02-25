package car

import (
	`fmt`

	"github.com/sshawnta/golangIntern/visitor/pkg/model"
)

//Implementation visitor for expansion functional
type Visitor interface {
	VisitCar(c car) int
}

//Active actions that can be performed on the car
type Car interface {
	FullInfo() string
	Price() (int, error)
	Accept(v Visitor) int
}

type car struct {
	model string
	price int
}

//Return price of car exemplar
func (c *car) Price() (int, error) {
	carInfo := c.makeDate()
	var err error
	if res, ok := carInfo[c.model]; ok {
		return res, nil
	}
	err = fmt.Errorf(model.NotFoundModel)
	return 0, err
}

//Get full info about car model
func (c *car) FullInfo() string {
	fmt.Println(c.model)
	return c.model
}

//Increases the price
func (c *car) Accept(v Visitor) int {
	res := v.VisitCar(c)
	return res
}

func (c *car) makeDate() map[string]int {
	return map[string]int{
		model.OpelCar:  model.OpelPrice,
		model.MazdaCar: model.MazdaPrice,
		model.BMWCar:   model.BMWPrice,
	}
}

//Constructor for car. Entry model of car
func NewCar(modelCar string) Car {
	return &car{
		model: modelCar,
	}
}
