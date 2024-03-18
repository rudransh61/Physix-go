package rigidbody

import (
	"physix/pkg/vector"
	"math"
)

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
	Torque      float64 
    AngularVelocity float64 
    AngularAcceleration float64 
}

// Rotate any body
func (rb *RigidBody) rotateCoordinates(theta float64) (vector.Vector) {
	//Get coordinates
	x := rb.Position.X 
	y := rb.Position.Y
	
	// Convert theta to radians
	radians := theta * (math.Pi / 180.0)

	// Define the 2D rotation matrix
	rotationMatrix := [2][2]float64{
		{math.Cos(radians), -math.Sin(radians)},
		{math.Sin(radians), math.Cos(radians)},
	}

	// Apply the rotation matrix to the coordinates
	newX := rotationMatrix[0][0]*x + rotationMatrix[0][1]*y
	newY := rotationMatrix[1][0]*x + rotationMatrix[1][1]*y

	return vector.Vector{X:newX, Y:newY}
}


// UpdateRotation updates the rotation of the rigid body based on its angular velocity.
func (rb *RigidBody) UpdateRotation(dt float64) {
    // Update rotation based on angular velocity
    angle := rb.AngularVelocity * dt
    rb.Position.X = math.Cos(angle)*rb.Position.X - math.Sin(angle)*rb.Position.Y
    rb.Position.Y = math.Sin(angle)*rb.Position.X + math.Cos(angle)*rb.Position.Y
}

// ApplyTorque applies a torque to the rigid body.
func (rb *RigidBody) ApplyTorque(torque float64) {
    rb.Torque += torque
}