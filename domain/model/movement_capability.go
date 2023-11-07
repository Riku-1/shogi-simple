package model

type MovementCapability struct {
	X int
	Y int
}

func (c MovementCapability) GetActual(isSente bool) (int, int) {
	if isSente {
		return c.X * -1, c.Y * -1
	}

	return c.X, c.Y
}
