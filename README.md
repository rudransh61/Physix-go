# Physix.go

A simple Physics Engine in GoLang

<div algin="center">
  <img src="/example.gif" width="200">
</div>

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

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

Then Run the example files from `./example` folder

for example : 
```
go run ./example/ex4.go //which is simple circular motion
```

## Documentation

## Vectors

Vectors are a datatype to store vectors.

Import the following file to use vectors
```go
package main 

import (
  //...other imports
  "physix/pkg/vector"
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


## Examples

Check examples in `./example` folder

## Contributing

New contributors are welcome!!
If you have any doubt related to its working you can ask to us by opening an issue

## License

`LICENSE.md` file

## Acknowledgments

Inspired from [Coding Train - Daniel Shiffman](https://www.youtube.com/channel/UCvjgXvBlbQiydffZU7m1_aw)
```
