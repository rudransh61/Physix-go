<u><h1> Physix.go</h1></u>

<div algin="center" style="margin-bottom:100px">
	<h2>A simple Physics Engine in GoLang☻</h2>
  <img src="/Phi 6.png" width="300">
</div>


<div algin="center">
  <img src="/example.gif" width="200">
  <img src="/example1.gif" width="200">
  <img src="/example2.gif" width="200">
</div>

## Introduction

Physix.go is simple, easy to use , fast , physics engine written in GoLang
With functions to perform physics calculations faster...

## Features
- Vector Calculations
- Physics Calculations
- Easy to use with [Ebiten.go](https://ebitengine.org/)

## Getting Started


### Prerequisites

GoLang must be installed
And [Ebiten](https://ebiten.org)

### Installation

Install 

To start , Clone this project
```
git clone https://github.com/rudransh61/Physix.go
```

Then Run the example files from `./examples` folder

for example : 
```
go run ./examples/ex4.go //which is simple circular motion
```

## Documentation

## Vectors

Vectors are a datatype to store vectors.

Import the following file to use vectors
```go
package main 

import (
  //...other imports
  "github.com/rudransh61/Physix-go/pkg/vector"
)
```

### To make a vector
```go
var MyVector := vector.Vector{X: 30, Y: 20}
// X is x component and Y is y component of Vector
```

Using Function
```go
var NewVec := vector.NewVector(x,y)
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
var DotProduct := Vec1.InnerProduct(Vec2)
```

##### Scale a Vector by a scalar
```go
var ScaledVector := Vec1.Scale(num)
```

##### Magnitude of a Vector
```go
var Magnitude := Vec1.Magnitude()
```

##### Normalize a Vector
```go
var NormalizeVector = Vec1.Normalize()
```

##### Distance between Head of 2 Vectors
```go
var distance = vector.Distance(Vec1,Vec2)
```

##### Perpendicular Vector of a given Vector
```go
vector.Vector Orthogonal_Vector = vector.Orthogonal(Vec1)
```

## Basics
In this Physics Engine , we call every physical entity a : <b>RigidBody</b>

There are 2 types of RigidBody According to their <b>Shape</b>
 - Rectangle
 - Circle

Every RigidBody have following properties :-

<pre>
 - <b>Position</b>     : <i>Vector</i>           # Required while initializing
 - <b>Velocity</b>     : <i>Vector</i>
 - <b>Force</b>        : <i>Vector</i>
 - <b>Mass</b>         : <i>float64</i>          # Required while initializing for Collision and Forces
 - <b>Shape</b>        : <i>string</i>           # Required while initializing for Collision
 - <b>Width</b>        : <i>float64</i>          # Required while initializing only for Shape :- "Rectangle"
 - <b>Height</b>       : <i>float64</i>          # Required while initializing only for Shape :- "Rectangle"
 - <b>Radius</b>       : <i>float64</i>          # Required while initializing only for Shape :- "Circle"
 - <b>IsMovable</b>    : <i>bool</i>             # Required while initializing for Collision and Forces
</pre>


## RigidBody
To create an instance of RigidBody you need to provide all the required fields .

First Import these files,
```golang
import (
  "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
)
```

Example :-
```golang
ball = &rigidbody.RigidBody{
		Position:  vector.Vector{X: 400, Y: 100},
		Velocity:  vector.Vector{X: 0, Y: 2},
		Mass:      1,
		Force:     vector.Vector{X: 0, Y: 5},
		IsMovable: true,
		Shape:     "Rectangle",
		Width:     50,
		Height:    50,
	}
```

To update position of a RigidBody, Use <b>ApplyForce</b> in a loop ,
Example :- 
```golang
for i := 0; i < 100; i++ {
        github.com/rudransh61/Physix-go.ApplyForce(ball1, vector.Vector{X: 10, Y: 0}, dt) // Make the vector (0,0) to apply no force
        // .. other code
    }
```

To access or change the  <b>Force</b> , <b>Velocity</b> , <b>Position</b>,
```golang
ball.Velocity // Get the velocity of ball as a vector.Vector
ball.Position.X += 5 // Increase the position of ball in X direction by 5
```

## Collision Detection
There are 2 types of collision systems for different shapes.
 - Rectangle-Rectangle collision
 - Circle-Circle collision

### Rectangle Collision
For example you have 2 Rectangles, Like this :-
```golang
rect = &rigidbody.RigidBody{
	Position: vector.Vector{X: 100, Y: 200},
	Velocity: vector.Vector{X: 50, Y: -50},
	Mass:     1.0,
	Shape : "Rectangle",
	Width : 100,
	Height : 90,
	IsMovable : true,
}
rect2 = &rigidbody.RigidBody{
	Position: vector.Vector{X: 400, Y: 300},
	Velocity: vector.Vector{X: 60, Y: 50},
	Mass:     2.0,
	Shape : "Rectangle",
	Width : 70,
	Height : 70,
	IsMovable : true,
}
```

Now you want to detect collision between them ,
```golang
if(collision.RectangleCollided(rect,rect2)){
	fmt.Println("Collided!")
}
```

And if you want to add a bounce effect in this collision according to the <b>Momentum Conservation</b> and <b>Energy Conservation</b>,
```golang
if(collision.RectangleCollided(rect,rect2)){
	float64 e = 0.9999999999;                // e is coefficient of restitution in collision
	collision.BounceOnCollision(ball,ball2,e)// NOTE :- e<1 is bit glitchy and goes wild, use it on your own risk :)
}
```
==NOTE :- e<1 is bit glitchy and goes wild, use it on your own risk :) ==

### Circle Collision

Now if you want to detect collisions between a circle and a circle,
```golang
if(collision.CircleCollided(rect,rect2)){
	fmt.Println("Collided!")
}
```
And same <a href='#'>BounceOnCollision</a> function for Bouncing ...


## Examples

Check examples in `./example` folder

## Polygon 

To create a polygon , we made a new entity <b>Polygon</b> which inherits all Properties of <b>RigidBody</b> with some additional features.

<pre>
 - <b>Position</b>     : <i>Vector</i>           # Required while initializing
 - <b>Velocity</b>     : <i>Vector</i>
 - <b>Force</b>        : <i>Vector</i>
 - <b>Mass</b>         : <i>float64</i>          # Required while initializing for Collision and Forces
 - <b>IsMovable</b>    : <i>bool</i>             # Required while initializing for Collision and Forces
 - <b>Vertices</b>     : <i>[] Vector</i>        # Required while initializing only for Shape :- "Polygon"
</pre>

To initialize a polygon , use **NewPolygon** function 
### Example 
```go
import (
	...
	"github.com/rudransh61/Physix-go/pkg/polygon"
	...
)
...
vertices := []vector.Vector{
	{X: 250, Y: 50}, 
	{X: 200, Y: 100}, 
	{X: 200, Y: 50}, 
	{X: 350, Y: 200}
}

ball = polygon.NewPolygon(vertices, 50, true)
```

## Some Dynamics

### Physics

To update our <button title="object">'entity'</button> , 
We have 2 functions which are <b>ApplyForce</b> and <b>ApplyForcePolygon</b> , as the name suggests one is for Rigidbody and one for polygons.

This function will move one frame forward or 'dt' time forward (which is time between 2 frames).

<mark>NOTE: Define dt(0.1 mostly) at top globally for good code</mark>

```go
polygonA := polygon.NewPolygon(vertices, 50, true) # Example

# Define dt
dt := 0.1

physix.ApplyForcePolygon(
	polygonA, 
	vector.Vector{X: 10, Y: 0}, 
	dt,
) 
# To applyForce on polygon and update its position and everything

```

Similarly for RigidBody : 
```go
rigid := &rigidbody.Rigidbody{...}

physix.ApplyForce(rigid, vector.Vector{X:10,Y:0}, dt)

```

To get both the utilities in the code import this file:
```go
import(
	...
	"github.com/rudransh61/Physix-go/dynamics/physics"
 	...
)
```

## Contributing

New contributors are welcome!!
If you have any doubt related to its working you can ask to us by opening an issue

## License

`LICENSE.md` file

## Acknowledgments

Inspired from [Coding Train - Daniel Shiffman](https://www.youtube.com/channel/UCvjgXvBlbQiydffZU7m1_aw)
