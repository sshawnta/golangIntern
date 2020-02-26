package car

import (
	`fmt`

	"github.com/sshawnta/golangIntern/visitor/pkg/model"
)

//Implementation visitor for expansion functional
type Visitor interface {
	VisitCar() (int, error)
}

//Active actions that can be performed on the car
type Car interface {
	FullInfo()(model string, price int, err error)
	Price() (int, error)
	Accept(v Visitor) (int,error)
}

type car struct {
	model string
	fullinfo map[string]int
}


//Return price of car exemplar
func (c *car) Price() (int, error) {
	//carInfo := c.makeDate()
	var err error
	if res, ok := c.fullinfo[c.model]; ok {
		return res, nil
	}
	err = fmt.Errorf(model.NotFoundModel)
	return 0, err
}

//Get full info about car model
func (c *car) FullInfo() (modelCar string, price int, err error) {
	if res, ok:=c.fullinfo[c.model]; ok{
		fmt.Println(c.model, res)
		return c.model, res,nil
	}
	err = fmt.Errorf(model.NotFoundModel)
	return c.model,0, err
}

//Increases the price
func (c *car) Accept(v Visitor) (int,error) {
	res, err := v.VisitCar()
	if err != nil{
		err = fmt.Errorf(model.NotFoundModel)
		return 0, err
	}
	c.fullinfo[c.model] = res
	return res, nil
}

/*func  (c *car)makeDate() map[string]int {
	return map[string]int{
		model.OpelCar:  model.OpelPrice,
		model.MazdaCar: model.MazdaPrice,
		model.BMWCar:   model.BMWPrice,
	}
}*/

//Constructor for car. Entry model of car
func NewCar(modelCar string, fullinfo map[string]int) Car {
	return &car{
		model: modelCar,
		fullinfo:fullinfo,
	}
}
