package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var f, _ = os.Open("input.in")
var reader = bufio.NewReader(f)
var scanner = bufio.NewScanner(f)

// var reader = bufio.NewReader(os.Stdin)
// var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func init() {
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func c(n, k int) int {
	switch {
	case n >= k:
		return n - k
	case k == 1:
		return 1
	case k%2 == 1:
		return 1 + min(c(n, k-1), c(n, k+1))
	default:
		return min(k-n, 1+c(n, k/2))
	}
}

func main() {
	n, k := nextInt(), nextInt()
	fmt.Println(c(n, k))
}
