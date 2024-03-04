package rigidbody

import "physics/pkg/vector"

// RigidBody represents a 2D rigid body.
type RigidBody struct {
	Position vector.Vector
	Velocity vector.Vector
	Mass     float64
	Circle   Circle  
}
