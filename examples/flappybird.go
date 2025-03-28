package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"image/color"
	"math/rand"
	"time"
)

const (
	screenWidth  = 400
	screenHeight = 600
	flapStrength = -30
	pipeWidth    = 50
	pipeGap      = 150
	pipeSpeed    = 2
)

type Bird struct {
	body *rigidbody.RigidBody
}

type Pipe struct {
	x, height float64
}

var (
	bird  = Bird{body: &rigidbody.RigidBody{Position: vector.Vector{X: 100, Y: screenHeight / 2}, Velocity: vector.Vector{X: 0, Y: 0}, Mass: 1, Radius: 10, IsMovable: true}}
	pipes []Pipe
	score int
	gameOver bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	spawnPipe()
}

func spawnPipe() {
	height := float64(rand.Intn(screenHeight-2*pipeGap) + pipeGap)
	pipes = append(pipes, Pipe{x: screenWidth, height: height})
}

func update() error {
	if gameOver {
		return nil
	}

	physix.ApplyForce(bird.body, vector.Vector{X: 0, Y: bird.body.Mass*9.8}, 0.1)
	// bird.body.Position = bird.body.Position.Add(bird.body.Velocity)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bird.body.Velocity.Y = flapStrength
	}

	for i := range pipes {
		pipes[i].x -= pipeSpeed
	}

	if len(pipes) > 0 && pipes[0].x < -pipeWidth {
		pipes = pipes[1:]
		spawnPipe()
		score++
	}

	if bird.body.Position.Y > screenHeight || bird.body.Position.Y < 0 {
		gameOver = true
	}

	for _, pipe := range pipes {
		if bird.body.Position.X+bird.body.Radius > pipe.x && bird.body.Position.X-bird.body.Radius < pipe.x+pipeWidth {
			if bird.body.Position.Y-bird.body.Radius < pipe.height-pipeGap/2 || bird.body.Position.Y+bird.body.Radius > pipe.height+pipeGap/2 {
				gameOver = true
			}
		}
	}
	

	return nil
}

func draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, bird.body.Position.X, bird.body.Position.Y, bird.body.Radius, color.RGBA{255, 255, 0, 255})

	for _, pipe := range pipes {
		ebitenutil.DrawRect(screen, pipe.x, 0, pipeWidth, pipe.height-pipeGap/2, color.RGBA{0, 255, 0, 255})
		ebitenutil.DrawRect(screen, pipe.x, pipe.height+pipeGap/2, pipeWidth, screenHeight, color.RGBA{0, 255, 0, 255})
	}
}

type Game struct{}

func (g *Game) Update() error { return update() }
func (g *Game) Draw(screen *ebiten.Image) { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return screenWidth, screenHeight }

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Flappy Bird Clone")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
