package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"physix/internal/physics"
	// "physix/pkg/rigidbody"
	"physix/pkg/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	dt           = 0.1
)

var particles []*physix.Particle

func update() error {
	for _, particle := range particles {
		particle.ApplyForce(vector.Vector{X: 0, Y: 9.8}) // Gravity force
		particle.Update(dt)
	}

	return nil
}

func draw(screen *ebiten.Image) {
	for _, particle := range particles {
		ebitenutil.DrawRect(screen, particle.Position.X, particle.Position.Y, 5, 5, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particle System Example")

	// Create particles
	particles = append(particles, physix.NewParticle(vector.Vector{X: 100, Y: 100}, 1))
	particles = append(particles, physix.NewParticle(vector.Vector{X: 150, Y: 100}, 1))

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
