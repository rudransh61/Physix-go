package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"physix/pkg/matrices" // Update this with your actual module path
)

type Game struct{}

var (
	cubeVertices = [][]float64{
		{-1.0, -1.0, -1.0},
		{-1.0, -1.0, 1.0},
		{-1.0, 1.0, -1.0},
		{-1.0, 1.0, 1.0},
		{1.0, -1.0, -1.0},
		{1.0, -1.0, 1.0},
		{1.0, 1.0, -1.0},
		{1.0, 1.0, 1.0},
	}

	cubeEdges = [][]int{
		{0.0, 1.0}, {1.0, 3.0}, {3.0, 2.0}, {2.0, 0.0},
		{4.0, 5.0}, {5.0, 7.0}, {7.0, 6.0}, {6.0, 4.0},
		{0.0, 4.0}, {1.0, 5.0}, {2.0, 6.0}, {3.0, 7.0},
	}

	screenWidth  = 640
	screenHeight = 480

	rotationAngles = []float64{0, 0, 0}
)

func (g *Game) Update() error {
	// Clear the screen
	// ebiten.Screen.Fill(color.RGBA{0, 0, 0, 255})

	// Rotate the cube around x, y, and z axes
	rotationMatrix := matrices.Identity3
	rotationMatrix = matrices.Multiply(rotationMatrix, matrices.RotationZ(rotationAngles[2]))
	rotationMatrix = matrices.Multiply(rotationMatrix, matrices.RotationY(rotationAngles[1]))
	rotationMatrix = matrices.Multiply(rotationMatrix, matrices.RotationX(rotationAngles[0]))

	// Project and draw each edge of the cube
	for _, edge := range cubeEdges {
		// Convert vertices to matrices
		vertex1Matrix := [][]float64{cubeVertices[edge[0]]}
		vertex2Matrix := [][]float64{cubeVertices[edge[1]]}

		// Rotate vertices
		vertex1 := matrices.Multiply(vertex1Matrix, rotationMatrix)[0]
		vertex2 := matrices.Multiply(vertex2Matrix, rotationMatrix)[0]

		// Draw the line
		drawLine(vertex1[0], vertex1[1], vertex2[0], vertex2[1], color.White)
	}

	// Increase rotation angles for animation
	rotationAngles[0] += 0.01
	rotationAngles[1] += 0.005
	rotationAngles[2] += 0.002

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Drawing is handled in the Update method
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the screen size
	return screenWidth, screenHeight
}

func drawLine(x1, y1, x2, y2 float64, clr color.Color) {
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, clr)
}

func main() {
	// Create a new window with the specified size
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rotating Cube")

	// Create an instance of the Game struct
	game := &Game{}

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println("Error:", err)
	}
}
