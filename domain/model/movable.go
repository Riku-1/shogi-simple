package model

type Movable interface {
	IsMovableTo(l Location) bool
	GetMovementCapabilities() []MovementCapability
}
