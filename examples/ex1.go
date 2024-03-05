// main.go
package main

import (
	"fmt"
	"time"

	"physix/internal/physics"
	"physix/pkg/rigidbody"
	"physix/pkg/vector"
	"physix/internal/collision"
)

func main() {
	// Create two rigid bodies for our bouncing balls
	ball1 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 50, Y: 50},
		Velocity: vector.Vector{X: 30, Y: 20},
		Mass:     1,
		Circle:   rigidbody.Circle{Position: vector.Vector{X: 50, Y: 50}, Radius: 5},
	}

	ball2 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 150, Y: 150},
		Velocity: vector.Vector{X: -20, Y: -10},
		Mass:     1,
		Circle:   rigidbody.Circle{Position: vector.Vector{X: 150, Y: 150}, Radius: 5},
	}

	// Simulation parameters
	dt := 0.1 // Time step for simulation

	for i := 0; i < 100; i++ {
        physix.UpdateRigidBody(ball1, vector.Vector{X: 0, Y: 0}, dt)
        physix.UpdateRigidBody(ball2, vector.Vector{X: 0, Y: 0}, dt)

        // Check for collision using circles
        if collision.CircleCollide(ball1.Circle, ball2.Circle) {
            // Handle collision response (e.g., reverse velocities)
            ball1.Velocity = ball1.Velocity.Scale(-1)
            ball2.Velocity = ball2.Velocity.Scale(-1)
        }

        fmt.Printf("Ball1: Position(%f, %f)\n", ball1.Position.X, ball1.Position.Y)
        fmt.Printf("Ball2: Position(%f, %f)\n", ball2.Position.X, ball2.Position.Y)
        fmt.Println("--------")

        time.Sleep(100 * time.Millisecond)
    }
}
