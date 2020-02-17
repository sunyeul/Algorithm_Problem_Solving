package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func scan() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func isGood(s string) bool {
	cnt := 0
	for _, e := range s {
		switch string(e) {
		case "a":
			cnt++
		case "e":
			cnt++
		case "i":
			cnt++
		case "o":
			cnt++
		case "u":
			cnt++
		}
	}
	if cnt > 0 && len(s)-cnt > 1 {
		return true
	}
	return false
}

func password(k, l int, arr []string, s string) {
	if k == l && isGood(s) {
		fmt.Println(s)
		return
	}
	for i := 0; i < len(arr); i++ {
		a := arr[i]
		switch {
		case len(s) == 0:
			password(k+1, l, arr[i+1:], s+a)
		case len(s) > 0 && s[len(s)-1] < []byte(a)[0]:
			password(k+1, l, arr[i+1:], s+a)
		}
	}
}

func main() {
	defer writer.Flush()

	L, C := scan(), scan()
	var letters []string
	for i := 0; i < C; i++ {
		scanner.Scan()
		letters = append(letters, scanner.Text())
	}
	sort.Strings(letters)
	password(0, L, letters, "")
}
