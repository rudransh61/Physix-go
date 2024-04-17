// main.go
package main

import (
	"fmt"
	// "time"

	"github.com/rudransh61/Physix-go/pkg/matrices"

	// "github.com/rudransh61/Physix-go/dynamics/collision"
)

func main() {
	matrixA := [][]float64{
		{1, 2},
		{3, 4},
	}

	matrixB := [][]float64{
		{5, 6},
		{7, 8},
	}

	// Perform addition
	addResult, err := matrices.Add(matrixA, matrixB)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Addition Result:")
		printMatrix(addResult)
	}

	// Perform subtraction
	subtractResult, err := matrices.Subtract(matrixA, matrixB)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Subtraction Result:")
		printMatrix(subtractResult)
	}

	// Perform multiplication
	multiplyResult := matrices.Multiply(matrixA, matrixB)
	fmt.Println("Multiplication Result:")
	printMatrix(multiplyResult)
}

func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%7.2f ", value)
		}
		fmt.Println()
	}
}