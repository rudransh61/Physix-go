package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"image/color"
	"github.com/rudransh61/Physix-go/pkg/spring"
)

// Constants
const (
	Gravity = 100
	BallRadius = 20
	Mass = 10
)

// Falling ball
var (
	ball = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 400, Y: 100},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      Mass,
		Radius:    BallRadius,
		IsMovable: true,
	}

	ball2 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 300, Y: 100},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      Mass,
		Radius:    BallRadius,
		IsMovable: true,
	}

	pivot = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 500, Y: 100},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      Mass,
		Radius:    BallRadius,
		IsMovable: false,
	}

	springu = spring.NewSpring(ball, pivot, 100, 0.1)
	springuu = spring.NewSpring(ball2, ball, 100, 0.1)
)

// Update physics
func update() error {
	// ball.Velocity = ball.Velocity.Add(vector.Vector{X: 0, Y: Gravity})
	// ball.Position = ball.Position.Add(ball.Velocity)
	physix.ApplyForce(ball, vector.Vector{X: 0, Y: Gravity}, 0.1)
	physix.ApplyForce(ball2, vector.Vector{X: 0, Y: Gravity}, 0.1)
	springu.ApplyForce()
	springuu.ApplyForce()
	return nil
}

// Draw function
func draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.RGBA{0, 255, 0, 255})
	ebitenutil.DrawCircle(screen, ball2.Position.X, ball2.Position.Y, ball2.Radius, color.RGBA{0, 255, 0, 255})
	ebitenutil.DrawCircle(screen, pivot.Position.X, pivot.Position.Y, pivot.Radius, color.RGBA{0, 255, 0, 255})

	ebitenutil.DrawLine(screen,
		ball.Position.X, ball.Position.Y,
		pivot.Position.X, pivot.Position.Y,
		color.White)
	ebitenutil.DrawLine(screen,
		ball.Position.X, ball.Position.Y,
		ball2.Position.X, ball2.Position.Y,
		color.White)
}

// Game struct
type Game struct{}

func (g *Game) Update() error                 { return update() }
func (g *Game) Draw(screen *ebiten.Image)     { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 600 }

// Main function
func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Falling Ball")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
