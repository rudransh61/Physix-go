package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"physix/internal/physics"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
	// "log"
)

var (
	ball *rigidbody.RigidBody
	dt   = 0.1
)

func update() error {
	ball.ApplyTorque(1)
	// Update the physix simulation
	physix.ApplyForce(ball, ball.Force, dt)
	// physix.UpdateRotation(ball,dt)


	return nil
}

func draw(screen *ebiten.Image) {
	// Draw the rectangle using the physix engine's position
	ebitenutil.DrawRect(screen, ball.Position.X, ball.Position.Y, 50, 50, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Projectile Motion")

	// Initialize a rigid body with your physix engine
	ball = &rigidbody.RigidBody{
		Position: vector.Vector{X: 100, Y: 400},
		Velocity: vector.Vector{X: 30, Y: -20},
		Mass:     1,
		Force : vector.Vector{X: 0, Y: 0},
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
