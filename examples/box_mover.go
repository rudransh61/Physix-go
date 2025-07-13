package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	physix "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

var (
	boxes     []*rigidbody.RigidBody
	dt        = 0.1
	redBoxIdx = 0 // Index of the red box
	check     = 1
)

const (
	Mass   = 2
	Width  = 50
	Height = 50
)

func update() error {
	// Apply gravity and handle wall collisions for all boxes
	for _, box := range boxes {
		gravity := vector.Vector{X: 0, Y: 0}
		physix.ApplyForce(box, gravity, dt)
		physix.ApplyForce(box, box.Velocity.Scale(-2), dt)
		checkWallCollision(box)
	}

	// Check for collisions between boxes and resolve them
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			if collision.RectangleCollided(boxes[i], boxes[j]) {
				collision.PreventRectangleOverlap(boxes[i], boxes[j])
				collision.BounceOnCollision(boxes[i], boxes[j], 1.0)
				if check < 2 {
					collision.PreventRectangleOverlap(boxes[i], boxes[j])
					check = 4
				}
			}
		}
	}
	check = 4

	// Move the red box with WASD keys
	redbox := boxes[redBoxIdx]
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		redbox.Velocity.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		redbox.Velocity.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		redbox.Velocity.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		redbox.Velocity.Y += 1
	}

	return nil
}

func checkWallCollision(box *rigidbody.RigidBody) {
	// Bounce off the walls
	if box.Position.X < 100 || box.Position.X > 600 {
		// box.Velocity.X = -1*math.Abs(box.Velocity.X)
		if box.Position.X < 100 {
			box.Velocity.X = math.Abs(box.Velocity.X)
		}
		if box.Position.X > 600 {
			box.Velocity.X = -1 * math.Abs(box.Velocity.X)
		}
	}
	if box.Position.Y < 100 || box.Position.Y > 600 {
		// box.Velocity.Y = -1*math.Abs(box.Velocity.X)
		if box.Position.Y < 100 {
			box.Velocity.Y = math.Abs(box.Velocity.Y)
		}
		if box.Position.Y > 600 {
			box.Velocity.Y = -1 * math.Abs(box.Velocity.Y)
		}
	}
}

func draw(screen *ebiten.Image) {
	for i, box := range boxes {
		// Determine color
		var c color.RGBA
		if i == redBoxIdx {
			c = color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff} // Red color
		} else {
			c = color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff} // Green color
		}
		// Draw the box
		ebitenutil.DrawRect(screen, box.Position.X, box.Position.Y, box.Width, box.Height, c)
	}

	// Draw boundaries
	ebitenutil.DrawRect(screen, 100.0, 100.0, 550, 10, color.RGBA{R: 0, G: 0xff, B: 0, A: 0}) // Up
	ebitenutil.DrawRect(screen, 100.0, 100.0, 10, 550, color.RGBA{R: 0, G: 0xff, B: 0, A: 0}) // Left
	ebitenutil.DrawRect(screen, 650.0, 100.0, 10, 550, color.RGBA{R: 0, G: 0xff, B: 0, A: 0}) // Right
	ebitenutil.DrawRect(screen, 100.0, 650.0, 550, 10, color.RGBA{R: 0, G: 0xff, B: 0, A: 0}) // Down
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Bouncing boxes")

	// Initialize rigid bodies (boxes)
	n := 5 // Number of boxes
	initializeboxes(n)

	// Run the game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// initializeboxes initializes n boxes with common properties
func initializeboxes(n int) {
	boxes = make([]*rigidbody.RigidBody, n)
	for i := 0; i < n; i++ {
		boxes[i] = &rigidbody.RigidBody{
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
