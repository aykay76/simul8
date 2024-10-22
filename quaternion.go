package simul8

import (
	"fmt"
	"math"
)

// Quaternion represents a quaternion with real (w) and imaginary (x, y, z) parts.
type Quaternion struct {
	W, X, Y, Z float64
}

// New creates a new quaternion.
func New(w, x, y, z float64) Quaternion {
	return Quaternion{W: w, X: x, Y: y, Z: z}
}

// Add adds two quaternions.
func (q Quaternion) Add(q2 Quaternion) Quaternion {
	return Quaternion{
		W: q.W + q2.W,
		X: q.X + q2.X,
		Y: q.Y + q2.Y,
		Z: q.Z + q2.Z,
	}
}

// Sub subtracts two quaternions.
func (q Quaternion) Sub(q2 Quaternion) Quaternion {
	return Quaternion{
		W: q.W - q2.W,
		X: q.X - q2.X,
		Y: q.Y - q2.Y,
		Z: q.Z - q2.Z,
	}
}

// Mul multiplies two quaternions.
func (q Quaternion) Mul(q2 Quaternion) Quaternion {
	return Quaternion{
		W: q.W*q2.W - q.X*q2.X - q.Y*q2.Y - q.Z*q2.Z,
		X: q.W*q2.X + q.X*q2.W + q.Y*q2.Z - q.Z*q2.Y,
		Y: q.W*q2.Y - q.X*q2.Z + q.Y*q2.W + q.Z*q2.X,
		Z: q.W*q2.Z + q.X*q2.Y - q.Y*q2.X + q.Z*q2.W,
	}
}

// Norm returns the norm (magnitude) of the quaternion.
func (q Quaternion) Norm() float64 {
	return math.Sqrt(q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z)
}

// Conjugate returns the conjugate of the quaternion.
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{W: q.W, X: -q.X, Y: -q.Y, Z: -q.Z}
}

// Inverse returns the inverse of the quaternion.
func (q Quaternion) Inverse() Quaternion {
	norm := q.Norm()
	if norm == 0 {
		panic("cannot invert a quaternion with zero norm")
	}
	conjugate := q.Conjugate()
	return Quaternion{
		W: conjugate.W / (norm * norm),
		X: conjugate.X / (norm * norm),
		Y: conjugate.Y / (norm * norm),
		Z: conjugate.Z / (norm * norm),
	}
}

// String returns the string representation of the quaternion.
func (q Quaternion) String() string {
	return fmt.Sprintf("(%f + %fi + %fj + %fk)", q.W, q.X, q.Y, q.Z)
}
