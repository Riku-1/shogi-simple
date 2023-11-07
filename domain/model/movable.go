package model

type Movable interface {
	IsMovableTo(l Location) bool
	GetMovementCapabilities() []MovementCapability
	MoveTo(l Location) (Movable, error)
	IsBelongToSente() bool
	// GetMovableLocation() []Location
}
