package main

import (
	"fmt"
	"github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/polygon"
	"github.com/rudransh61/Physix-go/pkg/vector"
	"time"
)

// main function for testing
func main() {
	vertices := []vector.Vector{{X: 0, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 2}, {X: 0, Y: 2}}
	Polygon := polygon.NewPolygon(vertices, 5.0, true)

	fmt.Println("Polygon Position:", Polygon.Position)
	fmt.Println("Polygon Vertices:", Polygon.Vertices)
	fmt.Println("Polygon Mass:", Polygon.Mass)
	fmt.Println("Polygon Shape:", Polygon.Shape)
	fmt.Println("Polygon Width:", Polygon.Width)
	fmt.Println("Polygon Height:", Polygon.Height)
	fmt.Println("Polygon Radius:", Polygon.Radius)
	fmt.Println("Polygon IsMovable:", Polygon.IsMovable)

	// Simulation parameters
	dt := 0.1 // Time step for simulation

	for i := 0; i < 100; i++ {
		physix.ApplyForcePolygon(Polygon, vector.Vector{X: 10, Y: 0}, dt)

		fmt.Printf("Polygon: Position(%f, %f)\n", Polygon.Position.X, Polygon.Position.Y)
		fmt.Println("--------")

		time.Sleep(100 * time.Millisecond)
	}
}
