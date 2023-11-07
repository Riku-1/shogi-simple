package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestKin(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("後手番の場合", func(t *testing.T) {
			kin := koma.Kin{Location: model.Location{X: 5, Y: 5}}
			t.Run("前は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 5, Y: 6}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右前は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 6, Y: 6}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 6, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右後は移動不可", func(t *testing.T) {
				if kin.IsMovableTo(model.Location{X: 6, Y: 4}) {
					t.Errorf("should be false")
				}
			})

			t.Run("左前は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 4, Y: 6}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 4, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左後は移動不可", func(t *testing.T) {
				if kin.IsMovableTo(model.Location{X: 4, Y: 4}) {
					t.Errorf("should be false")
				}
			})

			t.Run("後ろは移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 5, Y: 4}) {
					t.Errorf("should be true")
				}
			})
		})

		t.Run("先手番の場合", func(t *testing.T) {
			kin := koma.Kin{Location: model.Location{X: 5, Y: 5}, IsSente: true}
			t.Run("前は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 5, Y: 4}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右前は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 4, Y: 4}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 4, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右後は移動不可", func(t *testing.T) {
				if kin.IsMovableTo(model.Location{X: 4, Y: 6}) {
					t.Errorf("should be false")
				}
			})

			t.Run("左前は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 6, Y: 4}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左は移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 6, Y: 5}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左後は移動不可", func(t *testing.T) {
				if kin.IsMovableTo(model.Location{X: 6, Y: 6}) {
					t.Errorf("should be false")
				}
			})

			t.Run("後ろは移動可", func(t *testing.T) {
				if !kin.IsMovableTo(model.Location{X: 5, Y: 6}) {
					t.Errorf("should be true")
				}
			})
		})
	})
}
