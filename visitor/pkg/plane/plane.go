package plane

import (
	`fmt`

	"github.com/sshawnta/golangIntern/visitor/pkg/model"
)

// Visitor Implementation visitor for expansion functional
type Visitor interface {
	VisitPlane(p Plane) (res float64, err error)
}

// Plane Active actions that can be performed on the plane
type Plane interface {
	FullInfo() (err error)
	Price() (res float64, err error)
	Accept(v Visitor) (res float64, err error)
}

type plane struct {
	model    string
	fullInfo map[string]float64
}

// Price Return price of car exemplar
func (p *plane) Price() (res float64, err error) {
	if _, ok := p.fullInfo[p.model]; !ok {
		err = fmt.Errorf(model.NotFoundModel)
		return
	}
	res = p.fullInfo[p.model]
	return
}

// FullInfo Get full info about plane model
func (p *plane) FullInfo() (err error) {
	if _, ok := p.fullInfo[p.model]; !ok {
		err = fmt.Errorf(model.NotFoundModel)
		return
	}
	price := p.fullInfo[p.model]
	fmt.Println(p.model, price)
	return
}

// Accept Increases the sale
func (p *plane) Accept(v Visitor) (res float64, err error) {
	res, err = v.VisitPlane(p)
	if err != nil {
		return
	}
	p.fullInfo[p.model] = res
	return
}

// NewPlane Constructor for Plane. Entry model of plane
func NewPlane(modelPlane string, fullInfo map[string]float64) Plane {
	return &plane{
		model:    modelPlane,
		fullInfo: fullInfo,
	}
}
