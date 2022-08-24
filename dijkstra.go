// Extra
// Dijkstra Problem: https://codeforces.com/problemset/problem/20/C

package main

import (
	"bufio"
	. "container/heap"
	"fmt"
	"io"
	"os"
)

var os_in = os.Stdin
var os_out = os.Stdout

var in = bufio.NewReader(io.Reader(os_in))
var out = bufio.NewWriter(io.Writer(os.Stdout))

type node struct {
	value  int
	weight int64
}

type heap []node

func (h heap) Len() int              { return len(h) }
func (h heap) Less(i, j int) bool    { return h[i].weight < h[j].weight }
func (h heap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *heap) Push(v interface{})   { *h = append(*h, v.(node)) }
func (h *heap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func main() {
	defer out.Flush()

	var n int
	var m int

	fmt.Fscan(in, &n, &m)

	var adj = make([][]node, n+1)

	// var visited = make([]bool, n+1)

	for i := 0; i < m; i++ {
		var from int
		var to int
		var weight int64

		fmt.Fscan(in, &from, &to, &weight)

		// println(i)

		adj[from] = append(adj[from], node{to, weight})
		adj[to] = append(adj[to], node{from, weight})
	}

	// dfs(adj, &visited, 1)

	dijkstra(n, adj)
}

func dfs(adj [][]node, visited *[]bool, curr int) {

	fmt.Printf("%d\n", curr)
	(*visited)[curr] = true

	for _, node := range adj[curr] {
		if !(*visited)[node.value] {
			dfs(adj, visited, node.value)
		}
	}
}

func dijkstra(n int, adj [][]node) {
	const inf int64 = 1e18

	var distance = make([]int64, n+1)
	var predecessor = make([]int, n+1)

	for i := 0; i <= n; i++ {
		distance[i] = inf
	}
	distance[1] = 0

	queue := heap{{value: 1, weight: 0}}

	for len(queue) > 0 {
		// Need to optimize the next 2 lines
		curr_node := Pop(&queue).(node)
		curr_node_value := curr_node.value

		for _, next_node := range adj[curr_node_value] {
			next_node_value := next_node.value
			next_node_weight := next_node.weight
			if distance[curr_node_value]+next_node_weight < distance[next_node_value] {
				predecessor[next_node_value] = curr_node_value
				distance[next_node_value] = distance[curr_node_value] + next_node_weight
				Push(&queue, node{next_node_value, distance[next_node_value]})
			}
		}
	}

	if distance[n] == inf {
		fmt.Fprint(out, -1)
		return
	}

	last_node := n
	path := []int{}
	for last_node != 0 {
		path = append(path, last_node)
		last_node = predecessor[last_node]
	}

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Fprint(out, path[i], " ")
	}
}
