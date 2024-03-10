// internal/collision/collision.go
package collision

import (
	"physix/pkg/rigidbody"
	// "math"
	"physix/pkg/vector"
)


// CheckCollision checks if two rectangles (RigidBody instances) are colliding.
func RectangleCollided(rect1 *rigidbody.RigidBody, rect2 *rigidbody.RigidBody) bool {
	if(rect1.Shape==rect2.Shape && rect1.Shape=="Rectangle"){
		left1, top1, right1, bottom1 := rect1.Position.X, rect1.Position.Y, rect1.Position.X+rect1.Width, rect1.Position.Y+rect1.Height
		left2, top2, right2, bottom2 := rect2.Position.X, rect2.Position.Y, rect2.Position.X+rect2.Width, rect2.Position.Y+rect2.Height

		return right1 > left2 && left1 < right2 && bottom1 > top2 && top1 < bottom2
	}
	return false
	
}

func BounceOnCollision(rect1, rect2 *rigidbody.RigidBody, e float64) {
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



// Circle collision detection

func CircleCollided(circle1 *rigidbody.RigidBody, circle2 *rigidbody.RigidBody) bool {
	if(circle1.Shape==circle2.Shape && circle1.Shape=="Circle"){
		if(vector.Distance(circle1.Position, circle2.Position)<(circle1.Radius+circle2.Radius)){
			return true
		}else{
			return false
		}
	}
	return false
}