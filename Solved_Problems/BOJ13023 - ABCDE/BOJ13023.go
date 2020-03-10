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

func main() {
	defer writer.Flush()
	var visited [2000]int
	n, m := nextInt(), nextInt()

	adj := make([][]int, n)
	for j := 0; j < m; j++ {
		u, v := nextInt(), nextInt()
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(i, depth int) bool
	dfs = func(i, depth int) bool {
		if depth == 4 {
			return true
		}
		visited[i] = 1
		for _, v := range adj[i] {
			if visited[v] == 1 {
				continue
			}
			if dfs(v, depth+1) {
				return true
			}
		}
		visited[i] = 0
		return false
	}

	for i := 0; i < n; i++ {
		if dfs(i, 0) {
			fmt.Println(1)
			return
		}
	}
	fmt.Println(0)
	return
}
