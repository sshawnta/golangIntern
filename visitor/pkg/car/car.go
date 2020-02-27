package car

import (
	`fmt`

	"github.com/sshawnta/golangIntern/visitor/pkg/model"
)

// Visitor Implementation visitor for expansion functional
type Visitor interface {
	VisitCar(c Car) (int, error)
}

// Car Active actions that can be performed on the car
type Car interface {
	FullInfo() (modelCar string, price int, err error)
	Price() (res int, err error)
	Accept(v Visitor) (res int, err error)
}

type car struct {
	model    string
	fullInfo map[string]int
}

// Price Return price of car exemplar
func (c *car) Price() (res int, err error) {
	if _, ok := c.fullInfo[c.model]; !ok {
		err = fmt.Errorf(model.NotFoundModel)
		return
	}
	res = c.fullInfo[c.model]
	return

}

// FullInfo Get full info about car model
func (c *car) FullInfo() (modelCar string, price int, err error) {
	if _, ok := c.fullInfo[c.model]; !ok {
		err = fmt.Errorf(model.NotFoundModel)
		return
	}
	modelCar = c.model
	price = c.fullInfo[modelCar]
	fmt.Println(c.model, price)
	return
}

// Accept Increases the price
func (c *car) Accept(v Visitor) (res int, err error) {
	res, err = v.VisitCar(c)
	if err != nil {
		return
	}
	c.fullInfo[c.model] = res
	return
}

// NewCar Constructor for car. Entry model of car
func NewCar(modelCar string, fullInfo map[string]int) Car {
	return &car{
		model:    modelCar,
		fullInfo: fullInfo,
	}
}
