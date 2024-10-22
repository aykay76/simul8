package simul8

import (
	"errors"
	"math"
)

type Matrix struct {
	rows    int
	columns int
	data    [][]float64
}

// NewMatrix creates a new matrix with the given number of rows and columns
func NewMatrix(rows, columns int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, columns)
	}
	return &Matrix{rows: rows, columns: columns, data: data}
}

// Add adds two matrices and returns the result
func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.rows != other.rows || m.columns != other.columns {
		return nil, errors.New("matrices must have the same dimensions to add")
	}
	result := NewMatrix(m.rows, m.columns)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			result.data[i][j] = m.data[i][j] + other.data[i][j]
		}
	}
	return result, nil
}

// Subtract subtracts another matrix from the current matrix and returns the result
func (m *Matrix) Subtract(other *Matrix) (*Matrix, error) {
	if m.rows != other.rows || m.columns != other.columns {
		return nil, errors.New("matrices must have the same dimensions to subtract")
	}
	result := NewMatrix(m.rows, m.columns)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			result.data[i][j] = m.data[i][j] - other.data[i][j]
		}
	}
	return result, nil
}

// Multiply multiplies two matrices and returns the result
func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if m.columns != other.rows {
		return nil, errors.New("number of columns of the first matrix must equal the number of rows of the second matrix")
	}
	result := NewMatrix(m.rows, other.columns)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.columns; j++ {
			for k := 0; k < m.columns; k++ {
				result.data[i][j] += m.data[i][k] * other.data[k][j]
			}
		}
	}
	return result, nil
}

// Transpose returns the transpose of the matrix
func (m *Matrix) Transpose() *Matrix {
	result := NewMatrix(m.columns, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			result.data[j][i] = m.data[i][j]
		}
	}
	return result
}

// create from euler angles
func (m *Matrix) FromEulerAngles(angles Vector3) {
	cosX, sinX := math.Cos(angles.X), math.Sin(angles.X)
	cosY, sinY := math.Cos(angles.Y), math.Sin(angles.Y)
	cosZ, sinZ := math.Cos(angles.Z), math.Sin(angles.Z)

	m.data[0][0] = cosY * cosZ
	m.data[0][1] = cosY * sinZ
	m.data[0][2] = -sinY

	m.data[1][0] = sinX*sinY*cosZ - cosX*sinZ
	m.data[1][1] = sinX*sinY*sinZ + cosX*cosZ
	m.data[1][2] = sinX * cosY

	m.data[2][0] = cosX*sinY*cosZ + sinX*sinZ
	m.data[2][1] = cosX*sinY*sinZ - sinX*cosZ
	m.data[2][2] = cosX * cosY
}

// add determinant method
func (m *Matrix) Determinant() float64 {
	a := m.data[0][0]
	b := m.data[0][1]
	c := m.data[0][2]
	d := m.data[1][0]
	e := m.data[1][1]
	f := m.data[1][2]
	g := m.data[2][0]
	h := m.data[2][1]
	i := m.data[2][2]

	return a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
}

// invert the matrix
func (m *Matrix) Invert() {
	det := m.Determinant()
	if det == 0 {
		return
	}
	invDet := 1 / det

	a := m.data[0][0]
	b := m.data[0][1]
	c := m.data[0][2]
	d := m.data[1][0]
	e := m.data[1][1]
	f := m.data[1][2]
	g := m.data[2][0]
	h := m.data[2][1]
	i := m.data[2][2]

	m.data[0][0] = (e*i - f*h) * invDet
	m.data[0][1] = (c*h - b*i) * invDet
	m.data[0][2] = (b*f - c*e) * invDet
	m.data[1][0] = (f*g - d*i) * invDet
	m.data[1][1] = (a*i - c*g) * invDet
	m.data[1][2] = (c*d - a*f) * invDet
	m.data[2][0] = (d*h - e*g) * invDet
	m.data[2][1] = (g*b - a*h) * invDet
	m.data[2][2] = (a*e - b*d) * invDet
}

// multiply by vector
func (m *Matrix) MultiplyVector(v Vector3) Vector3 {
	return Vector3{
		X: m.data[0][0]*v.X + m.data[0][1]*v.Y + m.data[0][2]*v.Z,
		Y: m.data[1][0]*v.X + m.data[1][1]*v.Y + m.data[1][2]*v.Z,
		Z: m.data[2][0]*v.X + m.data[2][1]*v.Y + m.data[2][2]*v.Z,
	}
}

// multiply by float
func (m *Matrix) MultiplyScalar(s float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.data[i][j] *= s
		}
	}
}

// divide by float
func (m *Matrix) DivideScalar(s float64) {
	if s == 0 {
		return
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.data[i][j] /= s
		}
	}
}
