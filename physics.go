package simul8

// PhysicsObject represents an object with physical properties.
type PhysicsObject struct {
	Position, Velocity, Acceleration Vector2
	Mass                             float64
}

// Update updates the position and velocity of the physics object based on its acceleration and the time step.
func (obj *PhysicsObject) Update(dt float64) {
	obj.Velocity = obj.Velocity.Add(obj.Acceleration.Multiply(dt))
	obj.Position = obj.Position.Add(obj.Velocity.Multiply(dt))
}
