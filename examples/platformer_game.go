package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	// "math"
)
// platformer
var (
	ball      *rigidbody.RigidBody
	platform1 *rigidbody.RigidBody
	platform2 *rigidbody.RigidBody
	platform3 *rigidbody.RigidBody
	dt        = 0.1
	jumped    = false
	camX, camY float64
)
func update() error {
	// Camera movement
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		camX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		camX += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		camY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		camY += 5
	}

	// Handle input for the platform movement
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ball.Position.X += -15 * dt
		camX -= 16 * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ball.Position.X += 15 * dt
		camX += 14 * dt
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if ball.Velocity.Y < 2 && ball.Velocity.Y > -2 {
			ball.Velocity.Y -= 50
		}
	}

	if ball.Velocity.Y < 2 {
		jumped = false
	}

	// Update the physics simulation
	camY += ball.Velocity.Y * dt * 0.95
	physix.ApplyForce(ball, ball.Force, dt)
	physix.ApplyForce(platform1, ball.Force, dt)
	physix.ApplyForce(platform2, ball.Force, dt)
	physix.ApplyForce(platform3, ball.Force, dt)
	ball.Force.Y = 5

	// Check for collision between ball and platforms
	if collision.RectangleCollided(ball, platform1) {
		collision.PreventRectangleOverlap(ball, platform1)
		collision.BounceOnCollision(ball, platform1, 0.0)
	}
	if collision.RectangleCollided(ball, platform2) {
		collision.PreventRectangleOverlap(ball, platform2)
		collision.BounceOnCollision(ball, platform1, 0.0)
	}
	if collision.RectangleCollided(ball, platform3) {
		collision.PreventRectangleOverlap(ball, platform3)
		collision.BounceOnCollision(ball, platform1, 0.0)
	}

	return nil
}

func draw(screen *ebiten.Image) {
	// Apply camera transformation
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-camX, -camY)

	// Draw ball and platforms with camera offset
	ebitenutil.DrawRect(screen, ball.Position.X-camX, ball.Position.Y-camY, ball.Width, ball.Height, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	ebitenutil.DrawRect(screen, platform1.Position.X-camX, platform1.Position.Y-camY, platform1.Width, platform1.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	ebitenutil.DrawRect(screen, platform2.Position.X-camX, platform2.Position.Y-camY, platform2.Width, platform2.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	ebitenutil.DrawRect(screen, platform3.Position.X-camX, platform3.Position.Y-camY, platform3.Width, platform3.Height, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
}

func main() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Bouncing Ball - feat Gravity")

	// Initialize objects
	ball = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 400, Y: 100},
		Velocity:  vector.Vector{X: 0, Y: 2},
		Mass:      1,
		Force:     vector.Vector{X: 0, Y: 5},
		IsMovable: true,
		Shape:     "Rectangle",
		Width:     25,
		Height:    45,
	}

	platform1 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 100, Y: 600},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      rigidbody.Infinite_mass,
		IsMovable: false,
		Shape:     "Rectangle",
		Width:     200,
		Height:    50,
	}

	platform2 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 400, Y: 450},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      rigidbody.Infinite_mass,
		IsMovable: false,
		Shape:     "Rectangle",
		Width:     200,
		Height:    50,
	}

	platform3 = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 700, Y: 300},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      rigidbody.Infinite_mass,
		IsMovable: false,
		Shape:     "Rectangle",
		Width:     200,
		Height:    50,
	}

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
	return 500, 500
}
