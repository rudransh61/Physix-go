// internal/collision/collision.go
package collision

import (
	"physix/pkg/rigidbody"
	// "math"
	// "physix/pkg/vector"
)


// CheckCollision checks if two rectangles (RigidBody instances) are colliding.
func RectangleCollided(rect1 *rigidbody.RigidBody, rect2 *rigidbody.RigidBody , Width1 float64, Width2 float64, Height1 float64 , Height2 float64) bool {
	left1, top1, right1, bottom1 := rect1.Position.X, rect1.Position.Y, rect1.Position.X+Width1, rect1.Position.Y+Height1
	left2, top2, right2, bottom2 := rect2.Position.X, rect2.Position.Y, rect2.Position.X+Width2, rect2.Position.Y+Height2

	return right1 > left2 && left1 < right2 && bottom1 > top2 && top1 < bottom2
}

func BounceOnCollision(rect1, rect2 *rigidbody.RigidBody, e float64) {
	// Calculate the center of mass velocities
	v1 := rect1.Velocity
	m1 := rect1.Mass
	v2 := rect2.Velocity
	m2 := rect2.Mass

	rect1.Velocity = v1.Scale((m1-e*m2)/(m1+m2)).Add(v2.Scale((1+e)*m2/(m1+m2)))  
	rect2.Velocity = v2.Scale((m2-e*m1)/(m1+m2)).Add(v1.Scale((1+e)*m1/(m1+m2)))  
}