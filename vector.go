package simul8

import (
	"fmt"
	"math"
)

// Vector2 represents a vector in 2D space
type Vector2 struct {
	X, Y float64
}

// NewVector2 creates a new 2D vector
func NewVector2(x, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}

// Add adds two 2D vectors
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

// Subtract subtracts one 2D vector from another
func (v Vector2) Subtract(other Vector2) Vector2 {
	return Vector2{X: v.X - other.X, Y: v.Y - other.Y}
}

// Scale scales a 2D vector by a scalar
func (v Vector2) Scale(scalar float64) Vector2 {
	return Vector2{X: v.X * scalar, Y: v.Y * scalar}
}

// Dot calculates the dot product of two 2D vectors
func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Magnitude calculates the magnitude of the 2D vector
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalise normalises the 2D vector to a unit vector
func (v Vector2) Normalise() Vector2 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2{0, 0}
	}
	return Vector2{X: v.X / mag, Y: v.Y / mag}
}

// String returns a string representation of the 2D vector
func (v Vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

// Vector3 represents a vector in 3D space
type Vector3 struct {
	X, Y, Z float64
}

// NewVector3 creates a new 3D vector
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}

// Add adds two 3D vectors
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// Subtract subtracts one 3D vector from another
func (v Vector3) Subtract(other Vector3) Vector3 {
	return Vector3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

// Scale scales a 3D vector by a scalar
func (v Vector3) Scale(scalar float64) Vector3 {
	return Vector3{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

// Dot calculates the dot product of two 3D vectors
func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Magnitude calculates the magnitude of the 3D vector
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalise normalises the 3D vector to a unit vector
func (v Vector3) Normalise() Vector3 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector3{0, 0, 0}
	}
	return Vector3{X: v.X / mag, Y: v.Y / mag, Z: v.Z / mag}
}

// String returns a string representation of the 3D vector
func (v Vector3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

// Vector4 represents a vector in 4D space
type Vector4 struct {
	X, Y, Z, W float64
}

// NewVector4 creates a new 4D vector
func NewVector4(x, y, z, w float64) Vector4 {
	return Vector4{X: x, Y: y, Z: z, W: w}
}

// Add adds two 4D vectors
func (v Vector4) Add(other Vector4) Vector4 {
	return Vector4{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z, W: v.W + other.W}
}

// Subtract subtracts one 4D vector from another
func (v Vector4) Subtract(other Vector4) Vector4 {
	return Vector4{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z, W: v.W - other.W}
}

// Scale scales a 4D vector by a scalar
func (v Vector4) Scale(scalar float64) Vector4 {
	return Vector4{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar, W: v.W * scalar}
}

// Dot calculates the dot product of two 4D vectors
func (v Vector4) Dot(other Vector4) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

// Magnitude calculates the magnitude of the 4D vector
func (v Vector4) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalise normalises the 4D vector to a unit vector
func (v Vector4) Normalise() Vector4 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector4{0, 0, 0, 0}
	}
	return Vector4{X: v.X / mag, Y: v.Y / mag, Z: v.Z / mag, W: v.W / mag}
}

// String returns a string representation of the 4D vector
func (v Vector4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}
