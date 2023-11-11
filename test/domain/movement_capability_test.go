package domain

import (
	"shogi/domain/model"
	"testing"
)

func TestGetActual(t *testing.T) {
	t.Run("先手番の場合_反転", func(t *testing.T) {
		c := model.MovementCapability{X: 1, Y: 1}
		x, y := c.GetActual(true)
		if x != -1 {
			t.Errorf("x should be -1")
		}
		if y != -1 {
			t.Errorf("y should be -1")
		}
	})

	t.Run("後手番の場合_そのまま", func(t *testing.T) {
		c := model.MovementCapability{X: 1, Y: 1}
		x, y := c.GetActual(false)
		if x != 1 {
			t.Errorf("x should be 1")
		}
		if y != 1 {
			t.Errorf("y should be 1")
		}
	})
}
