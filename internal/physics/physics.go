package physix

import (
	"physix/pkg/vector"
	"physix/pkg/rigidbody"
)

// VerletIntegration function for position and velocity updates
func VerletIntegration(current, previous, velocity, force vector.Vector, mass, dt float64) (vector.Vector, vector.Vector) {
    // Update position using Verlet integration
    newPosition := current.Scale(2).Sub(previous).Add(force.Scale(dt * dt / mass))

    // Update velocity using Verlet integration
    newVelocity := newPosition.Sub(previous).Scale(1 / (2 * dt))

    return newPosition, newVelocity
}

// ApplyForce applies a force to a rigid body.
func ApplyForce(rb *rigidbody.RigidBody, force vector.Vector,dt float64) {
	// Use Newton's second law: F = ma -> a = F/m
	// rb.Force = rb.Force.Add(force)
	rb.Force = force
	acceleration := rb.Force.Scale(1 / rb.Mass)

	// Update velocity using acceleration and time step
	rb.Velocity = rb.Velocity.Add(acceleration.Scale(dt))

	// Update position using velocity and time step
	rb.Position = rb.Position.Add(rb.Velocity.Scale(dt))
}