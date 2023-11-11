package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
	"shogi/domain/model/koma/move"
)

type Kin struct {
	location model.Location
	isSente  bool
	isNari   bool
}

func LocateKin(l model.Location, isSente bool) Kin {
	return Kin{
		location: l,
		isSente:  isSente,
		isNari:   false,
	}
}

func (k Kin) GetCurrentLocation() model.Location {
	return k.location
}

func (k Kin) IsSente() bool {
	return k.isSente
}

func (k Kin) GetMovementCapabilities() []model.MovementCapability {
	if k.isNari {
		return move.GetKinMovementCapabilities()
	}

	return move.GetKinMovementCapabilities()
}

func (k Kin) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(k, l) {
		return Kin{}, &error2.InvalidMovementError{}
	}

	return Kin{
		location: l,
		isSente:  k.isSente,
		isNari:   k.isNari,
	}, nil
}
