package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"physix/internal/physics"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
)

var (
	ball   *rigidbody.RigidBody
	dt     = 0.1
	points []vector.Vector
	center = vector.Vector{500, 200}
)

func update() error {
	// Calculate the centripetal force for circular motion
	centripetalForce := calculateCentripetalForce(ball.Position, ball.Mass)

	// Apply the centripetal force to the rigid body
	physix.ApplyForce(ball, centripetalForce,dt)

	// Update the physix simulation
	// physix.UpdateRigidBody(ball, dt)

	// Add the current position to the list of points
	points = append(points, ball.Position)

	fmt.Println(ball.Position.X, ball.Position.Y)

	return nil
}

func calculateCentripetalForce(position vector.Vector, mass float64) vector.Vector {
	// Calculate the centripetal force required for circular motion
	speed := 20.0
	// F = -mv^2/R^2 * <R>
	rad := position.Sub(center)
	// radius := position.Magnitude()
	radius := 200.0
	Force := rad.Scale(-mass*speed*speed/(radius*radius))

	return Force
}

func draw(screen *ebiten.Image) {
	// Draw the ball using the physix engine's position
	ebitenutil.DrawRect(screen, center.X, center.Y, 20, 20, color.RGBA{R: 0, G: 0xff, B: 0})
	ebitenutil.DrawRect(screen, ball.Position.X, ball.Position.Y, 20, 20, color.RGBA{R: 0xff, G: 0, B: 0})

	// Draw a small white dot at each traced point
	for _, point := range points {
		ebitenutil.DrawRect(screen, point.X, point.Y, 2, 2, color.White)
	}
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Circular Motion")

	// Initialize a rigid body with your physix engine
	ball = &rigidbody.RigidBody{
		Position: vector.Vector{X: 500, Y: 400},
		Velocity: vector.Vector{X: 20, Y: 0},
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
