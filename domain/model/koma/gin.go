package koma

import (
	error2 "shogi/domain/error"
	"shogi/domain/model"
	"shogi/domain/model/koma/move"
)

type Gin struct {
	location model.Location
	isSente  bool
	isNari   bool
}

func (g Gin) Nari() model.Movable {
	return Gin{
		location: g.location,
		isSente:  g.isSente,
		isNari:   true,
	}
}

func LocateGin(l model.Location, isSente bool) Gin {
	return Gin{
		location: l,
		isSente:  isSente,
		isNari:   false,
	}
}

func (g Gin) IsNari() bool {
	return g.isNari
}

func (g Gin) GetCurrentLocation() model.Location {
	return g.location
}

func (g Gin) IsSente() bool {
	return g.isSente
}

func (g Gin) GetMovementCapabilities() []model.MovementCapability {
	if g.isNari {
		return move.GetKinMovementCapabilities()
	}

	return []model.MovementCapability{
		{X: 0, Y: 1},   // 前
		{X: 1, Y: 1},   // 右前
		{X: 1, Y: -1},  // 右後
		{X: -1, Y: 1},  // 左前
		{X: -1, Y: -1}, // 左後
	}
}

func (g Gin) MoveTo(l model.Location) (model.Movable, error) {
	if !model.IsMovableTo(g, l) {
		return Gin{}, &error2.InvalidMovementError{}
	}

	return Gin{
		location: l,
	}, nil
}
