package rigidbody

import "physix/pkg/vector"

var Infinite_mass float64 = 1e10

// RigidBody represents a 2D rigid body.
type RigidBody struct {
	Position    vector.Vector
	Velocity    vector.Vector
	Force       vector.Vector
	Mass        float64 
	Shape       string
	Width       float64
	Height      float64
	Radius      float64
	IsMovable   bool
}

