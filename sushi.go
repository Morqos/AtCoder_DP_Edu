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
var out = bufio.NewWriter(io.Writer(os_out))

var mod int = 1e9 + 7

const maxN int = 303

var prob = [maxN][maxN][maxN]float64{}
var ev = [maxN][maxN][maxN]float64{}

// dp[a][b][c] - probability to get state (a, b, c)

func main() {
	defer out.Flush()

	var n int

	fmt.Fscan(in, &n)

	var n_plates_per_n_sushis = [4]int{}
	// n_plates_per_n_sushis[i] - n plates containing i pieces of sushi

	for i := 0; i < n; i++ {
		var input int
		fmt.Fscan(in, &input)
		n_plates_per_n_sushis[input]++
	}

	prob[n_plates_per_n_sushis[1]][n_plates_per_n_sushis[2]][n_plates_per_n_sushis[3]] = 1

	for c := n; c >= 0; c-- {
		for b := n; b >= 0; b-- {
			for a := n; a >= 0; a-- {
				if a == 0 && b == 0 && c == 0 || a+b+c > n {
					continue
				}

				p_waste := float64(n-(a+b+c)) / float64(n)

				ev_waste := (p_waste / (1.0 - p_waste)) + 1.0
				ev[a][b][c] += ev_waste * prob[a][b][c]

				if a != 0 {
					p_go := float64(a) / float64((a + b + c))
					prob[a-1][b][c] += prob[a][b][c] * p_go
					ev[a-1][b][c] += ev[a][b][c] * p_go
				}
				if b != 0 {
					p_go := float64(b) / float64((a + b + c))
					prob[a+1][b-1][c] += prob[a][b][c] * p_go
					ev[a+1][b-1][c] += ev[a][b][c] * p_go
				}
				if c != 0 {
					p_go := float64(c) / float64((a + b + c))
					prob[a][b+1][c-1] += prob[a][b][c] * p_go
					ev[a][b+1][c-1] += ev[a][b][c] * p_go
				}
			}
		}
	}

	fmt.Printf("%.9f\n", ev[0][0][0])
}
