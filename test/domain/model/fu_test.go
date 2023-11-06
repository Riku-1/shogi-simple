package model

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestFu(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("9, 9にいる場合", func(t *testing.T) {
			fu := koma.Fu{Location: model.Location{X: 9, Y: 9}}

			t.Run("9, 8は移動可", func(t *testing.T) {
				if !fu.IsMovableTo(model.Location{X: 9, Y: 8}) {
					t.Errorf("should be true")
				}
			})

			t.Run("9, 7は移動不可", func(t *testing.T) {
				if fu.IsMovableTo(model.Location{X: 9, Y: 7}) {
					t.Errorf("should be false")
				}
			})
		})
	})
}
