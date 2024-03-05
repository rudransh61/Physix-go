// internal/collision/collision.go
package collision

import (
	"physix/pkg/rigidbody"
	"math"
)

// CircleOverlap checks if two circles overlap.
func CircleCollide(circle1, circle2 rigidbody.Circle) bool {
	distance := math.Sqrt(math.Pow(circle2.Position.X-circle1.Position.X, 2) + math.Pow(circle2.Position.Y-circle1.Position.Y, 2))
	sumRadii := circle1.Radius + circle2.Radius

	return distance < sumRadii
}