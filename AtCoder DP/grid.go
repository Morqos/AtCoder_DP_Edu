package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var os_in = os.Stdin
var os_out = os.Stdout

var in = bufio.NewReader(io.Reader(os_in))
var out = bufio.NewWriter(io.Writer(os.Stdout))

var mod int = 1e9 + 7

type cell struct {
	row int
	col int
}

func main() {
	defer out.Flush()

	var n int
	var m int

	fmt.Fscan(in, &n, &m)

	matrix := []string{}

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		matrix = append(matrix, s)
	}

	var dp = make([][]int, n)
	var visited = make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		visited[i] = make([]bool, m)
	}
	dp[0][0] = 1

	queue := []cell{cell{0, 0}}

	for len(queue) > 0 {

		curr_cell := queue[0]
		queue = queue[1:]

		// print_matrix(dp, n, m)

		if visited[curr_cell.row][curr_cell.col] {
			continue
		}
		visited[curr_cell.row][curr_cell.col] = true

		adjCells := getAdjCells(matrix, curr_cell, n, m)

		for _, _cell := range adjCells {
			dp[_cell.row][_cell.col] = (dp[_cell.row][_cell.col] + dp[curr_cell.row][curr_cell.col]) % mod
			queue = append(queue, _cell)
		}

	}

	// println(answer)

	fmt.Printf("%d\n", dp[n-1][m-1])

}

func getAdjCells(matrix []string, curr_cell cell, n, m int) []cell {

	var adjCells = []cell{}

	if curr_cell.row != n-1 && matrix[curr_cell.row+1][curr_cell.col] != '#' {
		adjCells = append(adjCells, cell{curr_cell.row + 1, curr_cell.col})
	}

	if curr_cell.col != m-1 && matrix[curr_cell.row][curr_cell.col+1] != '#' {
		adjCells = append(adjCells, cell{curr_cell.row, curr_cell.col + 1})
	}

	return adjCells
}

func print_matrix(dp [][]int, n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}
