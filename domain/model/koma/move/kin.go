package move

import "shogi/domain/model"

func GetKinMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
		{X: -1, Y: 1},
	}
}
