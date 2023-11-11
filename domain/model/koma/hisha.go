package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Hisha struct {
	location model.Location
	isSente  bool
	isNari   bool
}

func LocateHisha(l model.Location, isSente bool) Hisha {
	return Hisha{
		location: l,
		isSente:  isSente,
		isNari:   false,
	}
}

func (h Hisha) IsNari() bool {
	return h.isNari
}

func (h Hisha) Nari() model.Movable {
	return Hisha{
		location: h.location,
		isSente:  h.isSente,
		isNari:   true,
	}
}

func (h Hisha) GetCurrentLocation() model.Location {
	return h.location
}

func (h Hisha) IsSente() bool {
	return h.isSente
}

func (h Hisha) GetMovementCapabilities() []model.MovementCapability {
	baseMc := []model.MovementCapability{
		// 上
		{X: 0, Y: 1},
		{X: 0, Y: 2},
		{X: 0, Y: 3},
		{X: 0, Y: 4},
		{X: 0, Y: 5},
		{X: 0, Y: 6},
		{X: 0, Y: 7},
		{X: 0, Y: 8},

		// 右
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 5, Y: 0},
		{X: 6, Y: 0},
		{X: 7, Y: 0},
		{X: 8, Y: 0},

		// 左
		{X: -1, Y: 0},
		{X: -2, Y: 0},
		{X: -3, Y: 0},
		{X: -4, Y: 0},
		{X: -5, Y: 0},
		{X: -6, Y: 0},
		{X: -7, Y: 0},
		{X: -8, Y: 0},

		// 下
		{X: 0, Y: -1},
		{X: 0, Y: -2},
		{X: 0, Y: -3},
		{X: 0, Y: -4},
		{X: 0, Y: -5},
		{X: 0, Y: -6},
		{X: 0, Y: -7},
		{X: 0, Y: -8},
	}

	if !h.IsNari() {
		return baseMc
	}

	addedMc := []model.MovementCapability{
		{X: 1, Y: 1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: -1, Y: -1},
	}

	return append(baseMc, addedMc...)
}

func (h Hisha) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(h, l) {
		return Hisha{}, &error2.InvalidMovementError{}
	}

	return Hisha{
		location: l,
	}, nil
}
