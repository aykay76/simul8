package simul8

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	rows, columns := 3, 3
	matrix := NewMatrix(rows, columns)

	if matrix.rows != rows || matrix.columns != columns {
		t.Errorf("Expected matrix dimensions to be %dx%d, got %dx%d", rows, columns, matrix.rows, matrix.columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix.data[i][j] != 0 {
				t.Errorf("Expected matrix element at (%d, %d) to be 0, got %f", i, j, matrix.data[i][j])
			}
		}
	}
}

func TestMatrix_Add(t *testing.T) {
	m1 := &Matrix{
		rows:    2,
		columns: 2,
		data: [][]float64{
			{1, 2},
			{3, 4},
		},
	}

	m2 := &Matrix{
		rows:    2,
		columns: 2,
		data: [][]float64{
			{5, 6},
			{7, 8},
		},
	}

	expected := &Matrix{
		rows:    2,
		columns: 2,
		data: [][]float64{
			{6, 8},
			{10, 12},
		},
	}

	result, err := m1.Add(m2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, got %v", expected, result)
	}
}

func TestMatrix_Subtract(t *testing.T) {
	m1 := &Matrix{
		rows:    2,
		columns: 2,
		data: [][]float64{
			{5, 6},
			{7, 8},
		},
	}

	m2 := &Matrix{
		rows:    2,
		columns: 2,
		data: [][]float64{
			{1, 2},
			{3, 4},
		},
	}

	expected := &Matrix{
		rows:    2,
		columns: 2,
		data: [][]float64{
			{4, 4},
			{4, 4},
		},
	}

	result, err := m1.Subtract(m2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, got %v", expected, result)
	}
}

func TestMatrix_Multiply(t *testing.T) {
	m1 := &Matrix{
		rows:    2,
		columns: 3,
		data: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	m2 := &Matrix{
		rows:    3,
		columns: 2,
		data: [][]float64{
			{7, 8},
			{9, 10},
			{11, 12},
		},
	}

	expected := &Matrix{
		rows:    2,
		columns: 2,
		data: [][]float64{
			{58, 64},
			{139, 154},
		},
	}

	result, err := m1.Multiply(m2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, got %v", expected, result)
	}
}

func TestMatrix_Transpose(t *testing.T) {
	m := &Matrix{
		rows:    2,
		columns: 3,
		data: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	expected := &Matrix{
		rows:    3,
		columns: 2,
		data: [][]float64{
			{1, 4},
			{2, 5},
			{3, 6},
		},
	}

	result := m.Transpose()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, got %v", expected, result)
	}
}
