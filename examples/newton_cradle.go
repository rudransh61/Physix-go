package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	physix "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/spring"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

const (
	Gravity    = 100
	BallRadius = 10
	Mass       = 10
	Spacing    = 20
	RestLength = 80
	SpringK    = 0.1
)

var (
	balls   []*rigidbody.RigidBody
	springs []*spring.Spring
	pivots  []*rigidbody.RigidBody
)

func initCradle() {
	numBalls := 5
	startX := 300
	startY := 100

	for i := 0; i < numBalls; i++ {
		pivot := &rigidbody.RigidBody{
			Position:  vector.Vector{X: float64(startX + i*Spacing), Y: float64(startY)},
			Velocity:  vector.Vector{X: 0, Y: 0},
			Mass:      Mass,
			Radius:    BallRadius,
			IsMovable: false,
		}
		pivots = append(pivots, pivot)

		ball := &rigidbody.RigidBody{
			Position:  vector.Vector{X: float64(startX + i*Spacing), Y: float64(startY + RestLength)},
			Velocity:  vector.Vector{X: 0, Y: 0},
			Mass:      Mass,
			Radius:    BallRadius,
			Shape:     "Circle",
			IsMovable: true,
		}
		balls = append(balls, ball)

		springs = append(springs, spring.NewSpring(ball, pivot, RestLength, SpringK))
	}

	// Give the first ball an initial displacement
	balls[0].Velocity = vector.Vector{X: -50, Y: 0}
}

func update() error {
	for _, ball := range balls {
		physix.ApplyForce(ball, vector.Vector{X: 0, Y: Gravity}, 0.1)
	}

	for i := 0; i < len(balls); i++ {
		for j := 0; j < len(balls); j++ {
			if collision.CircleCollided(balls[i], balls[j]) {
				collision.PreventCircleOverlap(balls[i], balls[j])
				collision.BounceOnCollision(balls[i], balls[j], 1.0)
			}
		}
	}

	for _, s := range springs {
		s.ApplyForce()
	}

	return nil
}

func draw(screen *ebiten.Image) {
	for _, ball := range balls {
		ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.RGBA{0, 255, 0, 255})
	}

	for _, pivot := range pivots {
		ebitenutil.DrawCircle(screen, pivot.Position.X, pivot.Position.Y, pivot.Radius, color.RGBA{255, 0, 0, 255})
	}

	for i, ball := range balls {
		ebitenutil.DrawLine(screen, ball.Position.X, ball.Position.Y, pivots[i].Position.X, pivots[i].Position.Y, color.White)
	}
}

type Game struct{}

func (g *Game) Update() error                                     { return update() }
func (g *Game) Draw(screen *ebiten.Image)                         { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 600 }

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Newton's Cradle")
	initCradle()
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
