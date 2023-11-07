package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Gyoku struct {
	model.Location
}

func (g Gyoku) IsMovableTo(l model.Location) bool {
	if !model.CanLocation(l.X, l.Y) {
		return false
	}

	return g.Location.IsMovableTo(g, l)
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
	if !g.IsMovableTo(l) {
		return Gyoku{}, &error2.InvalidMovementError{}
	}

	return Gyoku{
		Location: l,
	}, nil
}
