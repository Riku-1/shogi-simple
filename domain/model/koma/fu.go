package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Fu struct {
	model.Location
	IsSente bool
}

func (f Fu) GetCurrentLocation() model.Location {
	return f.Location
}

func (f Fu) IsBelongToSente() bool {
	return f.IsSente
}

func (f Fu) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: 1},
	}
}

func (f Fu) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(f, l) {
		return Fu{}, &error2.InvalidMovementError{}
	}

	return Fu{
		Location: l,
	}, nil
}
