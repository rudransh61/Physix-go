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

func BounceOnCollision(body1, body2 *rigidbody.RigidBody, e float64) {
	// fmt.Println("Entering BounceOnCollision function")
	// defer fmt.Println("Exiting BounceOnCollision function")

	if body1.IsMovable && body2.IsMovable {
		// Calculate the center of mass velocities
		v1 := body1.Velocity
		m1 := body1.Mass
		v2 := body2.Velocity
		m2 := body2.Mass

		//body1.Velocity = v1.Scale((m1 - e*m2) / (m1 + m2)).Add(v2.Scale((1 + e) * m2 / (m1 + m2)))
		//body2.Velocity = v2.Scale((m2 - e*m1) / (m1 + m2)).Add(v1.Scale((1 + e) * m1 / (m1 + m2)))
		// Velocity after
		center1 := vector.NewVector(0, 0)
		if body1.Shape == "Circle" {
			center1.X = body1.Position.X
			center1.Y = body1.Position.Y
		} else {
			center1.X = body1.Position.X + body1.Width/2
			center1.Y = body1.Position.Y + body1.Height/2
		}

		center2 := vector.NewVector(0, 0)
		if body2.Shape == "Circle" {
			center2.X = body2.Position.X
			center2.Y = body2.Position.Y
		} else {
			center2.X = body2.Position.X + body2.Width/2
			center2.Y = body2.Position.Y + body2.Height/2
		}
		r12 := center2.Sub(center1)
		r21 := center1.Sub(center2)
		v1alongr12 := r12.Normalize().Scale(v1.InnerProduct(r12.Normalize()))
		v2alongr21 := r21.Normalize().Scale(v2.InnerProduct(r21.Normalize()))
		v1perp := v1.Sub(v1alongr12)
		v2perp := v2.Sub(v2alongr21)
		newVelocity1normal := v1alongr12.Scale((m1 - e*m2) / (m1 + m2)).Add(v2alongr21.Scale((1 + e) * m2 / (m1 + m2)))
		newVelocity2normal := v2alongr21.Scale((m2 - e*m1) / (m1 + m2)).Add(v1alongr12.Scale((1 + e) * m1 / (m1 + m2)))
		newVecloity1 := newVelocity1normal.Add(v1perp)
		newVecloity2 := newVelocity2normal.Add(v2perp)

		body1.Velocity = newVecloity1
		body2.Velocity = newVecloity2
	} else if body1.IsMovable && !body2.IsMovable {
		// Bounce only rect1
		body1.Velocity = body1.Velocity.Scale(-e)
	} else if !body1.IsMovable && body2.IsMovable {
		// Bounce only rect2
		body2.Velocity = body2.Velocity.Scale(-e)
	}
	// No bounce if both are static
}
