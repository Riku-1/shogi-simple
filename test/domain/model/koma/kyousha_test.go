package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestKyousha(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("後手番の場合", func(t *testing.T) {
			kyousha := koma.LocateKyousha(model.Location{X: 1, Y: 1}, false)

			t.Run("前1マスは移動可", func(t *testing.T) {
				_, err := kyousha.MoveTo(model.Location{X: 1, Y: 2})

				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("前8マスは移動可", func(t *testing.T) {
				_, err := kyousha.MoveTo(model.Location{X: 1, Y: 9})

				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("横には移動できない", func(t *testing.T) {
				_, err := kyousha.MoveTo(model.Location{X: 2, Y: 1})

				if err == nil {
					t.Errorf("should be error")
				}
			})
		})

		t.Run("先手番の場合", func(t *testing.T) {
			kyousha := koma.LocateKyousha(model.Location{X: 9, Y: 9}, true)

			t.Run("前1マスは移動可", func(t *testing.T) {
				_, err := kyousha.MoveTo(model.Location{X: 9, Y: 8})

				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("前8マスは移動可", func(t *testing.T) {
				_, err := kyousha.MoveTo(model.Location{X: 9, Y: 1})

				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("横には移動できない", func(t *testing.T) {
				_, err := kyousha.MoveTo(model.Location{X: 8, Y: 9})

				if err == nil {
					t.Errorf("should be error")
				}
			})
		})
	})
}
