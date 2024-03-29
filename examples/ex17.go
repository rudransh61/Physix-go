package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
	"math"
	"physix/internal/collision"
	"physix/internal/physics"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
)

var (
	balls      []*rigidbody.RigidBody
	dt         = 0.1
	redBallIdx = 0 // Index of the red ball
	check = 1
)

const (
	Mass   = 2
	Width  = 50
	Height = 50
)

func update() error {
	// Apply gravity and handle wall collisions for all balls
	for _, ball := range balls {
		gravity := vector.Vector{X: 0, Y: 0}
		physix.ApplyForce(ball, gravity, dt)
		checkWallCollision(ball)
	}

	// Check for collisions between balls and resolve them
	for i := 0; i < len(balls); i++ {
		for j := i + 1; j < len(balls); j++ {
			if collision.RectangleCollided(balls[i], balls[j]) {
				if(check<2){
					resolveCollision(balls[i], balls[j])
					// check = 4
				}
				collision.BounceOnCollision(balls[i], balls[j], 1.0)
			}
		}
	}
	check = 4

	// Move the red ball with arrow keys
	redBall := balls[redBallIdx]
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		redBall.Velocity.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		redBall.Velocity.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		redBall.Velocity.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		redBall.Velocity.Y += 1
	}

	return nil
}

func resolveCollision(ball1, ball2 *rigidbody.RigidBody) {
	// Calculate the vector between the centers of the balls
	distance := ball1.Position.Sub(ball2.Position)
	// Calculate the distance between the centers of the balls
	distanceMagnitude := distance.Magnitude()
	// Calculate the minimum distance where the balls stop overlapping
	minimumDistance := ball1.Width + ball2.Width

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

func checkWallCollision(ball *rigidbody.RigidBody) {
	// Bounce off the walls
	// Bounce off the walls
	if ball.Position.X < 100 || ball.Position.X > 600 {
		// ball.Velocity.X = -1*math.Abs(ball.Velocity.X)
		if(ball.Position.X < 100){
			ball.Velocity.X = math.Abs(ball.Velocity.X)
		}
		if(ball.Position.X > 600){
			ball.Velocity.X = -1*math.Abs(ball.Velocity.X)
		}
	}
	if ball.Position.Y < 100 || ball.Position.Y > 600 {
		// ball.Velocity.Y = -1*math.Abs(ball.Velocity.X)
		if(ball.Position.Y < 100){
			ball.Velocity.Y = math.Abs(ball.Velocity.Y)
		}
		if(ball.Position.Y > 600){
			ball.Velocity.Y = -1*math.Abs(ball.Velocity.Y)
		}
	}
}

func draw(screen *ebiten.Image) {
	for i, ball := range balls {
		// Determine color
		var c color.RGBA
		if i == redBallIdx {
			c = color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff} // Red color
		} else {
			c = color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff} // Green color
		}
		// Draw the ball
		ebitenutil.DrawRect(screen, ball.Position.X, ball.Position.Y, ball.Width, ball.Height, c)
	}

	// Draw boundaries
	ebitenutil.DrawRect(screen, 100.0, 100.0, 550, 10, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // Up
	ebitenutil.DrawRect(screen, 100.0, 100.0, 10, 550, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // Left
	ebitenutil.DrawRect(screen, 650.0, 100.0, 10, 550, color.RGBA{R: 0, G: 0xff, B: 0, A: 0})  // Right
	ebitenutil.DrawRect(screen, 100.0, 650.0, 550, 10, color.RGBA{R: 0, G: 0xff, B: 0, A: 0}) // Down
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Bouncing Balls")

	// Initialize rigid bodies (balls)
	n := 5 // Number of balls
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
			Velocity:  vector.Vector{X: 0, Y: 0},
			Mass:      Mass,
			Shape:     "Rectangle",
			Width:     Width,
			Height:    Height,
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
