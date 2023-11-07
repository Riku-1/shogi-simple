package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Kyousha struct {
	model.Location
	IsSente bool
}

func (k Kyousha) GetCurrentLocation() model.Location {
	return k.Location
}

func (k Kyousha) IsBelongToSente() bool {
	return k.IsSente
}

func (k Kyousha) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: 1},
		{X: 0, Y: 2},
		{X: 0, Y: 3},
		{X: 0, Y: 4},
		{X: 0, Y: 5},
		{X: 0, Y: 6},
		{X: 0, Y: 7},
		{X: 0, Y: 8},
	}
}

func (k Kyousha) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(k, l) {
		return Kyousha{}, &error2.InvalidMovementError{}
	}

	return Kyousha{
		Location: l,
	}, nil
}
