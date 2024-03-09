package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"physix/internal/physics"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
	"math/rand"
)

var (
	particles []*rigidbody.RigidBody
	dt        = 0.1
)

const numParticles = 10

func update() error {
	for _, particle := range particles {
		// Update the physix simulation for each particle
		physix.ApplyForce(particle, calculateRandomForce(), dt)
		// physix.UpdateRigidBody(particle, dt)

		// Print the velocity of each particle
		fmt.Println("Velocity:", particle.Velocity.X, particle.Velocity.Y)
	}

	return nil
}

func calculateRandomForce() vector.Vector {
	// Calculate a random force for particle movement
	return vector.Vector{
		X: (rand.Float64() - 0.5) * 10, // Adjust the force magnitude as needed
		Y: (rand.Float64() - 0.5) * 10, // Adjust the force magnitude as needed
	}
}

func draw(screen *ebiten.Image) {
	// Draw each particle using the physix engine's position
	for _, particle := range particles {
		ebitenutil.DrawCircle(screen, particle.Position.X, particle.Position.Y, particle.Radius, color.RGBA{R: 0xff, G: 0, B: 0xff, A: 0xff})
	}
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Particle Simulation")

	// Initialize particles with your physix engine
	particles = make([]*rigidbody.RigidBody, numParticles)
	for i := 0; i < numParticles; i++ {
		particles[i] = &rigidbody.RigidBody{
			Position: vector.Vector{X: rand.Float64() * 400, Y: rand.Float64() * 400},
			Velocity: vector.Vector{X: (rand.Float64() - 0.5) * 10, Y: (rand.Float64() - 0.5) * 10},
			Mass:     1,
			IsMovable: true,
			Shape:    "Circle",
			Radius:   10,
		}
	}

	// Run the game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// Game represents the game state.
type Game struct{}

// Update updates the game logic.
func (g *Game) Update() error {
	return update()
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) {
	draw(screen)
}

// Layout returns the game's screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.ScreenSizeInFullscreen()
}
