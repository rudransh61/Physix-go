// dynamics/collision/collision.go
package collision

import (
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/polygon"
	"math"
	"github.com/rudransh61/Physix-go/pkg/vector"
    // "fmt"
)

// CheckCollision checks if two rectangles (RigidBody instances) are colliding.
func RectangleCollided(rect1 *rigidbody.RigidBody, rect2 *rigidbody.RigidBody) bool {
    // // fmt.Println("Entering RectangleCollided function")
    // defer fmt.Println("Exiting RectangleCollided function")
    
    if rect1.Shape == rect2.Shape && rect1.Shape == "Rectangle" {
        left1, top1, right1, bottom1 := rect1.Position.X, rect1.Position.Y, rect1.Position.X+rect1.Width, rect1.Position.Y+rect1.Height
        left2, top2, right2, bottom2 := rect2.Position.X, rect2.Position.Y, rect2.Position.X+rect2.Width, rect2.Position.Y+rect2.Height

        return right1 > left2 && left1 < right2 && bottom1 > top2 && top1 < bottom2
    }
    return false
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

// Circle collision detection

func CircleCollided(circle1 *rigidbody.RigidBody, circle2 *rigidbody.RigidBody) bool {
    // fmt.Println("Entering CircleCollided function")
    // defer fmt.Println("Exiting CircleCollided function")
    
    if circle1.Shape == circle2.Shape && circle1.Shape == "Circle" {
        if vector.Distance(circle1.Position, circle2.Position) < (circle1.Radius + circle2.Radius) {
            return true
        } else {
            return false
        }
    }
    return false
}

// Polygon collision detection SAT implementation.
// Link : https://dyn4j.org/2010/01/sat/

// Project calculates the projection of a polygon onto a given axis.
func Project(p *polygon.Polygon, axis vector.Vector) (float64, float64) {
    // fmt.Println("Entering Project function")
    // defer fmt.Println("Exiting Project function")
    // 
    min := axis.InnerProduct(p.Vertices[0])
    max := min
    for i := 1; i < len(p.Vertices); i++ {
        d := axis.InnerProduct(p.Vertices[i])
        if d < min {
            min = d
        } else if d > max {
            max = d
        }
    }
    return min, max
}

// Overlap checks if two intervals overlap.
func Overlap(min1, max1, min2, max2 float64) bool {
    // fmt.Println("Entering Overlap function")
    // defer fmt.Println("Exiting Overlap function")
    
    return !(max1 < min2 || max2 < min1)
}

// Collides checks if two polygons are colliding using the SAT algorithm.
// If they are colliding, resolves the collision by applying appropriate impulses.
func PolygonCollision(poly1, poly2 *polygon.Polygon) bool {
    // fmt.Println("Entering PolygonCollision function")
    // defer fmt.Println("Exiting PolygonCollision function")
    
    for i := 0; i < len(poly1.Vertices); i++ {
        edge := vector.Vector{
            X: poly1.Vertices[(i+1)%len(poly1.Vertices)].X - poly1.Vertices[i].X,
            Y: poly1.Vertices[(i+1)%len(poly1.Vertices)].Y - poly1.Vertices[i].Y,
        }
        normal := vector.Orthogonal(edge).Normalize()
        min1, max1 := Project(poly1, normal)
        min2, max2 := Project(poly2, normal)
        if !Overlap(min1, max1, min2, max2) {
            return false
        }
    }

    for i := 0; i < len(poly2.Vertices); i++ {
        edge := vector.Vector{
            X: poly2.Vertices[(i+1)%len(poly2.Vertices)].X - poly2.Vertices[i].X,
            Y: poly2.Vertices[(i+1)%len(poly2.Vertices)].Y - poly2.Vertices[i].Y,
        }
        normal := vector.Orthogonal(edge).Normalize()
        min1, max1 := Project(poly1, normal)
        min2, max2 := Project(poly2, normal)
        if !Overlap(min1, max1, min2, max2) {
            return false
        }
    }

    // If we reach here, the polygons are colliding
    // Resolve the collision by applying impulses
    // resolveCollision(poly1, poly2)

    return true
}

// ResolveCollision resolves the collision between two polygons by applying appropriate impulses and torques.
func ResolveCollision(poly1, poly2 *polygon.Polygon, e, torqueCoefficient float64) {
    // fmt.Println("Entering ResolveCollision function")
    // defer fmt.Println("Exiting ResolveCollision function")
    
    // Find the MTV (Minimum Translation Vector) to separate the polygons
    mtv := FindMTV(poly1, poly2)

    // Apply the MTV to separate the polygons
    poly1.Move(mtv)
    poly2.Move(mtv.Scale(-1))

    // Calculate relative velocity
    relativeVelocity := poly1.Velocity.Sub(poly2.Velocity)

    // Calculate velocity along the collision normal
    velocityAlongNormal := relativeVelocity.InnerProduct(mtv)

    // If velocities are separating, no collision resolution needed
    if velocityAlongNormal > 0 {
        return
    }

    // Calculate impulse scalar
    j := -(1.0 + e) * velocityAlongNormal / (1/poly1.Mass + 1/poly2.Mass)

    // Apply impulses to resolve collision
    impulse := mtv.Scale(j)
    impulseMag := math.Min(impulse.Magnitude(), 50.0)
    impulse = mtv.Normalize().Scale(impulseMag)
    poly1.ApplyImpulse(impulse)
    poly2.ApplyImpulse(impulse.Scale(-1))

    // Calculate torque impulses
    torque1 := calculateTorqueImpulse(poly1, mtv)
    torque2 := calculateTorqueImpulse(poly2, mtv)

    // Apply torque impulses to resolve collision
    poly1.ApplyTorque(torque1 * torqueCoefficient)
    poly2.ApplyTorque(-torque2 * torqueCoefficient)
}

// calculateTorqueImpulse calculates the torque impulse for a polygon given the MTV.
func calculateTorqueImpulse(poly *polygon.Polygon, mtv vector.Vector) float64 {
    // fmt.Println("Entering calculateTorqueImpulse function")
    // defer fmt.Println("Exiting calculateTorqueImpulse function")
    
    // Calculate the perpendicular distance from centroid to the collision point
    // This is used to calculate torque
    centroidToCollision := PerpendicularDistance(poly.Position, poly.Position, mtv) // Pass polygon centroid position as the first argument

    // Calculate the torque impulse
    torque := centroidToCollision * mtv.Magnitude() // Adjust this calculation based on your method

    return torque
}

// FindMTV finds the Minimum Translation Vector (MTV) to separate two polygons.
func FindMTV(poly1, poly2 *polygon.Polygon) vector.Vector {
    // fmt.Println("Entering FindMTV function")
    // defer fmt.Println("Exiting FindMTV function")
    
    minOverlap := math.MaxFloat64
    mtv := vector.Vector{}

    // Loop through edges of poly1
    for i := 0; i < len(poly1.Vertices); i++ {
        edge := poly1.Vertices[(i+1)%len(poly1.Vertices)].Sub(poly1.Vertices[i])
        normal := vector.Orthogonal(edge).Normalize()

        // Project polygons onto the normal
        min1, max1 := polygon.Project(*poly1, normal)
        min2, max2 := polygon.Project(*poly2, normal)

        // Check for overlap
        overlap := math.Min(max1, max2) - math.Max(min1, min2)
        if overlap <= 0 {
            return vector.Vector{} // No overlap, no MTV
        }

        // Update MTV if this overlap is smaller
        if overlap < minOverlap {
            minOverlap = overlap
            mtv = normal.Scale(overlap)
        }
    }

    // Loop through edges of poly2
    for i := 0; i < len(poly2.Vertices); i++ {
        edge := poly2.Vertices[(i+1)%len(poly2.Vertices)].Sub(poly2.Vertices[i])
        normal := vector.Orthogonal(edge).Normalize()

        // Project polygons onto the normal
        min1, max1 := polygon.Project(*poly1, normal)
        min2, max2 := polygon.Project(*poly2, normal)

        // Check for overlap
        overlap := math.Min(max1, max2) - math.Max(min1, min2)
        if overlap <= 0 {
            return vector.Vector{} // No overlap, no MTV
        }

        // Update MTV if this overlap is smaller
        if overlap < minOverlap {
            minOverlap = overlap
            mtv = normal.Scale(overlap)
        }
    }

    return mtv
}

// CirclePolygonCollision checks collision between a circle and a polygon (rectangle in this case).
func CirclePolygonCollision(circle *rigidbody.RigidBody, poly *polygon.Polygon) bool {
    // fmt.Println("Entering CirclePolygonCollision function")
    // defer fmt.Println("Exiting CirclePolygonCollision function")
    
    // Translate the circle's position into the coordinate system of the polygon
    circleX := circle.Position.X - poly.Position.X
    circleY := circle.Position.Y - poly.Position.Y

    // Closest point on the polygon to the circle
    closestX, closestY := poly.ClosestPoint(circleX, circleY)

    // Calculate the distance between the circle's center and this closest point
    distanceX := circleX - closestX
    distanceY := circleY - closestY

    // If the distance is less than the circle's radius, they are colliding
    return (distanceX*distanceX + distanceY*distanceY) <= (circle.Radius * circle.Radius)
}

// PerpendicularDistance calculates the perpendicular distance from a point to a line defined by two other points.
func PerpendicularDistance(point, linePoint1, linePoint2 vector.Vector) float64 {
    // fmt.Println("Entering PerpendicularDistance function")
    // defer fmt.Println("Exiting PerpendicularDistance function")
    
    // Vector from linePoint1 to linePoint2
    lineVector := linePoint2.Sub(linePoint1)

    // Vector from linePoint1 to the point
    pointVector := point.Sub(linePoint1)

    // Projection of pointVector onto lineVector
    projection := pointVector.InnerProduct(lineVector) / lineVector.Magnitude()

    // Calculate the perpendicular distance
    perpendicularDistance := pointVector.Magnitude() - projection

    return perpendicularDistance
}
