package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Gyoku struct {
	location model.Location
	isSente  bool
}

func LocateGyoku(l model.Location, isSente bool) Gyoku {
	return Gyoku{
		location: l,
		isSente:  isSente,
	}
}

func (g Gyoku) GetCurrentLocation() model.Location {
	return g.location
}

func (g Gyoku) IsSente() bool {
	return g.isSente
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
		location: l,
	}, nil
}
