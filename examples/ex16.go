package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/polygon"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"image/color"
)

var (
	ball  *polygon.Polygon
	ball2 *rigidbody.RigidBody
	dt    = 0.01
)

func update() error {
	// Apply a force to simulate gravity
	gravity := vector.Vector{X: 0, Y: 10}
	physix.ApplyForcePolygon(ball, gravity, dt)
	physix.ApplyForce(ball2, gravity, dt)

	if ball.Position.Y < 0 || ball.Position.Y > 400 {
		ball.Velocity.Y = -10
	}

	// If colliding, handle collision response
	if collision.CirclePolygonCollision(ball2, ball) {
		fmt.Println("collision")
		// Resolve collision
		// collision.ResolveCollision(ball, ball2,0.995)
	}

	// Update positions based on velocities
	ball.UpdatePosition()

	return nil
}

func draw(screen *ebiten.Image) {
	// Draw the ball using the github.com/rudransh61/Physix-go engine's position
	// calculateCentroid calculates the centroid of a polygon given its vertices.
	for _, v := range ball.Vertices {
		ebitenutil.DrawRect(screen, v.X, v.Y, 10, 10, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	}
	for _, v1 := range ball.Vertices {
		for _, v2 := range ball.Vertices {
			ebitenutil.DrawLine(screen, v1.X, v1.Y, v2.X, v2.Y, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
		}
	}
	// ball2
	ebitenutil.DrawCircle(screen, ball2.Position.X, ball2.Position.Y, ball2.Radius, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	// fmt.Println(ball.Position.Y)
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Bouncing Ball")

	vertices := []vector.Vector{{X: 250, Y: 50}, {X: 200, Y: 100}, {X: 200, Y: 50}, {X: 350, Y: 200}}

	// Initialize a rigid body with your github.com/rudransh61/Physix-go engine
	ball = polygon.NewPolygon(vertices, 50, true)
	ball2 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 0, Y: 0},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      1.0,
		Shape:     "Circle",
		Radius:    100,
		IsMovable: true,
	}
	ball2.Velocity.X = 100
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
