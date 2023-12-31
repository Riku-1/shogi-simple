package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestFu(t *testing.T) {
	t.Run("MoveTo", func(t *testing.T) {
		t.Run("後手番", func(t *testing.T) {
			// 5, 5にいるとする
			fu := koma.LocateFu(model.Location{X: 5, Y: 5}, false)

			t.Run("5, 6は移動可", func(t *testing.T) {
				_, err := fu.MoveTo(model.Location{X: 5, Y: 6})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("5, 4は移動不可", func(t *testing.T) {
				_, err := fu.MoveTo(model.Location{X: 5, Y: 4})
				if err == nil {
					t.Errorf("should be error")
				}
			})
		})

		t.Run("先手番", func(t *testing.T) {
			// 5, 5にいるとする
			fu := koma.LocateFu(model.Location{X: 5, Y: 5}, true)

			t.Run("5, 4は移動可", func(t *testing.T) {
				_, err := fu.MoveTo(model.Location{X: 5, Y: 4})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("5, 6は移動不可", func(t *testing.T) {
				_, err := fu.MoveTo(model.Location{X: 5, Y: 6})
				if err == nil {
					t.Errorf("should be error")
				}
			})
		})

		t.Run("成った後の動き", func(t *testing.T) {
			fu := koma.LocateFu(model.Location{X: 5, Y: 5}, false)
			promoted := fu.Nari()

			// 他は省略
			t.Run("5, 6は移動可", func(t *testing.T) {
				_, err := promoted.MoveTo(model.Location{X: 5, Y: 6})
				if err != nil {
					t.Errorf("should be nil")
				}
			})
		})
	})
}
