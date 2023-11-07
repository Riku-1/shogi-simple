package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
)

type Gin struct {
	model.Location
	IsSente bool
}

func (g Gin) IsBelongToSente() bool {
	return g.IsSente
}

func (g Gin) IsMovableTo(l model.Location) bool {
	if !model.CanLocation(l.X, l.Y) {
		return false
	}

	return g.Location.IsMovableTo(g, l)
}

func (g Gin) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: 1},   // 前
		{X: 1, Y: 1},   // 右前
		{X: 1, Y: -1},  // 右後
		{X: -1, Y: 1},  // 左前
		{X: -1, Y: -1}, // 左後
	}
}

func (g Gin) MoveTo(l model.Location) (model.Movable, error) {
	if !g.IsMovableTo(l) {
		return Gin{}, &error2.InvalidMovementError{}
	}

	return Gin{
		Location: l,
	}, nil
}
