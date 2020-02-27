package frame

import (
	`github.com/sshawnta/golangIntern/visitor/pkg/car`
	`github.com/sshawnta/golangIntern/visitor/pkg/plane`
)

// Description of mutable behavior u can change carPrise or make sale for Plane
type Visitor interface {
	VisitCar(c car.Car) (price int, err error)
	VisitPlane(p plane.Plane) (res float64, err error)
}

type visitor struct {
	slogan int
	sale   float64
}

// implementation expansion of car
func (v *visitor) VisitCar(c car.Car) (price int, err error) {
	price, err = c.Price()
	if err != nil {
		return
	}
	price = price + v.slogan
	return
}

// implementation expansion of plane, make sales price
func (v *visitor) VisitPlane(p plane.Plane) (res float64, err error) {
	price, err := p.Price()
	if err != nil {
		return
	}
	res = price - (price * (v.sale / 100))
	return
}

// Constructor for Visitor. Entry value of change. 1st - for cat 2th - for plane
func NewVisitor(add int, sal float64) Visitor {
	return &visitor{
		slogan: add,
		sale:   sal,
	}
}
