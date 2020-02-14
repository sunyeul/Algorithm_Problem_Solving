package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func init() {
	scanner.Split(bufio.ScanWords)
}

func scan() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	n, k := scan(), scan()
	dp := make([]int, k+1)

	for m := 1; m <= n; m++ {
		w, v := scan(), scan()
		for j := k; j >= w; j-- {
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}
	fmt.Println(dp[k])
}
