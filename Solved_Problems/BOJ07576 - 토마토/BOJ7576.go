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

type pos struct {
	i, j int
}

type status struct {
	pos
	day int
}

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

	m, n := nextInt(), nextInt()

	box := make([][]int, n, n)
	queue := make([]status, 0, m*n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := nextInt()
			box[i] = append(box[i], tmp)
			if tmp == 1 {
				queue = append(queue, status{pos{i, j}, 0})
			}
		}
	}

	move := []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	MaxDay := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.day > MaxDay {
			MaxDay = curr.day
		}

		for _, d := range move {
			next := pos{curr.i + d.i, curr.j + d.j}
			if next.i < 0 || next.i >= n || next.j < 0 || next.j >= m {
				continue
			}

			if box[next.i][next.j] != 0 {
				continue
			}
			box[next.i][next.j] = 1
			queue = append(queue, status{pos{next.i, next.j}, curr.day + 1})
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if box[i][j] == 0 {
				fmt.Println(-1)
				return
			}
		}
	}
	fmt.Println(MaxDay)
	return
}
