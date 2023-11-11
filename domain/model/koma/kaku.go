package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Kaku struct {
	location model.Location
	isSente  bool
	isNari   bool
}

func LocateKaku(l model.Location, isSente bool) Kaku {
	return Kaku{
		location: l,
		isSente:  isSente,
		isNari:   false,
	}
}

func (k Kaku) Nari() model.Movable {
	return Kaku{
		location: k.location,
		isSente:  k.isSente,
		isNari:   true,
	}
}

func (k Kaku) IsNari() bool {
	return k.isNari
}

func (k Kaku) GetCurrentLocation() model.Location {
	return k.location
}

func (k Kaku) IsSente() bool {
	return k.isSente
}

// GetMovementCapabilities
// NOTE: 選択肢が多いので盤外の選択肢は除外できるように実装を変更してもいいかもしれない
func (k Kaku) GetMovementCapabilities() []model.MovementCapability {
	baseMc := []model.MovementCapability{
		// 右上
		{X: 1, Y: 1},
		{X: 2, Y: 2},
		{X: 3, Y: 3},
		{X: 4, Y: 4},
		{X: 5, Y: 5},
		{X: 6, Y: 6},
		{X: 7, Y: 7},
		{X: 8, Y: 8},

		// 右下
		{X: 1, Y: -1},
		{X: 2, Y: -2},
		{X: 3, Y: -3},
		{X: 4, Y: -4},
		{X: 5, Y: -5},
		{X: 6, Y: -6},
		{X: 7, Y: -7},
		{X: 8, Y: -8},

		// 左下
		{X: -1, Y: -1},
		{X: -2, Y: -2},
		{X: -3, Y: -3},
		{X: -4, Y: -4},
		{X: -5, Y: -5},
		{X: -6, Y: -6},
		{X: -7, Y: -7},
		{X: -8, Y: -8},

		// 左上
		{X: -1, Y: 1},
		{X: -2, Y: 2},
		{X: -3, Y: 3},
		{X: -4, Y: 4},
		{X: -5, Y: 5},
		{X: -6, Y: 6},
		{X: -7, Y: 7},
		{X: -8, Y: 8},
	}

	if !k.IsNari() {
		return baseMc
	}

	addedMc := []model.MovementCapability{
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
	}

	return append(baseMc, addedMc...)
}

func (k Kaku) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(k, l) {
		return Kaku{}, &error2.InvalidMovementError{}
	}

	return Kaku{
		location: l,
	}, nil
}
