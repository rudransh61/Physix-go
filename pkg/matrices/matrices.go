package matrices

import (
	"errors"
	"math"
)

var Identity3 = [][]float64{
	{1.0, 0.0, 0.0},
	{0.0, 1.0, 0.0},
	{0.0, 0.0, 1.0},
}

var Identity2 = [][]float64{
	{1.0, 0.0},
	{0.0, 1.0},
}

// RotationX returns a 3x3 rotation matrix for rotating around the X axis by the given angle (in radians).
func RotationX(angle float64) [][]float64 {
    cos := math.Cos(angle)
    sin := math.Sin(angle)

    return [][]float64{
        {1, 0, 0},
        {0, cos, -sin},
        {0, sin, cos},
    }
}

// RotationY returns a 3x3 rotation matrix for rotating around the Y axis by the given angle (in radians).
func RotationY(angle float64) [][]float64 {
    cos := math.Cos(angle)
    sin := math.Sin(angle)

    return [][]float64{
        {cos, 0, sin},
        {0, 1, 0},
        {-sin, 0, cos},
    }
}

// RotationZ returns a 3x3 rotation matrix for rotating around the Z axis by the given angle (in radians).
func RotationZ(angle float64) [][]float64 {
    cos := math.Cos(angle)
    sin := math.Sin(angle)

    return [][]float64{
        {cos, -sin, 0},
        {sin, cos, 0},
        {0, 0, 1},
    }
}

// Add performs matrix addition and returns the result.
func Add(matrixA, matrixB [][]float64) ([][]float64, error) {
	if !isValidMatrix(matrixA) || !isValidMatrix(matrixB) || len(matrixA) != len(matrixB) || len(matrixA[0]) != len(matrixB[0]) {
		return nil, errors.New("invalid matrices for addition")
	}

	result := make([][]float64, len(matrixA))
	for i := range result {
		result[i] = make([]float64, len(matrixA[0]))
		for j := range result[i] {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}

	return result, nil
}

// Subtract performs matrix subtraction and returns the result.
func Subtract(matrixA, matrixB [][]float64) ([][]float64, error) {
	if !isValidMatrix(matrixA) || !isValidMatrix(matrixB) || len(matrixA) != len(matrixB) || len(matrixA[0]) != len(matrixB[0]) {
		return nil, errors.New("invalid matrices for subtraction")
	}

	result := make([][]float64, len(matrixA))
	for i := range result {
		result[i] = make([]float64, len(matrixA[0]))
		for j := range result[i] {
			result[i][j] = matrixA[i][j] - matrixB[i][j]
		}
	}

	return result, nil
}

// Multiply performs matrix multiplication and returns the result.
func Multiply(matrixA, matrixB [][]float64) [][]float64 {
	if !isValidMatrix(matrixA) || !isValidMatrix(matrixB) || len(matrixA[0]) != len(matrixB) {
		return [][]float64{}
	}

	result := make([][]float64, len(matrixA))
	for i := range result {
		result[i] = make([]float64, len(matrixB[0]))
		for j := range result[i] {
			for k := range matrixA[0] {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}

// isValidMatrix checks if a given matrix is valid (non-empty).
func isValidMatrix(matrix [][]float64) bool {
	return len(matrix) > 0 && len(matrix[0]) > 0
}
