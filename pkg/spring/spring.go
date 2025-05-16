package spring

import (
	"fmt"

	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	// "math"
)

// Spring struct
type Spring struct {
	BallA, BallB *rigidbody.RigidBody
	RestLength   float64
	Stiffness    float64
	Damping      float64
}

// NewSpring creates a new spring connecting two balls with an optional relaxed length.
func NewSpring(ballA, ballB *rigidbody.RigidBody, stiffness, damping float64, relaxedLength ...float64) *Spring {
	var restLength float64
	if len(relaxedLength) > 0 {
		restLength = relaxedLength[0] // Extract the first value from the slice
	} else {
		restLength = vector.Distance(ballA.Position, ballB.Position)
	}

	return &Spring{BallA: ballA, BallB: ballB, RestLength: restLength, Stiffness: stiffness, Damping: damping}
}

// ApplyForce applies Hooke's Law
func (s *Spring) ApplyForce() {
	dt := 0.1
	delta := s.BallB.Position.Sub(s.BallA.Position)
	fmt.Printf("delta:%v", delta)
	distance := delta.Magnitude()
	fmt.Printf("dist:%v", distance)
	direction := delta.Normalize()
	// Hooke's Law: F = -k(x - L)
	force := direction.Scale(s.Stiffness).Scale(distance - s.RestLength)
	fmt.Print(force)

	// Damping force to stabilize oscillations
	relativeVelocity := vector.ComponentAlong(s.BallB.Velocity.Sub(s.BallA.Velocity), delta)
	dampingForce := relativeVelocity.Scale(s.Damping)

	// Apply forces
	if s.BallA.IsMovable && s.BallA.Mass != 0 {
		acc := force.Add((dampingForce)).Scale(1 / s.BallA.Mass)
		s.BallA.Velocity = s.BallA.Velocity.Add(acc.Scale(dt))
	}
	if s.BallB.IsMovable && s.BallB.Mass != 0 {
		acc := force.Add((dampingForce)).Scale(1 / s.BallB.Mass)
		s.BallB.Velocity = s.BallB.Velocity.Sub(acc.Scale(dt))
	}
}
