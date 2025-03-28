package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	physix "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

const (
	Gravity    = 0
	BallRadius = 30
	Mass       = 50
)

var (
	ball1 *rigidbody.RigidBody
	ball2 *rigidbody.RigidBody
)

func initBalls() {
	ball1 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 330, Y: 200},
		Velocity:  vector.Vector{X: 5, Y: 5},
		Mass:      Mass,
		Radius:    BallRadius,
		Shape:     "Circle",
		IsMovable: true,
	}

	ball2 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 500, Y: 300},
		Velocity:  vector.Vector{X: -5, Y: -5},
		Mass:      Mass,
		Radius:    BallRadius,
		Shape:     "Circle",
		IsMovable: true,
	}
}

func update() error {
	physix.ApplyForce(ball1, vector.Vector{X: 0, Y: Gravity}, 0.1)
	physix.ApplyForce(ball2, vector.Vector{X: 0, Y: Gravity}, 0.1)

	if collision.CircleCollided(ball1, ball2) {
		collision.PreventCircleOverlap(ball1, ball2)
		collision.BounceOnCollision(ball1, ball2, 1.0)
	}

	return nil
}

func draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, ball1.Position.X, ball1.Position.Y, ball1.Radius, color.RGBA{0, 255, 0, 255})
	ebitenutil.DrawCircle(screen, ball2.Position.X, ball2.Position.Y, ball2.Radius, color.RGBA{0, 0, 255, 255})
}

type Game struct{}

func (g *Game) Update() error                                     { return update() }
func (g *Game) Draw(screen *ebiten.Image)                         { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 600 }

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Oblique Collision")
	initBalls()
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
