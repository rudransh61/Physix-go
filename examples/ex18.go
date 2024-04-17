package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image/color"
    "math"
    "github.com/rudransh61/Physix-go/internal/collision"
    "github.com/rudransh61/Physix-go/internal/physics"
    "github.com/rudransh61/Physix-go/pkg/polygon"
    "github.com/rudransh61/Physix-go/pkg/rigidbody"
    "github.com/rudransh61/Physix-go/pkg/vector"
    "fmt"
)

var (
    ball    *polygon.Polygon
    bat     *polygon.Polygon
    dt      = 0.01 // Adjusted time step
    batAngle = 0.0   // Initial angle of the bat
)

func update() error {
    // Apply a force to simulate gravity
    gravity := vector.Vector{X: 0, Y: 1000}
    github.com/rudransh61/Physix-go.ApplyForcePolygon(ball, gravity, dt)

    // Rotate the bat
    batAngle += 0.01 // Adjust the rotation speed as needed

    // If colliding, handle collision response
    if collision.PolygonCollision(ball, bat) {
        fmt.Println("collision")
        // Resolve collision
        collision.ResolveCollision(ball, bat, 1.0, 0.01) // Adjusted torque factor
		ball.UpdatePosition()
    }

    // Update positions based on velocities

    return nil
}

func draw(screen *ebiten.Image) {
    // Draw the ball using the github.com/rudransh61/Physix-go engine's position
    for _, v := range ball.Vertices {
        ebitenutil.DrawRect(screen, v.X, v.Y, 10, 10, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
    }

    // Draw the bat using the github.com/rudransh61/Physix-go engine's position and rotation
    for _, v := range bat.Vertices {
        // Rotate each vertex of the bat
        rotatedVertex := rotateVertex(v, batAngle, bat.Position)
        ebitenutil.DrawRect(screen, rotatedVertex.X, rotatedVertex.Y, 10, 10, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
    }

    // Draw lines for the bat
    for i := 0; i < len(bat.Vertices); i++ {
        v1 := rotateVertex(bat.Vertices[i], batAngle, bat.Position)
        v2 := rotateVertex(bat.Vertices[(i+1)%len(bat.Vertices)], batAngle, bat.Position)
        ebitenutil.DrawLine(screen, v1.X, v1.Y, v2.X, v2.Y, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff})
    }
}

func rotateVertex(v vector.Vector, angle float64, pivot vector.Vector) vector.Vector {
    sinTheta := math.Sin(angle)
    cosTheta := math.Cos(angle)
    translatedX := v.X - pivot.X
    translatedY := v.Y - pivot.Y
    rotatedX := translatedX*cosTheta - translatedY*sinTheta
    rotatedY := translatedX*sinTheta + translatedY*cosTheta
    return vector.Vector{X: rotatedX + pivot.X, Y: rotatedY + pivot.Y}
}

func main() {
    // Set up the window
    ebiten.SetWindowSize(400, 400)
    ebiten.SetWindowTitle("Falling Ball")

    // Define vertices for the long bat
    batVertices := []vector.Vector{{X: 150, Y: 350}, {X: 350, Y: 350}, {X: 350, Y: 400}, {X: 150, Y: 400}} // Make the bat longer

    // Define vertices for the small square ball
    ballVertices := []vector.Vector{{X: 200, Y: 200}, {X: 210, Y: 200}, {X: 210, Y: 210}, {X: 200, Y: 210}} // Make the ball square and small

    // Initialize a rigid body with your github.com/rudransh61/Physix-go engine
    ball = polygon.NewPolygon(ballVertices, 50, true)

    // Initialize the bat
    bat = polygon.NewPolygon(batVertices, rigidbody.Infinite_mass, false) // Set mass to 0 since bat is not movable

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
