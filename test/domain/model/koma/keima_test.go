package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestKeima(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("5, 5にいる場合", func(t *testing.T) {
			keima := koma.Keima{Location: model.Location{X: 5, Y: 5}}
			t.Run("前2マス,右1マスは移動可", func(t *testing.T) {
				if !keima.IsMovableTo(model.Location{X: 4, Y: 3}) {
					t.Errorf("should be true")
				}
			})

			t.Run("前2マス,左1マスは移動可", func(t *testing.T) {
				if !keima.IsMovableTo(model.Location{X: 6, Y: 3}) {
					t.Errorf("should be true")
				}
			})

			t.Run("前は移動不可", func(t *testing.T) {
				if keima.IsMovableTo(model.Location{X: 5, Y: 4}) {
					t.Errorf("should be false")
				}
			})
		})
	})
}
