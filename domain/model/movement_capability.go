package model

type MovementCapability struct {
	X int
	Y int
}

func (c MovementCapability) GetActual(IsSente bool) (int, int) {
	if IsSente {
		return c.X * -1, c.Y * -1
	}

	return c.X, c.Y
}
