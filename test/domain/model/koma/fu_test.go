package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestFu(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("後手番", func(t *testing.T) {
			// 5, 5にいるとする
			fu := koma.Fu{Location: model.Location{X: 5, Y: 5}, IsSente: false}

			t.Run("5, 6は移動可", func(t *testing.T) {
				if !fu.IsMovableTo(model.Location{X: 5, Y: 6}) {
					t.Errorf("should be true")
				}
			})

			t.Run("5, 4は移動不可", func(t *testing.T) {
				if fu.IsMovableTo(model.Location{X: 5, Y: 4}) {
					t.Errorf("should be false")
				}
			})
		})

		t.Run("先手番", func(t *testing.T) {
			// 5, 5にいるとする
			fu := koma.Fu{Location: model.Location{X: 5, Y: 5}, IsSente: true}

			t.Run("5, 4は移動可", func(t *testing.T) {
				if !fu.IsMovableTo(model.Location{X: 5, Y: 4}) {
					t.Errorf("should be true")
				}
			})

			t.Run("5, 6は移動不可", func(t *testing.T) {
				if fu.IsMovableTo(model.Location{X: 5, Y: 6}) {
					t.Errorf("should be false")
				}
			})
		})
	})
}
