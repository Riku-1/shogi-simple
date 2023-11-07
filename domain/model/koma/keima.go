package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Keima struct {
	model.Location
}

func (k Keima) IsMovableTo(to model.Location) bool {
	if !model.CanLocation(to.X, to.Y) {
		return false
	}

	return k.Location.IsMovableTo(k, to)
}

func (k Keima) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: -1, Y: -2},
		{X: 1, Y: -2},
	}
}

func (k Keima) MoveTo(l model.Location) (model.Movable, error) {
	if !k.IsMovableTo(l) {
		return Keima{}, &error2.InvalidMovementError{}
	}

	return Keima{
		Location: l,
	}, nil
}
