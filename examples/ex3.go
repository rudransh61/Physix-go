package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"github.com/rudransh61/Physix-go/internal/physics"
)

var (
	ball *rigidbody.RigidBody
	dt   = 0.1
)

func update() error {
	// Apply a force to simulate gravity
	gravity := vector.Vector{X: 0, Y: 2}
	physix.ApplyForce(ball, gravity, dt)

	// Bounce off the walls
	if ball.Position.X < 0 || ball.Position.X > 400 {
		ball.Velocity.X *= -1
	}
	if ball.Position.Y < 0 || ball.Position.Y > 400 {
		ball.Velocity.Y *= -1
	}

	return nil
}

func draw(screen *ebiten.Image) {
	// Draw the ball using the github.com/rudransh61/Physix-go engine's position
	ebitenutil.DrawRect(screen, ball.Position.X, ball.Position.Y, 20, 20, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Bouncing Ball")

	// Initialize a rigid body with your github.com/rudransh61/Physix-go engine
	ball = &rigidbody.RigidBody{
		Position: vector.Vector{X: 100, Y: 200},
		Velocity: vector.Vector{X: 50, Y: -50},
		Mass:     1,
		IsMovable : true,
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
