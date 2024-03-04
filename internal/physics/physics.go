package physics

import (
	"physics/pkg/vector"
	"physics/pkg/rigidbody"
)

// UpdateRigidBody updates the position of a rigid body using Euler integration.
func UpdateRigidBody(rb *rigidbody.RigidBody, force vector.Vector, dt float64) {
	// Calculate acceleration using Newton's second law: F = ma -> a = F/m
	acceleration := force.Scale(1 / rb.Mass)

	// Update velocity using acceleration and time step
	rb.Velocity = rb.Velocity.Add(acceleration.Scale(dt))

	// Update position using velocity and time step
	rb.Position = rb.Position.Add(rb.Velocity.Scale(dt))
}
