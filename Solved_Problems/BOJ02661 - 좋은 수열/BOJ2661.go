package main

import (
	"fmt"
)

var stop bool = false

func isGoodSeq(str string) bool {
	l := len(str)
	for i := 1; i <= l/2; i++ {
		if str[l-i*2:l-i] == str[l-i:l] {
			return false
		}
	}
	return true
}

func GoodSeq(k, n int, str string) {
	if k == n {
		fmt.Println(str)
		stop = true
		return
	}

	for _, i := range []string{"1", "2", "3"} {
		tmp := str + i
		if isGoodSeq(tmp) {
			GoodSeq(k+1, n, tmp)
		}
		if stop {
			return
		}
	}
	return
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	GoodSeq(0, n, "")
}
