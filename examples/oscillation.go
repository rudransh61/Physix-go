package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	physix "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/spring"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

const (
	Gravity          = 10
	BallRadius       = 20
	Mass             = 10
	SpringRestLength = 100
	SpringStiffness  = 100
)

var (
	mass = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 400, Y: 110},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      Mass,
		Radius:    BallRadius,
		IsMovable: true,
		Shape:     "Circle",
	}

	anchor = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 400, Y: 100},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      0,
		Radius:    BallRadius,
		IsMovable: false,
	}

	springSim = spring.NewSpring(mass, anchor, SpringStiffness, 0)
)

func update() error {
	physix.ApplyForce(mass, vector.Vector{X: 0, Y: Gravity}, 0.1)
	springSim.ApplyForce()
	fmt.Printf("Mass Position: %v\n", mass.Position)
	return nil
}

func draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, mass.Position.X, mass.Position.Y, mass.Radius, color.RGBA{0, 255, 0, 255})
	ebitenutil.DrawCircle(screen, anchor.Position.X, anchor.Position.Y, anchor.Radius, color.RGBA{255, 0, 0, 255})
	ebitenutil.DrawLine(screen,
		mass.Position.X, mass.Position.Y,
		anchor.Position.X, anchor.Position.Y,
		color.White)
}

type Game struct{}

func (g *Game) Update() error                                     { return update() }
func (g *Game) Draw(screen *ebiten.Image)                         { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 600 }

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hanging Mass with Spring")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
