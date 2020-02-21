package plane

import (
	`fmt`

	`github.com/sshawnta/visitor/pkg/model`
)

////Implementation visitor for expansion functional
type Visitor interface {
	VisitPlane(p Plane) float64
}

//Active actions that can be performed on the plane
type Plane interface {
	FullInfo() string
	Price() (float64, error)
	Accept(v Visitor) float64
}

type plane struct {
	model string
	price float64
}

//Return car of plane exemplar
func (p *plane) Price() (float64, error) {
	planeInfo := map[string]float64{
		model.BoeingPlane: model.BoeingPrice,
		model.AirPlane:    model.AirBusPrice,
		model.IlPlane:     model.IlPrice,
	}
	var err error
	if res, ok := planeInfo[p.model]; ok {
		return res, nil
	}
	err = fmt.Errorf(model.NotFoundModel)
	return 0, err
}

//Get full info about plane model
func (p *plane) FullInfo() string {
	res := p.model
	fmt.Println(p.model)
	return res
}

//Increases the sale
func (p *plane) Accept(v Visitor) float64 {
	res := v.VisitPlane(p)
	return res
}

func NewPlane(modelPlane string) Plane {
	return &plane{
		model: modelPlane,
	}
}
