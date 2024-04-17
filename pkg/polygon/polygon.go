package polygon

import (
	"github.com/rudransh61/Physix-go/pkg/vector"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"math"
)

//Polygon 2d
type Polygon struct {
	rigidbody.RigidBody
	Vertices []vector.Vector
    Rotation float64 // Added rotation property
    Torque   float64 // Added torque property

}

// NewPolygon creates a new polygon with given properties.
func NewPolygon(vertices []vector.Vector, mass float64, IsMovable bool) *Polygon {
	polygon := &Polygon{
		RigidBody: rigidbody.RigidBody{
			Position:  CalculateCentroid(vertices),
			Velocity:  vector.Vector{X: 0, Y: 0},
			Force:     vector.Vector{X: 0, Y: 0},
			Mass:      mass,
			Shape:     "polygon",
			IsMovable: IsMovable,
			Restitution : 1.0,
		},
        Rotation: 0, // Initial rotation is set to 0
        Torque:   0, // Initial torque is set to 0
		Vertices: vertices,
	}
	return polygon
}

// CalculateCentroid calculates the centroid of a polygon given its vertices.
func CalculateCentroid(vertices []vector.Vector) vector.Vector {
	var centroid vector.Vector
	for _, v := range vertices {
		centroid.X += v.X
		centroid.Y += v.Y
	}
	centroid.X /= float64(len(vertices))
	centroid.Y /= float64(len(vertices))
	return centroid
}

//Update position of polygon
func (p *Polygon) UpdatePosition(){
	p.Position = CalculateCentroid(p.Vertices)
    p.Rotation = p.Torque / p.Mass // Update rotation based on torque and mass
}



// IMPART Impulse on a body
func (rb *Polygon) ApplyImpulse(impulse vector.Vector) {
    // Calculate the change in velocity using impulse and mass
    change_velocity := impulse.Scale(1/rb.Mass);
    rb.Velocity = rb.Velocity.Add(change_velocity)

    rb.Rotation += rb.Torque / rb.Mass // Update rotation based on torque and mass
}

// Project calculates the projection of a polygon onto a given axis.
func Project(p Polygon, axis vector.Vector) (float64, float64) {
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

// Move adjusts the position of the polygon by the given displacement vector.
func (p *Polygon) Move(displacement vector.Vector) {
    for i := range p.Vertices {
        p.Vertices[i] = p.Vertices[i].Add(displacement)
    }
}


// ClosestPoint finds the closest point on the polygon to a given point (x, y).
func (poly *Polygon) ClosestPoint(x, y float64) (float64, float64) {
    minDistanceSquared := math.MaxFloat64
    closestX, closestY := 0.0, 0.0

    for i := 0; i < len(poly.Vertices); i++ {
        vertexX := poly.Vertices[i].X + poly.Position.X
        vertexY := poly.Vertices[i].Y + poly.Position.Y
        dx := x - vertexX
        dy := y - vertexY
        distanceSquared := dx*dx + dy*dy
        if distanceSquared < minDistanceSquared {
            minDistanceSquared = distanceSquared
            closestX = vertexX
            closestY = vertexY
        }
    }

    return closestX, closestY
}