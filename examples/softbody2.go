package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	physix "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/spring"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

// Global variables
var (
	square    []*rigidbody.RigidBody // Four particles of the square
	springs   []*spring.Spring       // Springs connecting the square particles
	platform1 *rigidbody.RigidBody
	platform2 *rigidbody.RigidBody
	dt        = 0.05
)

// Constants
const (
	Gravity   = 9.8
	Stiffness = 100.0 // Spring stiffness
	Damping   = 5.0   // Spring damping
)

// Physics update function
func update() error {
	gravity := vector.Vector{X: 0, Y: Gravity}

	// Apply gravity to each particle
	for _, v := range square {
		physix.ApplyForce(v, gravity.Scale(v.Mass), dt)
		if collision.CircleRectangleCollided(v, platform1) {
			collision.PreventCircleRectangleOverlap(v, platform1)
			collision.BounceOnCollision(v, platform1, 1.0)
		}
		if collision.CircleRectangleCollided(v, platform2) {
			collision.PreventCircleRectangleOverlap(v, platform2)
			collision.BounceOnCollision(v, platform2, 1.0)
		}
	}

	// Apply spring forces
	for _, s := range springs {
		s.ApplyForce()
	}

	return nil
}

// Draw function
func draw(screen *ebiten.Image) {
	// Draw springs (edges of the square)
	for _, s := range springs {
		ebitenutil.DrawLine(screen,
			s.BallA.Position.X, s.BallA.Position.Y,
			s.BallB.Position.X, s.BallB.Position.Y,
			color.White)
	}

	// Draw particles of the square
	for _, v := range square {
		ebitenutil.DrawCircle(screen, v.Position.X, v.Position.Y, v.Radius, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	}

	// Draw platforms
	ebitenutil.DrawRect(screen, platform1.Position.X, platform1.Position.Y, platform1.Width, platform1.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	ebitenutil.DrawRect(screen, platform2.Position.X, platform2.Position.Y, platform2.Width, platform2.Height, color.RGBA{R: 0, G: 0xff, B: 0xff, A: 0xff})
}

// Initialize simulation
func initializeSimulation() {
	square = make([]*rigidbody.RigidBody, 4)
	springs = make([]*spring.Spring, 6)

	// Define square vertices (particles)
	square[0] = &rigidbody.RigidBody{Position: vector.Vector{X: 200, Y: 100}, Shape: "Circle", Mass: 50, Radius: 5, IsMovable: true}
	square[1] = &rigidbody.RigidBody{Position: vector.Vector{X: 250, Y: 100}, Shape: "Circle", Mass: 50, Radius: 5, IsMovable: true}
	square[2] = &rigidbody.RigidBody{Position: vector.Vector{X: 200, Y: 150}, Shape: "Circle", Mass: 50, Radius: 5, IsMovable: true}
	square[3] = &rigidbody.RigidBody{Position: vector.Vector{X: 250, Y: 150}, Shape: "Circle", Mass: 50, Radius: 5, IsMovable: true}

	// Create springs between square vertices
	springs[0] = spring.NewSpring(square[0], square[1], Stiffness, Damping)
	springs[1] = spring.NewSpring(square[1], square[3], Stiffness, Damping)
	springs[2] = spring.NewSpring(square[3], square[2], Stiffness, Damping)
	springs[3] = spring.NewSpring(square[2], square[0], Stiffness, Damping)
	springs[4] = spring.NewSpring(square[3], square[0], Stiffness, Damping)
	springs[5] = spring.NewSpring(square[2], square[1], Stiffness, Damping)

	// Create platform1 (lower one)
	platform1 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 100, Y: 300},
		Mass:      rigidbody.Infinite_mass,
		Shape:     "Rectangle",
		Width:     400,
		Height:    20,
		IsMovable: false,
	}

	// Create platform2 (higher one)
	platform2 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 245, Y: 200},
		Mass:      rigidbody.Infinite_mass,
		Shape:     "Rectangle",
		Width:     400,
		Height:    5,
		IsMovable: false,
	}
}

// Game struct
type Game struct{}

func (g *Game) Update() error                                     { return update() }
func (g *Game) Draw(screen *ebiten.Image)                         { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 500, 400 }

// Main function
func main() {
	ebiten.SetWindowSize(500, 400)
	ebiten.SetWindowTitle("Soft Body Square with Two Platforms")
	initializeSimulation()
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
