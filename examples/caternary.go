package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	physix "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/spring"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

const (
	ScreenWidth   = 800
	ScreenHeight  = 600
	Gravity       = 98
	Friction      = 98
	BallRadius    = 10
	BallMass      = 50
	SegmentCount  = 20
	SegmentLength = 20
	SpringK       = 100.0
	StrongSpringK = 250 // Increase stiffness for end springs
	Damping       = 5
)

type Game struct {
	segments []*rigidbody.RigidBody
	springs  []*spring.Spring
}

func (g *Game) initChain() {
	startX := ScreenWidth/2 - (SegmentCount * SegmentLength / 2)
	startY := 100

	firstFixed := &rigidbody.RigidBody{
		Position:  vector.Vector{X: float64(startX), Y: float64(startY)},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      BallMass,
		Radius:    BallRadius,
		IsMovable: false,
	}
	g.segments = append(g.segments, firstFixed)

	for i := 1; i < SegmentCount; i++ {
		segment := &rigidbody.RigidBody{
			Position:  vector.Vector{X: float64(startX + i*SegmentLength), Y: float64(startY)},
			Velocity:  vector.Vector{X: 0, Y: 0},
			Mass:      BallMass,
			Radius:    BallRadius,
			IsMovable: true,
		}
		g.segments = append(g.segments, segment)

		// Apply strong spring only at the first and last connections
		k := SpringK
		if i == 1 || i == SegmentCount-1 {
			k = StrongSpringK
		}
		g.springs = append(g.springs, spring.NewSpring(g.segments[i-1], segment, k, Damping))
	}

	lastFixed := &rigidbody.RigidBody{
		Position:  vector.Vector{X: float64(startX + SegmentCount*SegmentLength), Y: float64(startY)},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      BallMass,
		Radius:    BallRadius,
		IsMovable: false,
	}
	g.segments = append(g.segments, lastFixed)
	g.springs = append(g.springs, spring.NewSpring(g.segments[len(g.segments)-2], lastFixed, StrongSpringK, Damping))
}

func (g *Game) Update() error {
	for _, segment := range g.segments {
		if segment.IsMovable {
			physix.ApplyForce(segment, vector.Vector{X: 0, Y: Gravity}.Sub(segment.Velocity.Scale(Friction)), 0.1)
		}
	}

	for _, s := range g.springs {
		s.ApplyForce()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, segment := range g.segments {
		ebitenutil.DrawCircle(screen, segment.Position.X, segment.Position.Y, segment.Radius, color.RGBA{0, 255, 0, 255})
	}

	for i := 1; i < len(g.segments); i++ {
		ebitenutil.DrawLine(screen, g.segments[i-1].Position.X, g.segments[i-1].Position.Y, g.segments[i].Position.X, g.segments[i].Position.Y, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Hanging Chain")
	game := &Game{}
	game.initChain()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
