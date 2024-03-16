package physix

import (
	"physix/pkg/vector"
	"physix/pkg/rigidbody"
	"physix/pkg/polygon"
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

// ApplyForcePolygon applies force to a polygon.
func ApplyForcePolygon(pg *polygon.Polygon, force vector.Vector, dt float64) {
	if pg.IsMovable {
		// Use Newton's second law: F = ma -> a = F/m
		pg.Force = force
		acceleration := vector.Vector{
			X: pg.Force.X / pg.Mass,
			Y: pg.Force.Y / pg.Mass,
		}

		// Update velocity using acceleration and time step
		pg.Velocity.X += acceleration.X * dt
		pg.Velocity.Y += acceleration.Y * dt

		// Update position using velocity and time step for each vertex
		for i := range pg.Vertices {
			pg.Vertices[i].X += pg.Velocity.X * dt
			pg.Vertices[i].Y += pg.Velocity.Y * dt
		}
	}
}