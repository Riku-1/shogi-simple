package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestKyousha(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("後手番の場合", func(t *testing.T) {
			kyousha := koma.Kyousha{Location: model.Location{X: 1, Y: 1}}

			t.Run("前1マスは移動可", func(t *testing.T) {
				if !kyousha.IsMovableTo(model.Location{X: 1, Y: 2}) {
					t.Errorf("should be true")
				}
			})

			t.Run("前8マスは移動可", func(t *testing.T) {
				if !kyousha.IsMovableTo(model.Location{X: 1, Y: 9}) {
					t.Errorf("should be true")
				}
			})

			t.Run("横には移動できない", func(t *testing.T) {
				if kyousha.IsMovableTo(model.Location{X: 2, Y: 1}) {
					t.Errorf("should be false")
				}
			})
		})

		t.Run("先手番の場合", func(t *testing.T) {
			kyousha := koma.Kyousha{Location: model.Location{X: 9, Y: 9}, IsSente: true}

			t.Run("前1マスは移動可", func(t *testing.T) {
				if !kyousha.IsMovableTo(model.Location{X: 9, Y: 8}) {
					t.Errorf("should be true")
				}
			})

			t.Run("前8マスは移動可", func(t *testing.T) {
				if !kyousha.IsMovableTo(model.Location{X: 9, Y: 1}) {
					t.Errorf("should be true")
				}
			})

			t.Run("横には移動できない", func(t *testing.T) {
				if kyousha.IsMovableTo(model.Location{X: 8, Y: 9}) {
					t.Errorf("should be false")
				}
			})
		})
	})
}
