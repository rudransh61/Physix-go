# Documentation 

- [Introduction](#introduction)
- [Installation](#installation)
- [Vectors](#vectors)
   - [Create Vector](#create-vector)
   - [Addition](#addition)
   - [Subtraction](#subtraction)
   - [Inner Product](#inner-product)
   - [Scale Vector](#scale-vector)
   - [Magnitude](#magnitude)
   - [Normalize Vector](#normalize-vector)
   

## Introduction 
Physix.go is a simple, easy-to-use, and fast physics engine written in GoLang. It provides functions to perform physics calculations efficiently, including particle-based physics simulations.

Note : There are some functions for Polygons and all , but they are not good at working , So Please use it like a particle-based physics engine and not for polygons. (Not recommended)

## Installation

So , First of all you need to install [Go](https://go.dev/doc/install)

```bash
go get github.com/rudransh61/Physix-go@v1.0.0  
```

Or you can also start with cloning the repository.

```bash
git clone https://github.com/rudransh61/Physix.go
```

Now run the file `./examples/ex1.go` or copy this code and run it

```go
package main

import (
	"fmt"
	"time"

	"github.com/rudransh61/Physix-go/dynamics/physics"
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
        physix.ApplyForce(ball1, vector.Vector{X: 10, Y: 0}, dt)
        physix.ApplyForce(ball2, vector.Vector{X: 0, Y: 10}, dt)

        fmt.Printf("Ball1: Position(%f, %f)\n", ball1.Position.X, ball1.Position.Y)
        fmt.Printf("Ball2: Position(%f, %f)\n", ball2.Position.X, ball2.Position.Y)
        fmt.Println("--------")

        time.Sleep(100 * time.Millisecond)
    }
}
```

It will Run without any error if everything is working fine.

## Vectors

A Vector is a 1x2 matrix , or an object with 2 values `X`and `Y`.

Like this
```go
type Vector struct {
	X, Y float64
}
```

It will be used everywhere in the library, to store positions , forces and accelerations.

### Create Vector

import this file to use it,
`github.com/rudransh61/Physix-go/pkg/vector`

and create a vector like this:

```go
vec1 := vector.Vector{X: 1, Y: 2}

// or

vec2 := vector.NewVector(1, 2)
```

Operations on Vectors:

### Addition
```go
vec_add12 = vec1.Add(vec2)
```

### Subtraction
```go
vec_sub12 = vec1.Sub(vec2)
```

### Inner Product
```go
vec_inner_product = vec1.Dot(vec2)
```

### Scale Vector
```go
vec_scale = vec1.Scale(2)
```

### Magnitude
```go
vec_mag = vec1.Magnitude()
```

### Normalize Vector
```go
vec_norm = vec1.Normalize()
```
### Distance
```go
vec_dist = vector.Distance(vec1,vec2)
```
## Rigid Body

A Rigid Body is a physical object that has mass and can be affected by forces. It has a position, velocity, and mass.

import this file to use it,
`github.com/rudransh61/Physix-go/pkg/rigidbody`

### Create Rigid Body

```go
body := &rigidbody.RigidBody{
		Position: vector.Vector{X: 50, Y: 50},
		Velocity: vector.Vector{X: 30, Y: 20},
		Mass:     1,
		
		Shape:   'Circle', // or 'Rectangle'
		Radius:  10, // Only for Circle
		Width:   20, // Only for Rectangle
		Height:  30, // Only for Rectangle

		IsMovable: true, // If false, it will not be affected by forces

	}
```

To update the Position and Apply Force on it, use this function:

```go
dt := 0.1 // Time step for simulation
physix.ApplyForce(ball, Force_vector, dt) // Apply force
```

NOTE: Import `github.com/rudransh61/Physix-go/dynamics/physics` to use this functions in `physix`. 

Or access Velocity, Position and Mass of the Rigid Body like this:
```go
ball.Velocity // Get the velocity of the ball as a vector.Vector
ball.Position.X += 5 // Increase the position of the ball in X direction by 5
```