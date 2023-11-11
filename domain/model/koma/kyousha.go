package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Kyousha struct {
	location model.Location
	isSente  bool
	isNari   bool
}

func LocateKyousha(location model.Location, isSente bool) Kyousha {
	return Kyousha{
		location: location,
		isSente:  isSente,
		isNari:   false,
	}
}

func (k Kyousha) Nari() model.Movable {
	return Kyousha{
		location: k.location,
		isSente:  k.isSente,
		isNari:   true,
	}
}

func (k Kyousha) IsNari() bool {
	return k.isNari
}

func (k Kyousha) GetCurrentLocation() model.Location {
	return k.location
}

func (k Kyousha) IsSente() bool {
	return k.isSente
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
		location: l,
		isSente:  k.isSente,
		isNari:   k.isNari,
	}, nil
}
