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

func main() {
	defer out.Flush()

	var n int

	fmt.Fscan(in, &n)

	var probHead = make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &probHead[i])
	}

	var pHeads = make([][]float64, n+1)

	for i := 0; i <= n; i++ {
		pHeads[i] = make([]float64, n+1)
	}

	pHeads[0][0] = 1

	for j := 1; j <= n; j++ {
		_pHead, _pTail := probHead[j-1], 1-probHead[j-1]

		pHeads[0][j] = _pTail * pHeads[0][j-1]

		for i := 1; i <= j; i++ {
			pHeads[i][j] = _pHead*pHeads[i-1][j-1] + _pTail*pHeads[i][j-1]
		}
	}

	var ans float64
	for i := n/2 + 1; i <= n; i++ {
		ans += pHeads[i][n]
	}

	fmt.Printf("%.9f", ans)

}
