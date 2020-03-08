package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	defer writer.Flush()
	n := nextInt()

	W := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			W[i] = append(W[i], nextInt())
		}
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<uint(n))
	}

	var travelling func(prev, visited int) int
	travelling = func(prev, visited int) int {
		if memo[prev][visited] != 0 {
			return memo[prev][visited]
		}

		if visited == (1<<uint(n))-1 {
			if W[prev][0] != 0 {
				return W[prev][0]
			}
			return math.MaxInt32

		}

		res := math.MaxInt32
		for i := 0; i < n; i++ {
			if (visited>>uint(i))&1 == 1 {
				continue
			}

			if W[prev][i] == 0 {
				continue
			}
			cost := travelling(i, visited+(1<<uint(i))) + W[prev][i]
			if cost < res {
				res = cost
			}
		}
		memo[prev][visited] = res
		return res
	}
	res := travelling(0, 1)
	fmt.Println(res)
}
