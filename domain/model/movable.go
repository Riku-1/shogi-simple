package model

type Movable interface {
	GetCurrentLocation() Location
	GetMovementCapabilities() []MovementCapability
	MoveTo(l Location) (Movable, error)
	IsSente() bool
}

func IsMovableTo(m Movable, to Location) bool {
	ml := GetMovableLocations(m)

	for _, l := range ml {
		if l.X == to.X && l.Y == to.Y {
			return true
		}
	}

	return false
}

func GetMovableLocations(m Movable) []Location {
	cl := m.GetCurrentLocation()
	movementCapabilities := m.GetMovementCapabilities()
	var locations []Location

	for _, movementCapability := range movementCapabilities {
		moveX, moveY := movementCapability.GetActual(m.IsSente())
		x := cl.X + moveX
		y := cl.Y + moveY
		if IsValidLocation(x, y) {
			locations = append(locations, Location{X: x, Y: y})
		}
	}

	return locations
}
