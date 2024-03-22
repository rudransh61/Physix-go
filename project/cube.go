package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"math"
	"image/color"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	cubeVertices = [8][3]float64{
		{-1, -1, -1}, {-1, 1, -1}, {1, 1, -1}, {1, -1, -1},
		{-1, -1, 1}, {-1, 1, 1}, {1, 1, 1}, {1, -1, 1},
	}
	cubeFaces = [6][4]int{
		{0, 1, 2, 3}, {1, 5, 6, 2}, {5, 4, 7, 6},
		{4, 0, 3, 7}, {0, 1, 5, 4}, {3, 2, 6, 7},
	}
)

type Game struct {
	angle float64
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.angle += 0.02
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, face := range cubeFaces {
		var points [4][2]float64
		for i, v := range face {
			x := cubeVertices[v][0]
			y := cubeVertices[v][1]
			z := cubeVertices[v][2]

			// Rotate around the y-axis
			x, z = x*math.Cos(g.angle)-z*math.Sin(g.angle), x*math.Sin(g.angle)+z*math.Cos(g.angle)

			// Rotate around the x-axis
			y, z = y*math.Cos(g.angle)-z*math.Sin(g.angle), y*math.Sin(g.angle)+z*math.Cos(g.angle)

			// Projection
			scale := 200 / (z + 3)
			points[i][0] = x*scale + screenWidth/2
			points[i][1] = y*scale + screenHeight/2
		}

		// Draw the face of the cube with the specified color
		ebitenutil.DrawLine(screen, points[0][0], points[0][1], points[1][0], points[1][1], color.RGBA{R: 255, G: 0, B: 0, A: 255})
		ebitenutil.DrawLine(screen, points[1][0], points[1][1], points[2][0], points[2][1], color.RGBA{R: 255, G: 0, B: 0, A: 255})
		ebitenutil.DrawLine(screen, points[2][0], points[2][1], points[3][0], points[3][1], color.RGBA{R: 255, G: 0, B: 0, A: 255})
		ebitenutil.DrawLine(screen, points[3][0], points[3][1], points[0][0], points[0][1], color.RGBA{R: 255, G: 0, B: 0, A: 255})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rotating Cube")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
