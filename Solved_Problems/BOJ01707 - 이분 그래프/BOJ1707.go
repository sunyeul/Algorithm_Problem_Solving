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
	k := nextInt()

	for a := 0; a < k; a++ {
		v, e := nextInt(), nextInt()
		G := make(map[int][]int)
		for b := 0; b < e; b++ {
			i, j := nextInt(), nextInt()
			G[i] = append(G[i], j)
			G[j] = append(G[j], i)
		}
		grouped := make([]int, v+1)

		DFS := func(start int) bool {
			stack := []int{start}
			grouped[start] = 1
			for len(stack) > 0 {
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				group := grouped[curr]
				for _, nbhd := range G[curr] {
					if grouped[nbhd] != 0 {
						if grouped[nbhd] == group {
							return false
						}
						continue
					}
					grouped[nbhd] = -group
					stack = append(stack, nbhd)
				}
			}
			return true
		}

		ans := true

		for start := 1; start <= v; start++ {
			if grouped[start] == 0 {
				if DFS(start) == false {
					ans = false
					break
				}
			}
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
