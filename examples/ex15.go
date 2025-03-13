package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"image/color"
	"math"
	"math/rand"
	// "fmt"
)

var (
	balls []*rigidbody.RigidBody
	dt    = 0.05
)

const (
	Mass   = 2
	Shape  = "Circle"
	Radius = 10
)

func checkwall(ball *rigidbody.RigidBody) {
	// Bounce off the walls
	if ball.Position.X < 100+ball.Radius || ball.Position.X > 600-ball.Radius {
		if ball.Position.X < 100+ball.Radius {
			ball.Velocity.X = math.Abs(ball.Velocity.X * 0.7)
			ball.Position.X = 100 + ball.Radius
		}
		if ball.Position.X > 600-ball.Radius {
			ball.Velocity.X = -0.7 * math.Abs(ball.Velocity.X)
			ball.Position.X = 600 - ball.Radius
		}
	}
	if ball.Position.Y < 100+ball.Radius || ball.Position.Y > 600-ball.Radius {
		if ball.Position.Y < 100+ball.Radius {
			ball.Velocity.Y = math.Abs(ball.Velocity.Y * 0.7)
			ball.Position.Y = 100 + ball.Radius
		}
		if ball.Position.Y > 600-ball.Radius {
			ball.Velocity.Y = -0.7 * math.Abs(ball.Velocity.Y)
			ball.Position.Y = 600 - ball.Radius
		}
	}
}

func update() error {
	gravity := vector.Vector{X: 0, Y: 15}
	substeps := 10;
	for i:=0;i<substeps;i++ {
		for _, ball := range balls {
			physix.ApplyForce(ball, gravity, dt)
			checkwall(ball)
			checkwall(ball)
		}
	
	
	
	for i := 0; i < len(balls); i++ {
		for j := i + 1; j < len(balls); j++ {
			if collision.CircleCollided(balls[i], balls[j]) {
				// fmt.Println("Collision!")
				resolveCollision(balls[i], balls[j])
				
				// collision.BounceOnCollision(balls[i], balls[j], 0.7)
				collision.BounceOnCollision(balls[i], balls[j], 0.7)
			}
		}
	}
}

	return nil
}

func resolveCollision(ball1, ball2 *rigidbody.RigidBody) {
	// Calculate the vector between the centers of the balls
	distance := ball1.Position.Sub(ball2.Position)
	// Calculate the distance between the centers of the balls
	distanceMagnitude := distance.Magnitude()
	// Calculate the minimum distance where the balls stop overlapping
	minimumDistance := ball1.Radius + ball2.Radius

	// Check if the balls are overlapping
	if distanceMagnitude < minimumDistance {
		// Calculate the direction to move the balls apart
		moveDirection := distance.Normalize()
		// Calculate the amount by which to move the balls apart
		moveAmount := minimumDistance - distanceMagnitude
		// Calculate the movement vectors for each ball
		moveVector1 := moveDirection.Scale(moveAmount / 2)
		moveVector2 := moveDirection.Scale(-moveAmount / 2)
		// Move the balls apart
		ball1.Position = ball1.Position.Add(moveVector1)
		ball2.Position = ball2.Position.Add(moveVector2)
	}
}

func draw(screen *ebiten.Image) {
	for _, ball := range balls {
		ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
	}

	//Boundary
	ebitenutil.DrawRect(screen, 600.0, 100.0, 10, 500, color.RGBA{R: 0, G: 0xff, B: 0, A: 0}) // right
	ebitenutil.DrawRect(screen, 90.0, 100.0, 10, 500, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // left
	ebitenutil.DrawRect(screen, 90.0, 100.0, 510, 10, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // up
	ebitenutil.DrawRect(screen, 90.0, 600.0, 510, 10, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // down
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
