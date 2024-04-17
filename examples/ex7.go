package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"github.com/rudransh61/Physix-go/internal/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"github.com/rudransh61/Physix-go/internal/collision"
	"fmt"
)

var (
	ball *rigidbody.RigidBody
	platform *rigidbody.RigidBody
	dt   = 0.1
)

func update() error {
	// Update the github.com/rudransh61/Physix-go simulation
	github.com/rudransh61/Physix-go.ApplyForce(ball, ball.Force, dt)
	github.com/rudransh61/Physix-go.ApplyForce(platform, ball.Force, dt)

	if(collision.RectangleCollided(ball,platform)){
		fmt.Println("Bounced!")
		collision.BounceOnCollision(ball,platform,1.0)
	}

	return nil
}

func draw(screen *ebiten.Image) {
	// Draw the rectangle using the github.com/rudransh61/Physix-go engine's position
	ebitenutil.DrawRect(screen, ball.Position.X, ball.Position.Y, ball.Width, ball.Height, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	ebitenutil.DrawRect(screen, platform.Position.X, platform.Position.Y, platform.Width, platform.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Bouncing Ball - feat Gravity")

	// Initialize a rigid body with your github.com/rudransh61/Physix-go engine
	ball = &rigidbody.RigidBody{
		Position: vector.Vector{X: 400, Y: 100},
		Velocity: vector.Vector{X: 0, Y: 2},
		Mass:     1,
		Force : vector.Vector{X: 0, Y: 5},
		IsMovable : true,
		Shape: "Rectangle",
		Width : 50,
		Height : 50,
	}

	platform = &rigidbody.RigidBody{
		Position : vector.Vector{X:100 , Y:600},
		Velocity : vector.Vector{X:0,Y:0},
		Mass : rigidbody.Infinite_mass,
		IsMovable: false,
		Shape: "Rectangle",
		Width : 1000,
		Height : 50,
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
