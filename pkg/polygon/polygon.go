package polygon

import (
	"physix/pkg/vector"
	"physix/pkg/rigidbody"
	// "math"
)

//Polygon 2d
type Polygon struct {
	rigidbody.RigidBody
	Vertices []vector.Vector
}

// NewPolygon creates a new polygon with given properties.
func NewPolygon(vertices []vector.Vector, mass float64, IsMovable bool) *Polygon {
	polygon := &Polygon{
		RigidBody: rigidbody.RigidBody{
			Position:  calculateCentroid(vertices),
			Velocity:  vector.Vector{X: 0, Y: 0},
			Force:     vector.Vector{X: 0, Y: 0},
			Mass:      mass,
			Shape:     "polygon",
			IsMovable: IsMovable,
		},
		Vertices: vertices,
	}
	return polygon
}

// calculateCentroid calculates the centroid of a polygon given its vertices.
func calculateCentroid(vertices []vector.Vector) vector.Vector {
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
	p.Position = calculateCentroid(p.Vertices)
}



// IMPART Impulse on a body
func (rb *Polygon) ApplyImpulse(impulse vector.Vector) {
    // Calculate the change in velocity using impulse and mass
    deltaV := vector.Vector{
        X: impulse.X / rb.Mass,
        Y: impulse.Y / rb.Mass,
    }

	// Update everything
    for i := range rb.Vertices {
        rb.Vertices[i].X += deltaV.X
        rb.Vertices[i].Y += deltaV.Y
    }
}