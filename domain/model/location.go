package model

import (
	"fmt"
)

const MinLocation = 1
const MaxLocation = 9

// Location 将棋の駒の位置を表す構造体
// 左下が1,1で右上が9,9(先手から見た形)を基準とする
type Location struct {
	X int
	Y int
}

// NewLocation Locationのコンストラクタ
func NewLocation(x, y int) (*Location, error) {
	if !IsValidLocation(x, y) {
		return nil, fmt.Errorf("xまたはyが範囲外です。x:%d, y:%d", x, y)
	}

	return &Location{
		X: x,
		Y: y,
	}, nil
}

func IsValidLocation(x, y int) bool {
	return MinLocation <= x && x <= MaxLocation && MinLocation <= y && y <= MaxLocation
}
