package main

import (
	"bufio";
	"fmt";
	"os";
	"strconv";
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

func minslice(array [3]int) int {
	var min int = array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
	}
	return min
}

func main() {
	defer writer.Flush()

	n := scan()
	cost := [3]int{scan(), scan(), scan()}

	for i := 0; i < n-1; i++ {
		var costCopy [3]int = cost

		R, G, B := scan(), scan(), scan()

		cost[0] = min(costCopy[1]+R, costCopy[2]+R)
		cost[1] = min(costCopy[0]+G, costCopy[2]+G)
		cost[2] = min(costCopy[0]+B, costCopy[1]+B)
	}
	fmt.Println(minslice(cost))
}
