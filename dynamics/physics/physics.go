package physix

import (
    "github.com/rudransh61/Physix-go/pkg/vector"
    "github.com/rudransh61/Physix-go/pkg/rigidbody"
    "github.com/rudransh61/Physix-go/pkg/polygon"
    "github.com/rudransh61/Physix-go/dynamics/collision"
    "github.com/rudransh61/Physix-go/pkg/broadphase"
    "math"
)

// Update updates the physics simulation.
func Update(objects []interface{}, dt float64) {
    // Create a spatial hash with an appropriate cell size
    spatialHash := broadphase.NewSpatialHash(50.0)

    // Insert objects into the spatial hash
    for _, obj := range objects {
        switch obj := obj.(type) {
        case *rigidbody.RigidBody:
            spatialHash.Add(obj, obj.Position)
        case *polygon.Polygon:
            spatialHash.Add(obj, obj.Position)
        }
    }

    // Broad-phase collision detection
    potentialCollisions := make(map[interface{}]map[interface{}]bool)
    for _, obj := range objects {
        var position vector.Vector
        switch obj := obj.(type) {
        case *rigidbody.RigidBody:
            position = obj.Position
        case *polygon.Polygon:
            position = obj.Position
        }

        nearbyObjects := spatialHash.Query(position)
        for _, other := range nearbyObjects {
            if obj == other {
                continue
            }
            if _, exists := potentialCollisions[obj]; !exists {
                potentialCollisions[obj] = make(map[interface{}]bool)
            }
            potentialCollisions[obj][other] = true
        }
    }

    // Narrow-phase collision detection and response
    for obj, colliders := range potentialCollisions {
        for collider := range colliders {
            switch obj := obj.(type) {
            case *rigidbody.RigidBody:
                switch collider := collider.(type) {
                case *rigidbody.RigidBody:
                    if collision.RectangleCollided(obj, collider) {
                        collision.BounceOnCollision(obj, collider, 1.0)
                    } else if collision.CircleCollided(obj, collider) {
                        collision.BounceOnCollision(obj, collider, 1.0)
                    }
                case *polygon.Polygon:
                    if collision.CirclePolygonCollision(obj, collider) {
                        collision.ResolveCollision(collider, collider, 1.0, 0.1)
                    }
                }
            case *polygon.Polygon:
                switch collider := collider.(type) {
                case *rigidbody.RigidBody:
                    if collision.CirclePolygonCollision(collider, obj) {
                        collision.ResolveCollision(obj, obj, 1.0, 0.1)
                    }
                case *polygon.Polygon:
                    if collision.PolygonCollision(obj, collider) {
                        collision.ResolveCollision(obj, collider, 1.0, 0.1)
                    }
                }
            }
        }
    }

    // Update positions and velocities based on forces and impulses
    for _, obj := range objects {
        switch obj := obj.(type) {
        case *rigidbody.RigidBody:
            ApplyForce(obj,obj.Force, dt)
            UpdateRotation(obj, dt)
        case *polygon.Polygon:
            ApplyForcePolygon(obj,obj.Force, dt)
        }
    }
}

// ApplyForcePolygon applies force to a polygon and rotates every vertex about the centroid.
func ApplyForcePolygon(pg *polygon.Polygon, force vector.Vector, dt float64) {
    if pg.IsMovable {
        // Use Newton's second law: F = ma -> a = F/m
        pg.Force = force
        acceleration := vector.Vector{
            X: pg.Force.X / pg.Mass,
            Y: pg.Force.Y / pg.Mass,
        }

        // Update velocity using acceleration and time step
        pg.Velocity.X += acceleration.X * dt
        pg.Velocity.Y += acceleration.Y * dt

        // Update position using velocity and time step for each vertex
        for i := range pg.Vertices {
            pg.Vertices[i].X += pg.Velocity.X * dt
            pg.Vertices[i].Y += pg.Velocity.Y * dt
        }

        // Calculate centroid
        centroid := polygon.CalculateCentroid(pg.Vertices)

        // Rotate each vertex about the centroid
        for i := range pg.Vertices {
            // Translate vertex to the origin
            translatedX := pg.Vertices[i].X - centroid.X
            translatedY := pg.Vertices[i].Y - centroid.Y

            // Rotate the vertex
            rotatedX := translatedX*math.Cos(pg.Rotation) - translatedY*math.Sin(pg.Rotation)
            rotatedY := translatedX*math.Sin(pg.Rotation) + translatedY*math.Cos(pg.Rotation)

            // Translate the vertex back to its original position
            pg.Vertices[i].X = rotatedX + centroid.X
            pg.Vertices[i].Y = rotatedY + centroid.Y
        }

        pg.UpdatePosition()
    }
}


// ApplyForce applies a force to a rigid body.
func ApplyForce(rb *rigidbody.RigidBody, force vector.Vector, dt float64) {
	if rb.IsMovable {
		// Use Newton's second law: F = ma -> a = F/m
		// rb.Force = rb.Force.Add(force)

		rb.Force = force
		acceleration := rb.Force.Scale(1 / rb.Mass)

		// Update velocity using acceleration and time step
		rb.Velocity = rb.Velocity.Add(acceleration.Scale(dt))

		// Update position using velocity and time step
		rb.Position = rb.Position.Add(rb.Velocity.Scale(dt))
	}
}

// UpdateRotation updates the rotation of the rigid body based on its angular velocity.
func UpdateRotation(rb *rigidbody.RigidBody, dt float64) {
	// I := 1.0 // Define MOI
	// // Update Angular velocity
	// if rb.Shape == "Circle" {
	// 	I = 0.5 * rb.Mass * rb.Radius * rb.Radius // Moment of Inertia of Circular Disk
	// } else if rb.Shape == "Rectangle" {
	// 	I = 0.16668 * rb.Mass * rb.Width * rb.Height // // Moment of Inertia of Plank defined by us.
	// 	// 1/24 *b*l*m
	// }
	// rb.AngularAcceleration = rb.Torque / I
	// rb.AngularVelocity += rb.AngularAcceleration*dt
	// // Update rotation based on angular velocity
	// angle := rb.AngularVelocity * dt
	// rb.Position.X = math.Cos(angle)*rb.Position.X - math.Sin(angle)*rb.Position.Y
	// rb.Position.Y = math.Sin(angle)*rb.Position.X + math.Cos(angle)*rb.Position.Y

}