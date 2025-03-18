
# Physix.go

<div align="center" style="margin-bottom:100px">
	<h2>A Simple Physics Engine in GoLang ☻</h2>
  <img src="/Phi 6.png" width="300">
</div>

<div align="center">
  <img src="/example.gif" width="200">
  <img src="/example1.gif" width="200">
  <img src="/example2.gif" width="200">
</div>

## Introduction

Physix.go is a simple, easy-to-use, and fast physics engine written in GoLang. It provides functions to perform physics calculations efficiently, including particle-based physics simulations.

## Features
- Vector Calculations
- Physics Calculations
- Spring Dynamics
- Easy to use with [Ebiten.go](https://ebitengine.org/)

## Getting Started

### Prerequisites

- GoLang must be installed.
- [Ebiten](https://ebiten.org) must be installed.

### Installation

To start, clone this project:
```bash
git clone https://github.com/rudransh61/Physix.go
```

Then run the example files from the `./examples` folder. For example:
```bash
go run ./examples/ex4.go # which is a simple circular motion
```

## Documentation

### Vectors

Vectors are a datatype to store vectors. Import the following file to use vectors:
```go
package main 

import (
  //...other imports
  "github.com/rudransh61/Physix-go/pkg/vector"
)
```

#### To make a vector
```go
var MyVector = vector.Vector{X: 30, Y: 20}
// X is the x component and Y is the y component of the Vector
```

Using Function
```go
var NewVec = vector.NewVector(x, y)
```

##### Add Vector
```go
var NewVector = Vec1.Add(Vec2)
```

##### Subtract Vector
```go
var NewVector = Vec1.Sub(Vec2)
```

##### Inner Product of 2 Vectors
```go
var DotProduct = Vec1.InnerProduct(Vec2)
```

##### Scale a Vector by a scalar
```go
var ScaledVector = Vec1.Scale(num)
```

##### Magnitude of a Vector
```go
var Magnitude = Vec1.Magnitude()
```

##### Normalize a Vector
```go
var NormalizeVector = Vec1.Normalize()
```

##### Distance between Heads of 2 Vectors
```go
var distance = vector.Distance(Vec1, Vec2)
```

##### Perpendicular Vector of a given Vector
```go
var Orthogonal_Vector = vector.Orthogonal(Vec1)
```

### RigidBody

To create an instance of RigidBody, you need to provide all the required fields. First, import these files:
```go
import (
  "github.com/rudransh61/Physix-go/dynamics/physics"
  "github.com/rudransh61/Physix-go/pkg/rigidbody"
)
```

Example:
```go
ball = &rigidbody.RigidBody{
	Position:  vector.Vector{X: 400, Y: 100},
	Velocity:  vector.Vector{X: 0, Y: 2},
	Mass:      1,
	Force:     vector.Vector{X: 0, Y: 5},
	IsMovable: true,
	Shape:     "Circle", // Example shape
	Radius:    10,       // Required for Circle
}
```

To update the position of a RigidBody, use **ApplyForce** in a loop:
```go
for i := 0; i < 100; i++ {
    physics.ApplyForce(ball, vector.Vector{X: 10, Y: 0}, dt) // Apply force
    // .. other code
}
```

To access or change the **Force**, **Velocity**, **Position**:
```go
ball.Velocity // Get the velocity of the ball as a vector.Vector
ball.Position.X += 5 // Increase the position of the ball in X direction by 5
```

### Collision Detection

There are two types of collision systems for different shapes:
- Rectangle-Rectangle collision
- Circle-Circle collision

#### Rectangle Collision

For example, you have two Rectangles:
```go
rect1 = &rigidbody.RigidBody{
	Position: vector.Vector{X: 100, Y: 200},
	Velocity: vector.Vector{X: 50, Y: -50},
	Mass:     1.0,
	Shape:    "Rectangle",
	Width:    100,
	Height:   90,
	IsMovable: true,
}
rect2 = &rigidbody.RigidBody{
	Position: vector.Vector{X: 400, Y: 300},
	Velocity: vector.Vector{X: 60, Y: 50},
	Mass:     2.0,
	Shape:    "Rectangle",
	Width:    70,
	Height:   70,
	IsMovable: true,
}
```

Now you want to detect collision between them:
```go
if collision.RectangleCollided(rect1, rect2) {
	fmt.Println("Collided!")
}
```

And if you want to add a bounce effect in this collision according to the **Momentum Conservation** and **Energy Conservation**:
```go
if collision.RectangleCollided(rect1, rect2) {
	float64 e = 0.9999999999; // e is the coefficient of restitution in collision
	collision.BounceOnCollision(rect1, rect2, e) // NOTE: e<1 is a bit glitchy and goes wild, use it at your own risk :)
}
```

#### Circle Collision

Now if you want to detect collisions between a circle and a circle:
```go
if collision.CircleCollided(rect1, rect2) {
	fmt.Println("Collided!")
}
```
And use the same `BounceOnCollision` function for bouncing.

### Springs

Springs can be used to simulate elastic connections between rigid bodies. To create a spring, you need to define two rigid bodies and specify the spring's stiffness and damping properties.

Example:
```go
spring := spring.NewSpring(ballA, ballB, stiffness, damping)
```

To apply the spring force in your simulation loop:
```go
spring.ApplyForce()
```

This will apply the forces based on Hooke's Law and damping to the connected rigid bodies.

## Examples

Check examples in the `./examples` folder.

## Some Dynamics

### Physics

To update our **entity**, we have two functions: **ApplyForce** and **ApplyForcePolygon**, as the name suggests, one is for RigidBody and one for polygons.

This function will move one frame forward or 'dt' time forward (which is the time between two frames).

**NOTE:** Define `dt` (0.1 mostly) at the top globally for good code.

```go
ball = &rigidbody.RigidBody{
	Position: vector.Vector{X: 400, Y: 100},
	Velocity: vector.Vector{X: 0, Y: 2},
	Mass:     1,
	Force:    vector.Vector{X: 0, Y: 5},
	IsMovable: true,
}

physics.ApplyForce(ball, vector.Vector{X: 10, Y: 0}, dt) // To apply force on the rigid body
```

To get both utilities in the code, import this file:
```go
import (
	...
	"github.com/rudransh61/Physix-go/dynamics/physics"
	...
)
```

## Contributing

New contributors are welcome! If you have any doubts related to its working, you can ask us by opening an issue.

## License

See `LICENSE.md` file.

## Acknowledgments

Inspired by [Coding Train - Daniel Shiffman](https://www.youtube.com/channel/UCvjgXvBlbQiydffZU7m1_aw)

