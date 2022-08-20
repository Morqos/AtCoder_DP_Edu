package main

import (
	"fmt"
)

var _map = make(map[int]int)

func main() {
	var n int
	var m int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)

	var adj = make([][]int, n+1)

	for i := 0; i < m; i++ {
		var from int
		var to int

		fmt.Scanf("%d", &from)
		fmt.Scanf("%d", &to)

		adj[from] = append(adj[from], to)
	}

	for i := 1; i < n+1; i++ {
		if _, ok := _map[i]; !ok {
			dfs(i, adj)
		}
	}
	answer := 0

	for _, el := range _map {
		answer = max(answer, el)
	}

	// println(answer)

	fmt.Printf("%d\n", answer)

}

func dfs(curr int, adj [][]int) int {
	if adj[curr] == nil {
		_map[curr] = 0
		return 0
	}
	if val, ok := _map[curr]; ok {
		return val
	}

	res := 0
	for _, next := range adj[curr] {
		if next == 0 {
			break
		}
		res = max(res, dfs(next, adj))
	}

	res++
	_map[curr] = res
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
