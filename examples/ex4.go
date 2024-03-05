package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math"
	"physix/internal/physics"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
	"fmt"
)

var (
	ball *rigidbody.RigidBody
	dt   = 0.1
)

func update() error {
	// Calculate the centripetal force for circular motion
	centripetalForce := calculateCentripetalForce(ball.Position, 50, ball.Mass)

	// Apply the centripetal force to the rigid body
	physix.ApplyForce(ball, centripetalForce)

	// Update the physix simulation
	physix.UpdateRigidBody(ball, centripetalForce, dt)
	fmt.Println(ball.Position.X,ball.Position.Y)

	return nil
}

func calculateCentripetalForce(position vector.Vector, radius, mass float64) vector.Vector {
	// Calculate the centripetal force required for circular motion
	speed := 5.0
	angularVelocity := speed / radius

	// Calculate the new position using parametric equations for circular motion
	angle := angularVelocity * dt
	newX := position.X + radius*math.Cos(angle)
	newY := position.Y + radius*math.Sin(angle)

	// Calculate the centripetal force based on the change in position
	centripetalForceX := (- newX + position.X) / dt * mass
	centripetalForceY := (- newY + position.Y) / dt * mass

	return vector.Vector{X: centripetalForceX, Y: centripetalForceY}
}

func draw(screen *ebiten.Image) {
	// Draw the ball using the physix engine's position
	ebitenutil.DrawRect(screen, ball.Position.X, ball.Position.Y, 20, 20, color.RGBA{R: 0xff, G: 0, B: 0})
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(800, 400)
	ebiten.SetWindowTitle("Circular Motion")

	// Initialize a rigid body with your physix engine
	ball = &rigidbody.RigidBody{
		Position: vector.Vector{X: 400, Y: 200},
		Velocity: vector.Vector{X: 0, Y: 0},
		Mass:     0.0001,
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
