package physix

import "physix/pkg/vector"

const dt = 0.1

type Particle struct {
	Position, PrevPosition, Velocity vector.Vector
	InverseMass                     float64
}

// NewParticle creates a new particle with the given position and mass.
func NewParticle(position vector.Vector, mass float64) *Particle {
	return &Particle{
		Position:      position,
		PrevPosition:  position,
		InverseMass:   1 / mass,
	}
}

// ApplyForce applies a force to the particle.
func (p *Particle) ApplyForce(force vector.Vector) {
	acceleration := force.Scale(p.InverseMass)
	p.Velocity = p.Velocity.Add(acceleration.Scale(dt))
}

// Modify the Particle struct to include Verlet integration in the Update method
func (p *Particle) Update(dt float64) {
    currentPosition := p.Position
    p.Position = currentPosition.Scale(2).Sub(p.PrevPosition).Add(p.Velocity.Scale(dt * dt))
    p.PrevPosition = currentPosition
}