package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"image/color"
	// "math"
)

// Global variables
var (
	triangle  []*rigidbody.RigidBody // Triangle vertices
	springs   []*Spring              // Springs connecting triangle vertices
	ball      *rigidbody.RigidBody   // Single falling ball
	dt        = 0.05                 // Time step
)

// Constants
const (
	Mass       = 20
	Shape      = "Circle"
	Radius     = 10
	Stiffness  = 10.0 // Spring stiffness
	Damping    = 2   // Spring damping
	Gravity    = 15    // Gravity force
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

// Physics update function
func update() error {
	gravity := vector.Vector{X: 0, Y: -Gravity-0.5}
	substeps := 1
	for i := 0; i < substeps; i++ {
		// Apply gravity
		for _, v := range triangle {
			physix.ApplyForce(v, gravity, dt)
		}
		physix.ApplyForce(ball,   vector.Vector{X: 0, Y: Gravity+50}, dt)

		// Apply spring forces
		for _, spring := range springs {
			spring.ApplyForce()
		}

		// Handle collisions between ball and triangle vertices
		for _, v := range triangle {
			if collision.CircleCollided(ball, v) {
				resolveCollision(ball, v)
				collision.BounceOnCollision(ball, v, 1.0)
			}
		}
	}

	return nil
}

// Resolve collision between ball and a vertex of the triangle
func resolveCollision(ball1, ball2 *rigidbody.RigidBody) {
	distance := ball1.Position.Sub(ball2.Position)
	distanceMagnitude := distance.Magnitude()
	minimumDistance := ball1.Radius + ball2.Radius

	if distanceMagnitude < minimumDistance {
		moveDirection := distance.Normalize()
		moveAmount := minimumDistance - distanceMagnitude
		moveVector1 := moveDirection.Scale(moveAmount / 2)
		moveVector2 := moveDirection.Scale(-moveAmount / 2)

		ball1.Position = ball1.Position.Add(moveVector1)
		ball2.Position = ball2.Position.Add(moveVector2)
	}
}

// Draw the simulation
func draw(screen *ebiten.Image) {
	// Draw springs (triangle edges)
	for _, spring := range springs {
		ebitenutil.DrawLine(screen,
			spring.BallA.Position.X, spring.BallA.Position.Y,
			spring.BallB.Position.X, spring.BallB.Position.Y,
			color.White)
	}

	// Draw triangle vertices
	for _, v := range triangle {
		ebitenutil.DrawCircle(screen, v.Position.X, v.Position.Y, v.Radius, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	}

	// Draw falling ball
	ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
}

// Initialize a triangle and a ball
func initializeSimulation() {
	triangle = make([]*rigidbody.RigidBody, 3)
	springs = make([]*Spring, 3)

	// Define triangle vertices
	triangle[0] = &rigidbody.RigidBody{Position: vector.Vector{X: 300, Y: 200}, Velocity: vector.Vector{X: 0, Y: 0}, Mass: Mass, Shape: Shape, Radius: Radius, IsMovable: true}
	triangle[1] = &rigidbody.RigidBody{Position: vector.Vector{X: 350, Y: 300}, Velocity: vector.Vector{X: 0, Y: 0}, Mass: Mass, Shape: Shape, Radius: Radius, IsMovable: true}
	triangle[2] = &rigidbody.RigidBody{Position: vector.Vector{X: 250, Y: 300}, Velocity: vector.Vector{X: 0, Y: 0}, Mass: Mass, Shape: Shape, Radius: Radius, IsMovable: true}

	// Create springs for triangle edges
	springs[0] = NewSpring(triangle[0], triangle[1], Stiffness, Damping)
	springs[1] = NewSpring(triangle[1], triangle[2], Stiffness, Damping)
	springs[2] = NewSpring(triangle[2], triangle[0], Stiffness, Damping)

	// Create falling ball
	ball = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 310, Y: 50}, // Initial position above the triangle
		Velocity:  vector.Vector{X: 0, Y: 50},
		Mass:      100,
		Shape:     Shape,
		Radius:    5,
		IsMovable: true,
	}
}

// Game struct
type Game struct{}

func (g *Game) Update() error { return update() }
func (g *Game) Draw(screen *ebiten.Image) { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 800 }

// Main function
func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Soft Body Triangle and Falling Ball")
	initializeSimulation()
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
