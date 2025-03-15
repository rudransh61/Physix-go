package spring
import (
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	// "math"
)


// Spring struct
type Spring struct {
	BallA, BallB *rigidbody.RigidBody
	RestLength   float64
	Stiffness    float64
	Damping      float64
}

// NewSpring creates a new spring connecting two balls
func NewSpring(ballA, ballB *rigidbody.RigidBody, stiffness, damping float64) *Spring {
	restLength := ballA.Position.Sub(ballB.Position).Magnitude()
	return &Spring{BallA: ballA, BallB: ballB, RestLength: restLength, Stiffness: stiffness, Damping: damping}
}

// ApplyForce applies Hooke's Law
func (s *Spring) ApplyForce() {
	delta := s.BallB.Position.Sub(s.BallA.Position)
	distance := delta.Magnitude()
	direction := delta.Normalize()

	// Hooke's Law: F = -k(x - L)
	force := direction.Scale(s.Stiffness * (distance - s.RestLength))

	// Damping force to stabilize oscillations
	relativeVelocity := s.BallB.Velocity.Sub(s.BallA.Velocity)
	dampingForce := relativeVelocity.Scale(s.Damping)

	// Apply forces
	s.BallA.Velocity = s.BallA.Velocity.Add(force.Add(dampingForce).Scale(1 / s.BallA.Mass))
	s.BallB.Velocity = s.BallB.Velocity.Sub(force.Add(dampingForce).Scale(1 / s.BallB.Mass))
}