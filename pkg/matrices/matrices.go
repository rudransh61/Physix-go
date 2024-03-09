package matrices

import (
	"errors"
)

// Add performs matrix addition and returns the result.
func Add(matrixA, matrixB [][]int) ([][]int, error) {
	if !isValidMatrix(matrixA) || !isValidMatrix(matrixB) || len(matrixA) != len(matrixB) || len(matrixA[0]) != len(matrixB[0]) {
		return nil, errors.New("invalid matrices for addition")
	}

	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixA[0]))
		for j := range result[i] {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}

	return result, nil
}

// Subtract performs matrix subtraction and returns the result.
func Subtract(matrixA, matrixB [][]int) ([][]int, error) {
	if !isValidMatrix(matrixA) || !isValidMatrix(matrixB) || len(matrixA) != len(matrixB) || len(matrixA[0]) != len(matrixB[0]) {
		return nil, errors.New("invalid matrices for subtraction")
	}

	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixA[0]))
		for j := range result[i] {
			result[i][j] = matrixA[i][j] - matrixB[i][j]
		}
	}

	return result, nil
}

// Multiply performs matrix multiplication and returns the result.
func Multiply(matrixA, matrixB [][]int) ([][]int, error) {
	if !isValidMatrix(matrixA) || !isValidMatrix(matrixB) || len(matrixA[0]) != len(matrixB) {
		return nil, errors.New("invalid matrices for multiplication")
	}

	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixB[0]))
		for j := range result[i] {
			for k := range matrixA[0] {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result, nil
}

// isValidMatrix checks if a given matrix is valid (non-empty).
func isValidMatrix(matrix [][]int) bool {
	return len(matrix) > 0 && len(matrix[0]) > 0
}
