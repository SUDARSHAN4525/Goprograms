package main

import "fmt"

func main() {
	var rows, cols int
	fmt.Print("Enter rows and columns of the matrices: ")
	fmt.Scan(&rows, &cols)

	mat1 := make([][]int, rows)
	mat2 := make([][]int, rows)
	result := make([][]int, rows)

	fmt.Println("Enter elements of first matrix:")
	for i := 0; i < rows; i++ {
		mat1[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Scan(&mat1[i][j])
		}
	}

	fmt.Println("Enter elements of second matrix:")
	for i := 0; i < rows; i++ {
		mat2[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Scan(&mat2[i][j])
		}
	}

	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < cols; k++ {
				result[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	fmt.Println("Resultant Matrix:")
	for _, row := range result {
		fmt.Println(row)
	}
}
