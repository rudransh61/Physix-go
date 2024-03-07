package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
	"physix/internal/physics"
	"physix/internal/collision"
	"fmt"
)

var (
	ball *rigidbody.RigidBody
	ball2 *rigidbody.RigidBody
	ball3 *rigidbody.RigidBody
	dt   = 0.1
)

func checkwall(ball *rigidbody.RigidBody){
	// Bounce off the walls
	if ball.Position.X < 100 || ball.Position.X > 600 {
		ball.Velocity.X *= -1
	}
	if ball.Position.Y < 100 || ball.Position.Y > 600 {
		ball.Velocity.Y *= -1
	}
}

func CheckBall(rect1, rect2 *rigidbody.RigidBody) bool {
	left1, top1, right1, bottom1 := rect1.Position.X, rect1.Position.Y, rect1.Position.X+rect1.Width, rect1.Position.Y+rect1.Height
	left2, top2, right2, bottom2 := rect2.Position.X, rect2.Position.Y, rect2.Position.X+rect2.Width, rect2.Position.Y+rect2.Height

	return right1 > left2 && left1 < right2 && bottom1 > top2 && top1 < bottom2
}

func update() error {
	physix.ApplyForce(ball, vector.Vector{X: 0, Y: 0}, dt)
	physix.ApplyForce(ball2, vector.Vector{X: 0, Y: 0}, dt)
	physix.ApplyForce(ball3, vector.Vector{X: 0, Y: 0}, dt)

	//checkwall
	checkwall(ball)
	checkwall(ball2)
	checkwall(ball3)

	if(collision.RectangleCollided(ball,ball2,70,70,70,70)){
		fmt.Println("Collided!")
	}
	if(collision.RectangleCollided(ball3,ball2,70,70,70,70)){
		fmt.Println("Collided!")
	}
	if(collision.RectangleCollided(ball3,ball,70,70,70,70)){
		fmt.Println("Collided!")
	}
	if(CheckBall(ball,ball2)){
		collision.BounceOnCollision(ball,ball2,1.0)
	}
	if(CheckBall(ball,ball3)){
		collision.BounceOnCollision(ball,ball3,1.0)
	}
	if(CheckBall(ball3,ball2)){
		collision.BounceOnCollision(ball3,ball2,1.0)
	}
	return nil
}

func draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, ball.Position.X, ball.Position.Y, 70, 70, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	ebitenutil.DrawRect(screen, ball2.Position.X, ball2.Position.Y, 70, 70, color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff})
	ebitenutil.DrawRect(screen, ball3.Position.X, ball3.Position.Y, 70, 70, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	ebitenutil.DrawRect(screen, 690.0, 100.0 ,10,600, color.RGBA{R:0,G:0xff,B:0,A:0})
	ebitenutil.DrawRect(screen, 90.0, 100.0 ,10,600, color.RGBA{R:0,G:0xff,B:0,A:0})
	ebitenutil.DrawRect(screen, 90.0, 100.0 ,600,10, color.RGBA{R:0,G:0xff,B:0,A:0})
	ebitenutil.DrawRect(screen, 90.0, 690.0 ,600,10, color.RGBA{R:0,G:0xff,B:0,A:0})
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Bouncing Ball")

	// Initialize a rigid body with your physix engine
	ball = &rigidbody.RigidBody{
		Position: vector.Vector{X: 100, Y: 200},
		Velocity: vector.Vector{X: 50, Y: -50},
		Mass:     1,
		Shape : "Rectangle",
		Width : 70,
		Height : 70,
	}
	ball2 = &rigidbody.RigidBody{
		Position: vector.Vector{X: 400, Y: 300},
		Velocity: vector.Vector{X: 60, Y: 50},
		Mass:     1,
		Shape : "Rectangle",
		Width : 70,
		Height : 70,
	}
	ball3 = &rigidbody.RigidBody{
		Position: vector.Vector{X: 400, Y: 400},
		Velocity: vector.Vector{X: -60, Y: 50},
		Mass:     5,
		Shape : "Rectangle",
		Width : 70,
		Height : 70,
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
