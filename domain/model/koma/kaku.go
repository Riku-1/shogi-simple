package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Kaku struct {
	model.Location
	IsSente bool
}

func (k Kaku) GetCurrentLocation() model.Location {
	return k.Location
}

func (k Kaku) IsBelongToSente() bool {
	return k.IsSente
}

// GetMovementCapabilities
// NOTE: 選択肢が多いので盤外の選択肢は除外できるように実装を変更してもいいかもしれない
func (k Kaku) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
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
}

func (k Kaku) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(k, l) {
		return Kaku{}, &error2.InvalidMovementError{}
	}

	return Kaku{
		Location: l,
	}, nil
}
