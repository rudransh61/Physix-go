package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"github.com/rudransh61/Physix-go/pkg/spring"
	"image/color"
	// "math"
)

// Global variables
var (
	triangle  []*rigidbody.RigidBody // Triangle vertices
	springs   []*spring.Spring              // Springs connecting triangle vertices
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
				collision.PreventCircleOverlap(ball, v)
				collision.BounceOnCollision(ball, v, 1.0)
			}
		}
	}

	return nil
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
	springs = make([]*spring.Spring, 3)

	// Define triangle vertices
	triangle[0] = &rigidbody.RigidBody{Position: vector.Vector{X: 300, Y: 200}, Velocity: vector.Vector{X: 0, Y: 0}, Mass: Mass, Shape: Shape, Radius: Radius, IsMovable: true}
	triangle[1] = &rigidbody.RigidBody{Position: vector.Vector{X: 350, Y: 300}, Velocity: vector.Vector{X: 0, Y: 0}, Mass: Mass, Shape: Shape, Radius: Radius, IsMovable: true}
	triangle[2] = &rigidbody.RigidBody{Position: vector.Vector{X: 250, Y: 300}, Velocity: vector.Vector{X: 0, Y: 0}, Mass: Mass, Shape: Shape, Radius: Radius, IsMovable: true}

	// Create springs for triangle edges
	springs[0] = spring.NewSpring(triangle[0], triangle[1], Stiffness, Damping)
	springs[1] = spring.NewSpring(triangle[1], triangle[2], Stiffness, Damping)
	springs[2] = spring.NewSpring(triangle[2], triangle[0], Stiffness, Damping)

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
