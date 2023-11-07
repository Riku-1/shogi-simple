package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Kin struct {
	model.Location
	IsSente bool
}

func (k Kin) GetCurrentLocation() model.Location {
	return k.Location
}

func (k Kin) IsBelongToSente() bool {
	return k.IsSente
}

func (k Kin) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
		{X: -1, Y: 1},
	}
}

func (k Kin) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(k, l) {
		return Kin{}, &error2.InvalidMovementError{}
	}

	return Kin{
		Location: l,
	}, nil
}
