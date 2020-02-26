package frame

import (
	`fmt`

	"github.com/sshawnta/golangIntern/visitor/pkg/model"
)

//Description of mutable behavior u can change carPrise or make sale for Plane
type Visitor interface {
	VisitCar(c car) (int, error)
	VisitPlane(p plane) float64
}


type car interface {
	FullInfo()(model string, price int, err error)
	Price() (int, error)
}

type plane interface {
	FullInfo() string
	Price() (float64, error)
}

type visitor struct {
	car car
	slogan int
	sale   float64
}

//implementation expansion of car
func (v *visitor) VisitCar(c car) (int, error) {
	price, err := c.Price()
	if err != nil {
		err = fmt.Errorf(model.NotFoundModel)
		return 0,err
	}
	price = price + v.slogan
	fmt.Println(price)
	return price, err
}

//implementation expansion of plane, make sales price
func (v *visitor) VisitPlane(p plane) float64 {
	price, err := p.Price()
	if err != nil {
		fmt.Println(model.NotFoundModel)
	}
	price = price - (price * (v.sale / 100))
	fmt.Println(price)
	return price
}

//Constructor for Visitor. Entry value of change. 1st - for cat 2th - for plane
func NewVisitor(car car, add int, sal float64) Visitor {
	return &visitor{
		car: car,
		slogan: add,
		sale:   sal,
	}
}
