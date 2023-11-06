package model

import (
	"fmt"
)

const MinLocation = 1
const MaxLocation = 9

// Location 将棋の駒の位置を表す構造体
// 右上が1,1で左下が9,9
type Location struct {
	X int
	Y int
}

// NewLocation Locationのコンストラクタ
// 解析など重い処理に使用したい場合は毎回バリデーションすると重いかもしれない。今回はそこまで実装する予定はないのでこのままで
func NewLocation(x, y int) (*Location, error) {
	if !CanLocation(x, y) {
		return nil, fmt.Errorf("xまたはyが範囲外です。x:%d, y:%d", x, y)
	}

	return &Location{
		X: x,
		Y: y,
	}, nil
}

func CanLocation(x, y int) bool {
	return MinLocation <= x && x <= MaxLocation && MinLocation <= y && y <= MaxLocation
}

func (l Location) IsMovableTo(m Movable, to Location) bool {
	for _, m := range m.GetMovementCapabilities() {
		if l.X+m.X == to.X && l.Y+m.Y == to.Y {
			return true
		}
	}

	return false
}
