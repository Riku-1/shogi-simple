package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Gyoku struct {
	model.Location
	IsSente bool
}

func (g Gyoku) GetCurrentLocation() model.Location {
	return g.Location
}

func (g Gyoku) IsBelongToSente() bool {
	return g.IsSente
}

func (g Gyoku) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 1, Y: -1},
		{X: 0, Y: -1},
		{X: -1, Y: -1},
		{X: -1, Y: 0},
		{X: -1, Y: 1},
	}
}

func (g Gyoku) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(g, l) {
		return Gyoku{}, &error2.InvalidMovementError{}
	}

	return Gyoku{
		Location: l,
	}, nil
}
