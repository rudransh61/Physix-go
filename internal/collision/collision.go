// internal/collision/collision.go
package collision

import (
	"physix/pkg/rigidbody"
	"physix/pkg/polygon"
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


// Polygon collision detection SAT implementation.
// Link : https://dyn4j.org/2010/01/sat/

// Project calculates the projection of a polygon onto a given axis.
func Project(p polygon.Polygon, axis vector.Vector) (float64, float64) {
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
	return !(max1 < min2 || max2 < min1)
}

// Collides checks if two polygons are colliding using the SAT algorithm.
func PolygonCollision(poly1, poly2 polygon.Polygon) bool {
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

	return true
}

// CalculateCollisionNormal calculates the collision normal between two polygons
func CalculateCollisionNormal(poly1, poly2 polygon.Polygon) vector.Vector {
	// Implement collision normal calculation
	// For simplicity, let's calculate the average normal of all edges
	// Calculate the edges of both polygons
	edges1 := calculateEdges(poly1.Vertices)
	edges2 := calculateEdges(poly2.Vertices)

	// Calculate normals of edges
	normals1 := calculateNormals(edges1)
	normals2 := calculateNormals(edges2)

	// Find the average normal of the edges
	averageNormal := averageVector(normals1...)
	averageNormal = averageNormal.Add(averageVector(normals2...))
	averageNormal = averageNormal.Normalize()

	return averageNormal
}

// ResolveCollision resolves collision between two polygons
func ResolveCollision(poly1, poly2 *polygon.Polygon, normal vector.Vector) {
    // Calculate relative velocity
    relativeVelocity := poly2.Velocity.Sub(poly1.Velocity)

    // Calculate relative velocity in terms of the normal direction
    velocityAlongNormal := relativeVelocity.InnerProduct(normal)

    // If velocities are separating, do not resolve collision
    if velocityAlongNormal > 0 {
        return
    }

    // Calculate restitution (bounciness)
    restitution := 2.0 // Example restitution coefficient

    // Calculate impulse scalar
    impulseScalar := -(1 + restitution) * velocityAlongNormal / (1 / poly1.Mass + 1 / poly2.Mass)

    // Apply impulse to the polygons
    impulse := normal.Scale(impulseScalar)

    poly1.Velocity = poly1.Velocity.Add(impulse.Scale(1 / poly1.Mass)).Scale(2)
    poly2.Velocity = poly2.Velocity.Sub(impulse.Scale(1 / poly2.Mass)).Scale(2)
}

// Helper function to calculate the dot product of two vectors
func dotProduct(v1, v2 vector.Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Helper function to calculate edges from polygon vertices
func calculateEdges(vertices []vector.Vector) []vector.Vector {
	edges := make([]vector.Vector, len(vertices))
	for i := 0; i < len(vertices); i++ {
		nextIndex := (i + 1) % len(vertices)
		edges[i] = vertices[nextIndex].Sub(vertices[i])
	}
	return edges
}

// Helper function to calculate normals from edges
func calculateNormals(edges []vector.Vector) []vector.Vector {
	normals := make([]vector.Vector, len(edges))
	for i, edge := range edges {
		// Calculate normal by rotating edge by 90 degrees counter-clockwise
		normals[i] = vector.Vector{X: -edge.Y, Y: edge.X}
	}
	return normals
}

// Helper function to calculate the average vector from multiple vectors
func averageVector(vectors ...vector.Vector) vector.Vector {
	sum := vector.Vector{}
	for _, v := range vectors {
		sum = sum.Add(v)
	}
	return sum.Scale(1 / float64(len(vectors)))
}
