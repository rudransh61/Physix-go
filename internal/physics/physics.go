package physix

import (
	"physix/pkg/vector"
	"physix/pkg/rigidbody"
	// "errors"
)

// ApplyForce applies a force to a rigid body.
func ApplyForce(rb *rigidbody.RigidBody, force vector.Vector,dt float64) {
	if(rb.IsMovable){
		// Use Newton's second law: F = ma -> a = F/m
		// rb.Force = rb.Force.Add(force)
		rb.Force = force
		acceleration := rb.Force.Scale(1 / rb.Mass)

		// Update velocity using acceleration and time step
		rb.Velocity = rb.Velocity.Add(acceleration.Scale(dt))

		// Update position using velocity and time step
		rb.Position = rb.Position.Add(rb.Velocity.Scale(dt))
	}
}