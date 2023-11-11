package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
	"shogi/domain/model/koma/move"
)

type Fu struct {
	model.Location
	isSente bool
	isNari  bool
}

func LocateFu(l model.Location, isSente bool) Fu {
	return Fu{
		Location: l,
		isSente:  isSente,
		isNari:   false,
	}
}

func (f Fu) Nari() model.Movable {
	return Fu{
		Location: f.Location,
		isSente:  f.isSente,
		isNari:   true,
	}
}

func (f Fu) IsNari() bool {
	return f.isNari
}

func (f Fu) GetCurrentLocation() model.Location {
	return f.Location
}

func (f Fu) IsSente() bool {
	return f.isSente
}

func (f Fu) GetMovementCapabilities() []model.MovementCapability {
	if f.IsNari() {
		return move.GetKinMovementCapabilities()
	}

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
