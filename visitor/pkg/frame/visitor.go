package frame

import (
	`fmt`

	`github.com/sshawnta/visitor/pkg/car`
	`github.com/sshawnta/visitor/pkg/model`
	`github.com/sshawnta/visitor/pkg/plane`
)

//Description of mutable behavior
type Visitor interface {
	VisitCar(c car.Car) int
	VisitPlane(p plane.Plane) float64
}

type visitor struct {
	slogan int
	sale   float64
}

//implementation expansion of car
func (v *visitor) VisitCar(c car.Car) int {
	price, err := c.Price()
	if err != nil {
		fmt.Println(model.NotFoundModel)
		return 0
	}
	price = price + v.slogan
	fmt.Println(price)
	return price
}

//implementation expansion of plane, make sales price
func (v *visitor) VisitPlane(p plane.Plane) float64 {
	price, err := p.Price()
	if err != nil {
		fmt.Println(model.NotFoundModel)
	}
	price = price - (price * (v.sale / 100))
	fmt.Println(price)
	return price
}

func NewVisitor(add int, sal float64) Visitor {
	return &visitor{
		slogan: add,
		sale:   sal,
	}
}
