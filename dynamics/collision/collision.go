// dynamics/collision/collision.go
package collision

import (
	"math"

	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

// CheckCollision checks if two rectangles (RigidBody instances) are colliding.
func RectangleCollided(rect1 *rigidbody.RigidBody, rect2 *rigidbody.RigidBody) bool {
	if rect1.Shape == rect2.Shape && rect1.Shape == "Rectangle" {
		left1, top1, right1, bottom1 := rect1.Position.X, rect1.Position.Y, rect1.Position.X+rect1.Width, rect1.Position.Y+rect1.Height
		left2, top2, right2, bottom2 := rect2.Position.X, rect2.Position.Y, rect2.Position.X+rect2.Width, rect2.Position.Y+rect2.Height

		return right1 > left2 && left1 < right2 && bottom1 > top2 && top1 < bottom2
	}
	return false
}

// Circle collision detection
func CircleCollided(circle1 *rigidbody.RigidBody, circle2 *rigidbody.RigidBody) bool {
	if circle1.Shape == circle2.Shape && circle1.Shape == "Circle" {
		return vector.Distance(circle1.Position, circle2.Position) < (circle1.Radius + circle2.Radius)
	}
	return false
}

// Circle-Rectangle collision detection
func CircleRectangleCollided(circle *rigidbody.RigidBody, rect *rigidbody.RigidBody) bool {
	if circle.Shape == "Circle" && rect.Shape == "Rectangle" {
		closestX := math.Max(rect.Position.X, math.Min(circle.Position.X, rect.Position.X+rect.Width))
		closestY := math.Max(rect.Position.Y, math.Min(circle.Position.Y, rect.Position.Y+rect.Height))
		distance := vector.Distance(vector.Vector{X: closestX, Y: closestY}, circle.Position)
		return distance < circle.Radius
	}
	return false
}

// Prevent Rectangle-Rectangle Overlap
func PreventRectangleOverlap(rect1, rect2 *rigidbody.RigidBody) {
	if RectangleCollided(rect1, rect2) {
		mtvX := math.Min(rect1.Position.X+rect1.Width-rect2.Position.X, rect2.Position.X+rect2.Width-rect1.Position.X)
		mtvY := math.Min(rect1.Position.Y+rect1.Height-rect2.Position.Y, rect2.Position.Y+rect2.Height-rect1.Position.Y)

		if rect1.IsMovable && rect2.IsMovable {
			if mtvX < mtvY {
				if rect1.Position.X < rect2.Position.X {
					rect1.Position.X -= mtvX / 2
					rect2.Position.X += mtvX / 2
				} else {
					rect1.Position.X += mtvX / 2
					rect2.Position.X -= mtvX / 2
				}
			} else {
				if rect1.Position.Y < rect2.Position.Y {
					rect1.Position.Y -= mtvY / 2
					rect2.Position.Y += mtvY / 2
				} else {
					rect1.Position.Y += mtvY / 2
					rect2.Position.Y -= mtvY / 2
				}
			}
		}

		if rect2.IsMovable && !rect1.IsMovable {
			if mtvX < mtvY {
				if rect2.Position.X < rect1.Position.X {
					rect2.Position.X -= mtvX
				} else {
					rect2.Position.X += mtvX
				}
			} else {
				if rect2.Position.Y < rect1.Position.Y {
					rect2.Position.Y -= mtvY
				} else {
					rect2.Position.Y += mtvY
				}
			}
		}
		if rect1.IsMovable && !rect2.IsMovable {
			if mtvX < mtvY {
				if rect1.Position.X < rect2.Position.X {
					rect1.Position.X -= mtvX
				} else {
					rect1.Position.X += mtvX
				}
			} else {
				if rect1.Position.Y < rect2.Position.Y {
					rect1.Position.Y -= mtvY
				} else {
					rect1.Position.Y += mtvY
				}
			}
		}

	}
}

// Prevent Circle-Circle Overlap
func PreventCircleOverlap(circle1, circle2 *rigidbody.RigidBody) {
	if CircleCollided(circle1, circle2) {
		delta := circle2.Position.Sub(circle1.Position)
		distance := delta.Magnitude()
		overlap := (circle1.Radius + circle2.Radius - distance) / 2
		if distance > 0 {
			correction := delta.Normalize().Scale(overlap)
			circle1.Position = circle1.Position.Sub(correction)
			circle2.Position = circle2.Position.Add(correction)
		}
	}
}

// Prevent Circle-Rectangle Overlap
func PreventCircleRectangleOverlap(circle, rect *rigidbody.RigidBody) {
	if CircleRectangleCollided(circle, rect) {
		closestX := math.Max(rect.Position.X, math.Min(circle.Position.X, rect.Position.X+rect.Width))
		closestY := math.Max(rect.Position.Y, math.Min(circle.Position.Y, rect.Position.Y+rect.Height))
		closestPoint := vector.Vector{X: closestX, Y: closestY}
		delta := circle.Position.Sub(closestPoint)
		distance := delta.Magnitude()
		overlap := circle.Radius - distance
		if distance > 0 {
			correction := delta.Normalize().Scale(overlap)
			circle.Position = circle.Position.Add(correction)
		}
	}
}

func BounceOnCollision(rect1, rect2 *rigidbody.RigidBody, e float64) {
	// fmt.Println("Entering BounceOnCollision function")
	// defer fmt.Println("Exiting BounceOnCollision function")

	if rect1.IsMovable && rect2.IsMovable {
		// Calculate the center of mass velocities
		v1 := rect1.Velocity
		m1 := rect1.Mass
		v2 := rect2.Velocity
		m2 := rect2.Mass

		rect1.Velocity = v1.Scale((m1 - e*m2) / (m1 + m2)).Add(v2.Scale((1 + e) * m2 / (m1 + m2)))
		rect2.Velocity = v2.Scale((m2 - e*m1) / (m1 + m2)).Add(v1.Scale((1 + e) * m1 / (m1 + m2)))
	} else if rect1.IsMovable && !rect2.IsMovable {
		// Bounce only rect1
		rect1.Velocity = rect1.Velocity.Scale(-e)
	} else if !rect1.IsMovable && rect2.IsMovable {
		// Bounce only rect2
		rect2.Velocity = rect2.Velocity.Scale(-e)
	}
	// No bounce if both are static
}
