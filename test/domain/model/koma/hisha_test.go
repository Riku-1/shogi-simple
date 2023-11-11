package koma

import (
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestHisha(t *testing.T) {
	t.Run("IsMovableTo", func(t *testing.T) {
		t.Run("後手番の場合", func(t *testing.T) {
			hisha := koma.LocateHisha(model.Location{X: 5, Y: 5}, false)

			t.Run("前は移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 5, Y: 6})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("前4マスは移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 5, Y: 9})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("右は移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 6, Y: 5})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("右4マスは移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 9, Y: 5})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("左は移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 4, Y: 5})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("左4マスは移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 1, Y: 5})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("後ろは移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 5, Y: 4})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("後ろ4マスは移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 5, Y: 1})
				if err != nil {
					t.Errorf("should be nil")
				}
			})

			t.Run("右前は移動不可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 6, Y: 6})
				if err == nil {
					t.Errorf("should be error")
				}
			})
		})

		t.Run("先手番の場合", func(t *testing.T) {
			hisha := koma.LocateHisha(model.Location{X: 5, Y: 5}, true)

			// 飛車に関しては先手後手で動作が変わらないので最低限
			t.Run("前は移動可", func(t *testing.T) {
				_, err := hisha.MoveTo(model.Location{X: 5, Y: 4})
				if err != nil {
					t.Errorf("should be nil")
				}
			})
		})

		t.Run("成った後の動き", func(t *testing.T) {
			hisha := koma.LocateHisha(model.Location{X: 5, Y: 5}, false)
			promoted := hisha.Nari()

			// 他は省略
			t.Run("斜めにも動ける", func(t *testing.T) {
				_, err := promoted.MoveTo(model.Location{X: 4, Y: 4})
				if err != nil {
					t.Errorf("should be nil")
				}
			})
		})
	})
}
