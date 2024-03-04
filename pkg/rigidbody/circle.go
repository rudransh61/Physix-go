// pkg/rigidbody/circle.go
package rigidbody

import "physics/pkg/vector"

// Circle represents a 2D circle.
type Circle struct {
    Position vector.Vector
    Radius   float64
}
