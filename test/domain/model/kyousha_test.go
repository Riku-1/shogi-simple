package model

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestKyousha(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("9, 9にいる場合", func(t *testing.T) {
			kyousha := koma.Kyousha{Location: model.Location{X: 9, Y: 9}}

			t.Run("9, 8は移動可", func(t *testing.T) {
				if !kyousha.IsMovableTo(model.Location{X: 9, Y: 8}) {
					t.Errorf("should be true")
				}
			})

			t.Run("9, 1は移動可", func(t *testing.T) {
				if !kyousha.IsMovableTo(model.Location{X: 9, Y: 1}) {
					t.Errorf("should be true")
				}
			})

			t.Run("横には移動できない", func(t *testing.T) {
				if kyousha.IsMovableTo(model.Location{X: 8, Y: 1}) {
					t.Errorf("should be false")
				}
			})
		})
	})
}
