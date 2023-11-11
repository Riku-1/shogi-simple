package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
	"shogi/domain/model/koma/move"
)

type Keima struct {
	location model.Location
	isSente  bool
	isNari   bool
}

func (k Keima) IsNari() bool {
	return k.isNari
}

func (k Keima) Nari() model.Movable {
	return Keima{
		location: k.location,
		isSente:  k.isSente,
		isNari:   true,
	}
}

func LocateKeima(l model.Location, isSente bool) Keima {
	return Keima{
		location: l,
		isSente:  isSente,
		isNari:   false,
	}
}

func (k Keima) GetCurrentLocation() model.Location {
	return k.location
}

func (k Keima) IsSente() bool {
	return k.isSente
}

func (k Keima) GetMovementCapabilities() []model.MovementCapability {
	if k.isNari {
		return move.GetKinMovementCapabilities()
	}

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
		location: l,
	}, nil
}
