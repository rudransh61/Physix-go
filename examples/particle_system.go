package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"image/color"
	// "math"
	"math/rand"
	// "fmt"
)

var (
	balls []*rigidbody.RigidBody
	down *rigidbody.RigidBody
	right *rigidbody.RigidBody
	up *rigidbody.RigidBody
	left *rigidbody.RigidBody
	dt    = 0.001
	e = 1.0
)

const (
	Mass   = 0.0002
	Shape  = "Circle"
	Radius = 10
)

func update() error {
	gravity := vector.Vector{X: 0, Y: 15}
	for _, ball := range balls {
		physix.ApplyForce(ball, gravity, dt)
		// checkwall(ball)
		if(collision.CircleRectangleCollided(ball, down)){
			// collision.PreventCircleRectangleOverlap(ball, down)
			collision.BounceOnCollision(ball, down, e)
		}
		if(collision.CircleRectangleCollided(ball, right)){
			// collision.PreventCircleRectangleOverlap(ball, right)
			collision.BounceOnCollision(ball, right, e)
		}
		if(collision.CircleRectangleCollided(ball, up)){
			// collision.PreventCircleRectangleOverlap(ball, up)
			collision.BounceOnCollision(ball, up, e)
		}
		if(collision.CircleRectangleCollided(ball, left)){
			// collision.PreventCircleRectangleOverlap(ball, left)
			collision.BounceOnCollision(ball, left, e)
		}
	}
	for steps:=0;steps<10;steps++{
		for i := 0; i < len(balls); i++ {
			for j := i + 1; j < len(balls); j++ {
				if collision.CircleCollided(balls[i], balls[j]) {
					// resolveCollision(balls[i], balls[j])
					collision.PreventCircleOverlap(balls[i], balls[j])
				}
			}
		}
		for _, ball := range balls {
			if(collision.CircleRectangleCollided(ball, down)){
				collision.PreventCircleRectangleOverlap(ball, down)
				// collision.BounceOnCollision(ball, down, e)
			}
			if(collision.CircleRectangleCollided(ball, right)){
				collision.PreventCircleRectangleOverlap(ball, right)
				// collision.BounceOnCollision(ball, right, e)
			}
			if(collision.CircleRectangleCollided(ball, up)){
				collision.PreventCircleRectangleOverlap(ball, up)
				// collision.BounceOnCollision(ball, up, e)
			}
			if(collision.CircleRectangleCollided(ball, left)){
				collision.PreventCircleRectangleOverlap(ball, left)
				// collision.BounceOnCollision(ball, left, e)
			}
		}
	}

	for i := 0; i < len(balls); i++ {
		for j := i + 1; j < len(balls); j++ {
			if collision.CircleCollided(balls[i], balls[j]) {
				// fmt.Println("Collision!")
				collision.BounceOnCollision(balls[i], balls[j], e)
			}
		}
	}

	return nil
}

func draw(screen *ebiten.Image) {
	for _, ball := range balls {
		ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	}

	//Boundary
	ebitenutil.DrawRect(screen, right.Position.X, right.Position.Y, right.Width, right.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0}) // right
	ebitenutil.DrawRect(screen, left.Position.X, left.Position.Y, left.Width, left.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // left
	ebitenutil.DrawRect(screen, up.Position.X, up.Position.Y, up.Width, up.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // up
	ebitenutil.DrawRect(screen, down.Position.X, down.Position.Y, down.Width, down.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // down
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Bouncing Balls")

	// Initialize rigid bodies (balls)
	n := 200 // Change this to the desired number of balls
	initializeBalls(n)

	// Run the game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// initializeBalls initializes n balls with common properties
func initializeBalls(n int) {
	balls = make([]*rigidbody.RigidBody, n)
	for i := 0; i < n; i++ {
		balls[i] = &rigidbody.RigidBody{
			Position:  vector.Vector{X: float64(rand.Intn(200) + 200), Y: float64(rand.Intn(200) + 200)},
			Velocity:  vector.Vector{X: float64(rand.Intn(20)), Y: float64(rand.Intn(20))},
			Mass:      Mass,
			Shape:     Shape,
			Radius:    Radius,
			IsMovable: true,
		}
	}
	down = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 100, Y: 600},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      0,
		Shape:     "Rectangle",
		Width:     510,
		Height:    10,
		IsMovable: false,
	}
	right = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 600, Y: 300},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      0,
		Shape:     "Rectangle",
		Width:     10,
		Height:    600,
		IsMovable: false,
	}
	left = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 100, Y: 300},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      0,
		Shape:     "Rectangle",
		Width:     10,
		Height:    600,
		IsMovable: false,
	}
	up = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 100, Y: 100},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      0,
		Shape:     "Rectangle",
		Width:     510,
		Height:    10,
		IsMovable: false,
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
