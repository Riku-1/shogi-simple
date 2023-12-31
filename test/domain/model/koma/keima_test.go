package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestKeima(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("後手番の場合", func(t *testing.T) {
			keima := koma.LocateKeima(model.Location{X: 5, Y: 5}, false)
			t.Run("前2マス,右1マスは移動可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 6, Y: 7})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("前2マス,左1マスは移動可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 4, Y: 7})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("前は移動不可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 5, Y: 6})
				if err == nil {
					t.Errorf("should be error")
				}
			})

			t.Run("先手番と違って4,3には移動不可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 4, Y: 3})
				if err == nil {
					t.Errorf("should be error")
				}
			})
		})

		t.Run("先手番の場合", func(t *testing.T) {
			keima := koma.LocateKeima(model.Location{X: 5, Y: 5}, true)
			t.Run("前2マス,右1マスは移動可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 4, Y: 3})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("前2マス,左1マスは移動可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 6, Y: 3})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("前は移動不可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 5, Y: 4})
				if err == nil {
					t.Errorf("should be error")
				}
			})

			t.Run("後手番と違って4,7には移動不可", func(t *testing.T) {
				_, err := keima.MoveTo(model.Location{X: 4, Y: 7})
				if err == nil {
					t.Errorf("should be error")
				}
			})
		})

		t.Run("成った後の動き", func(t *testing.T) {
			keima := koma.LocateKeima(model.Location{X: 5, Y: 5}, false)
			promoted := keima.Nari()

			// 他は省略
			t.Run("横にも動ける", func(t *testing.T) {
				_, err := promoted.MoveTo(model.Location{X: 4, Y: 5})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("成る前の動きはできない", func(t *testing.T) {
				_, err := promoted.MoveTo(model.Location{X: 4, Y: 7})
				if err == nil {
					t.Errorf("should not be nil")
				}
			})
		})
	})
}
