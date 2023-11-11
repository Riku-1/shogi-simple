package model

import (
	"github.com/stretchr/testify/assert"
	"shogi/domain/model"
	"shogi/domain/model/koma"
	"testing"
)

func TestGetMovableLocations(t *testing.T) {
	t.Run("移動場所が盤外の場合は移動不可", func(t *testing.T) {
		fu := koma.LocateFu(model.Location{X: 5, Y: 9}, false)
		movableLocations := model.GetMovableLocations(fu)
		if len(movableLocations) != 0 {
			t.Errorf("should be 0")
		}
	})

	t.Run("移動場所が盤外でない場合は移動可", func(t *testing.T) {
		fu := koma.Fu{Location: model.Location{X: 5, Y: 5}}
		movableLocations := model.GetMovableLocations(fu)
		if len(movableLocations) != 1 {
			t.Errorf("should be 1")
		}
	})

	t.Run("移動可能場所が正しく取得できること", func(t *testing.T) {
		hisha := koma.LocateHisha(model.Location{X: 5, Y: 5}, false)
		movableLocations := model.GetMovableLocations(hisha)

		expected := []model.Location{
			// 上下
			{X: 5, Y: 1},
			{X: 5, Y: 2},
			{X: 5, Y: 3},
			{X: 5, Y: 4},
			{X: 5, Y: 6},
			{X: 5, Y: 7},
			{X: 5, Y: 8},
			{X: 5, Y: 9},

			// 左右
			{X: 1, Y: 5},
			{X: 2, Y: 5},
			{X: 3, Y: 5},
			{X: 4, Y: 5},
			{X: 6, Y: 5},
			{X: 7, Y: 5},
			{X: 8, Y: 5},
			{X: 9, Y: 5},
		}

		assert.ElementsMatch(t, expected, movableLocations)
	})
}

func TestIsMovableTo(t *testing.T) {
	t.Run("移動できない場合_false", func(t *testing.T) {
		fu := koma.LocateFu(model.Location{X: 5, Y: 5}, false)
		if model.IsMovableTo(fu, model.Location{X: 5, Y: 4}) {
			t.Errorf("should be false")
		}
	})

	t.Run("移動できる場合_true", func(t *testing.T) {
		fu := koma.LocateFu(model.Location{X: 5, Y: 5}, false)
		if !model.IsMovableTo(fu, model.Location{X: 5, Y: 6}) {
			t.Errorf("should be true")
		}
	})
}
