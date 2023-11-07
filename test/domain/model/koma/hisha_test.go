package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestHisha(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("5, 5にいる場合", func(t *testing.T) {
			hisha := koma.Hisha{Location: model.Location{X: 5, Y: 5}}

			t.Run("前は移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 5, Y: 4}) {
					t.Errorf("should be true")
				}
			})

			t.Run("前4マスは移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 5, Y: 1}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右は移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 4, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右4マスは移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 1, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左は移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 6, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左4マスは移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 9, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("後ろは移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 5, Y: 6}) {
					t.Errorf("should be true")
				}
			})

			t.Run("後ろ4マスは移動可", func(t *testing.T) {
				if !hisha.IsMovableTo(model.Location{X: 5, Y: 9}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右前は移動不可", func(t *testing.T) {
				if hisha.IsMovableTo(model.Location{X: 4, Y: 4}) {
					t.Errorf("should be false")
				}
			})

		})
	})
}
