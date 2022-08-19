package main

import (
	"fmt"
)

var string1 string
var string2 string

type cell struct {
	howMany int
	prevRow int
	prevCol int
}

func main() {
	fmt.Scanln(&string1)
	fmt.Scanln(&string2)

	len1 := len(string1)
	len2 := len(string2)

	matrix := make([][]cell, len1+1)

	for i := 0; i < len1+1; i++ {
		matrix[i] = make([]cell, len2+1)
	}

	// for i := 0; i < len(string1); i++ {
	// 	for j := 0; j < len(string2); j++ {
	// 		print(matrix[i][j].howMany)
	// 	}
	// 	println()
	// }

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if string1[i] == string2[j] {
				maxSelf(&matrix[i+1][j+1], cell{howMany: matrix[i][j].howMany + 1, prevRow: i, prevCol: j})
			}

			maxSelf(&matrix[i+1][j], cell{howMany: matrix[i][j].howMany, prevRow: i, prevCol: j})
			maxSelf(&matrix[i][j+1], cell{howMany: matrix[i][j].howMany, prevRow: i, prevCol: j})

		}
	}

	answer := cell{0, 0, 0}
	var ansString string

	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			maxSelf(&answer, cell{matrix[i][j].howMany, i, j})
		}
	}

	for {
		row := answer.prevRow
		col := answer.prevCol
		if row == 0 || col == 0 {
			break
		}
		if matrix[row][col].prevRow == row-1 && matrix[row][col].prevCol == col-1 {
			ansString = string(string1[row-1]) + ansString
		}
		answer = matrix[row][col]
	}

	fmt.Println(string(ansString))
}

func maxSelf(x *cell, y cell) {
	if x.howMany < y.howMany {
		*x = y
	}
}
