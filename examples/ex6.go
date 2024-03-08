package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
	"physix/internal/physics"
	"physix/internal/collision"
	// "math"
	"fmt"
)

var (
	ball *rigidbody.RigidBody
	ball2 *rigidbody.RigidBody
	ball3 *rigidbody.RigidBody
	dt   = 0.1
)

func checkwall(ball *rigidbody.RigidBody) {
	// Bounce off the walls for X direction
	if ball.Position.X-ball.Radius < 100 {
		ball.Position.X = 100 + ball.Radius
		ball.Velocity.X *= -1
	} else if ball.Position.X-ball.Radius > 600 {
		ball.Velocity.X *= -1
	}

	// Bounce off the walls for Y direction
	if ball.Position.Y-ball.Radius < 100 {
		ball.Velocity.Y *= -1
	} else if ball.Position.Y+ball.Radius > 600 {
		ball.Velocity.Y *= -1
	}
}

func update() error {
	physix.ApplyForce(ball, vector.Vector{X: 0, Y: 0}, dt)
	physix.ApplyForce(ball2, vector.Vector{X: 0, Y: 0}, dt)
	physix.ApplyForce(ball3, vector.Vector{X: 0, Y: 0}, dt)

	// //checkwall
	checkwall(ball)
	checkwall(ball2)
	checkwall(ball3)

	if(collision.CircleCollided(ball,ball2)){
		fmt.Println("Collided!")
		collision.BounceOnCollision(ball,ball2,1.0)
	}
	if(collision.CircleCollided(ball3,ball2)){
		fmt.Println("Collided!")
		collision.BounceOnCollision(ball2,ball3,1.0)
	}
	if(collision.CircleCollided(ball3,ball)){
		fmt.Println("Collided!")
		collision.BounceOnCollision(ball,ball3,1.0)
	}
	return nil
}

func draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	ebitenutil.DrawCircle(screen, ball2.Position.X, ball2.Position.Y, ball2.Radius,  color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	ebitenutil.DrawCircle(screen, ball3.Position.X, ball3.Position.Y, ball3.Radius, color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff})
	// ebitenutil.DrawRect(screen, ball2.Position.X, ball2.Position.Y, 70, 70, color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff})
	// ebitenutil.DrawRect(screen, ball3.Position.X, ball3.Position.Y, 70, 70, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
	
	//Boundary
	ebitenutil.DrawRect(screen, 690.0, 100.0 ,10,500, color.RGBA{R:0,G:0xff,B:0,A:0}) // right
	ebitenutil.DrawRect(screen, 90.0, 100.0 ,10,500, color.RGBA{R:0,G:0xff,B:0,A:0}) // left
	ebitenutil.DrawRect(screen, 90.0, 100.0 ,600,10, color.RGBA{R:0,G:0xff,B:0,A:0}) // up
	ebitenutil.DrawRect(screen, 90.0, 600.0 ,600,10, color.RGBA{R:0,G:0xff,B:0,A:0}) // down 
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Bouncing Ball")

	// Initialize a rigid body with your physix engine
	ball = &rigidbody.RigidBody{
		Position: vector.Vector{X: 150, Y: 200},
		Velocity: vector.Vector{X: 0, Y: 0},
		Mass:   rigidbody.Infinite_mass,
		Shape : "Circle",
		Radius : 30,
		IsMovable : false,
	}
	ball2 = &rigidbody.RigidBody{
		Position: vector.Vector{X: 400, Y: 300},
		Velocity: vector.Vector{X: 60, Y: 50},
		Mass:     1,
		Shape : "Circle",
		Radius : 40,
		IsMovable : true,
	}
	ball3 = &rigidbody.RigidBody{
		Position: vector.Vector{X: 400, Y: 400},
		Velocity: vector.Vector{X: -60, Y: 50},
		Mass:     1,
		Shape : "Circle",
		Radius : 50,
		IsMovable : true,
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
