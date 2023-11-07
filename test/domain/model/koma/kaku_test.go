package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestKaku(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("後手番の場合", func(t *testing.T) {
			kaku := koma.Kaku{Location: model.Location{X: 5, Y: 5}}

			t.Run("右前は移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 6, Y: 6}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右前4マスは移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 9, Y: 9}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右後は移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 6, Y: 4}) {
					t.Errorf("should be true")
				}
			})

			t.Run("右後4マスは移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 9, Y: 1}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左前は移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 4, Y: 6}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左前4マスは移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 1, Y: 9}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左後は移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 4, Y: 4}) {
					t.Errorf("should be true")
				}
			})

			t.Run("左後4マスは移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 1, Y: 1}) {
					t.Errorf("should be true")
				}
			})

			t.Run("前は移動不可", func(t *testing.T) {
				if kaku.IsMovableTo(model.Location{X: 5, Y: 6}) {
					t.Errorf("should be false")
				}
			})

			t.Run("後ろは移動不可", func(t *testing.T) {
				if kaku.IsMovableTo(model.Location{X: 5, Y: 4}) {
					t.Errorf("should be false")
				}
			})

			t.Run("右は移動不可", func(t *testing.T) {
				if kaku.IsMovableTo(model.Location{X: 6, Y: 5}) {
					t.Errorf("should be false")
				}
			})

			t.Run("左は移動不可", func(t *testing.T) {
				if kaku.IsMovableTo(model.Location{X: 4, Y: 5}) {
					t.Errorf("should be false")
				}
			})
		})

		t.Run("先手番の場合", func(t *testing.T) {
			kaku := koma.Kaku{Location: model.Location{X: 5, Y: 5}, IsSente: true}

			// 角に関しては先手後手で動作が変わらないので最低限

			t.Run("右前は移動可", func(t *testing.T) {
				if !kaku.IsMovableTo(model.Location{X: 6, Y: 6}) {
					t.Errorf("should be true")
				}
			})
		})
	})
}
