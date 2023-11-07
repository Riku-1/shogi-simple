package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Fu struct {
	model.Location
}

func (f Fu) IsMovableTo(to model.Location) bool {
	if !model.CanLocation(to.X, to.Y) {
		return false
	}

	return f.Location.IsMovableTo(f, to)
}

func (f Fu) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: -1},
	}
}

func (f Fu) MoveTo(l model.Location) (model.Movable, error) {
	if !f.IsMovableTo(l) {
		return Fu{}, &error2.InvalidMovementError{}
	}

	return Fu{
		Location: l,
	}, nil
}
