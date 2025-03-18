package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"github.com/rudransh61/Physix-go/pkg/spring"
	"image/color"
	"math"
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
	Gravity   = 98
	Stiffness = 100.0 // Spring stiffness
	Damping   = 20  // Spring damping
)

// Physics update function
func update() error {
	gravity := vector.Vector{X: 0, Y: Gravity}

	// Apply gravity to each particle
	for _, v := range square {
		physix.ApplyForce(v, gravity.Scale(v.Mass), dt)
	}

	// Apply spring forces
	for _, s := range springs {
		s.ApplyForce()
	}

	// Check collision for each particle against both platforms
	// Check collision for each particle against both platforms
for _, v := range square {
    // Collision with platform1
    if v.Position.Y > platform1.Position.Y && 
       v.Position.X > platform1.Position.X && 
       v.Position.X < platform1.Position.X+platform1.Width {
        v.Velocity.Y = -math.Abs(v.Velocity.Y)*2
    }

    // Collision with platform2
    if v.Position.Y > platform2.Position.Y && 
       v.Position.Y < platform2.Position.Y+platform2.Height &&
       v.Position.X > platform2.Position.X && 
       v.Position.X < platform2.Position.X+platform2.Width {
        v.Velocity.Y = -v.Velocity.Y
    }
	// return  nil
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

	// Draw platform1
	ebitenutil.DrawRect(screen, platform1.Position.X, platform1.Position.Y, platform1.Width, platform1.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})

	// Draw platform2
	ebitenutil.DrawRect(screen, platform2.Position.X, platform2.Position.Y, platform2.Width, platform2.Height, color.RGBA{R: 0, G: 0xff, B: 0xff, A: 0xff}) // Different color
}

// Initialize simulation
func initializeSimulation() {
	square = make([]*rigidbody.RigidBody, 4)
	springs = make([]*spring.Spring, 6)

	// Define square vertices (particles)
	square[0] = &rigidbody.RigidBody{Position: vector.Vector{X: 200, Y: 100}, Mass: 50, Radius: 5, IsMovable: true}
	square[1] = &rigidbody.RigidBody{Position: vector.Vector{X: 250, Y: 100}, Mass: 50, Radius: 5, IsMovable: true}
	square[2] = &rigidbody.RigidBody{Position: vector.Vector{X: 200, Y: 150}, Mass: 50, Radius: 5, IsMovable: true}
	square[3] = &rigidbody.RigidBody{Position: vector.Vector{X: 250, Y: 150}, Mass: 50, Radius: 5, IsMovable: true}

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

func (g *Game) Update() error { return update() }
func (g *Game) Draw(screen *ebiten.Image) { draw(screen) }
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
