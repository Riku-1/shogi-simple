package model

import (
	"shogi/domain/model"
	"testing"
)

func TestLocation(t *testing.T) {
	t.Run("NewLocation", func(t *testing.T) {
		t.Run("xとyが範囲内の場合はLocationを返す", func(t *testing.T) {
			_, err := model.NewLocation(model.MinLocation, model.MinLocation)
			if err != nil {
				t.Errorf("err should be nil, but %v", err)
			}
		})

		t.Run("xが範囲外の場合はエラーを返す", func(t *testing.T) {
			_, err := model.NewLocation(model.MinLocation, model.MinLocation-1)
			if err == nil {
				t.Errorf("err should not be nil")
			}
		})

		t.Run("yが範囲外の場合はエラーを返す", func(t *testing.T) {
			_, err := model.NewLocation(model.MinLocation, model.MaxLocation+1)
			if err == nil {
				t.Errorf("err should not be nil")
			}
		})
	})
}
