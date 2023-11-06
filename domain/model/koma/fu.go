package koma

import "shogi/domain/model"

type Fu struct {
	model.Location
}

func (f Fu) IsMovableTo(to model.Location) bool {
	if !model.CanLocation(to.X, to.Y) {
		return false
	}

	return f.Location.IsMovableTo(f, to)
}

func (f Fu) GetMovementCapabilities() []model.MovementCapability {
	return []model.MovementCapability{
		{X: 0, Y: -1},
	}
}

func (f Fu) MoveTo(l model.Location) (Fu, error) {
	if !f.IsMovableTo(l) {
		return Fu{}, nil
	}

	return Fu{
		Location: l,
	}, nil
}
