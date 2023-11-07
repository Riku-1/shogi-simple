package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Keima struct {
	model.Location
	IsSente bool
}

func (k Keima) GetCurrentLocation() model.Location {
	return k.Location
}

func (k Keima) IsBelongToSente() bool {
	return k.IsSente
}

func (k Keima) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 1, Y: 2},
		{X: -1, Y: 2},
	}
}

func (k Keima) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(k, l) {
		return Keima{}, &error2.InvalidMovementError{}
	}

	return Keima{
		Location: l,
	}, nil
}
