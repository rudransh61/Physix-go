package main

import (
	"fmt"
	"time"

	"github.com/rudransh61/Physix-go/internal/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

func main() {
	// Create two rigid bodies for our bouncing balls
	ball1 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 50, Y: 50},
		Velocity: vector.Vector{X: 30, Y: 20},
		Mass:     1,
	}

	ball2 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 150, Y: 150},
		Velocity: vector.Vector{X: -20, Y: -10},
		Mass:     1,
	}

	// Simulation parameters
	dt := 0.1 // Time step for simulation

	for i := 0; i < 100; i++ {
		// Update the position of the balls in the simulation
		github.com/rudransh61/Physix-go.ApplyForce(ball1, vector.Vector{X: 0, Y: 0}, dt)
		github.com/rudransh61/Physix-go.ApplyForce(ball2, vector.Vector{X: 0, Y: 0}, dt)

		// Print the current positions of the balls
		fmt.Printf("Ball1: Position(%f, %f)\n", ball1.Position.X, ball1.Position.Y)
		fmt.Printf("Ball2: Position(%f, %f)\n", ball2.Position.X, ball2.Position.Y)
		fmt.Println("--------")

		// Introduce a delay to visualize the simulation
		time.Sleep(100 * time.Millisecond)
	}
}
