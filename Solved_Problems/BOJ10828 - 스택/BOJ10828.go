package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var f, _ = os.Open("input.in")

// var reader = bufio.NewReader(os.Stdin)
// var scanner = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReader(f)
var scanner = bufio.NewScanner(f)
var writer = bufio.NewWriter(os.Stdout)

func init() {
	scanner.Split(bufio.ScanWords)
}

func scan() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

type (
	// Node : nested list
	Node struct {
		next  *Node
		value int
	}
	// Stack : stack
	Stack struct {
		top  *Node
		size int
	}
)

// Push : 정수 x를 스택에 넣는 연산이다.
func (s *Stack) Push(x int) {
	s.top = &Node{s.top, x}
	s.size++
}

// Pop : 스택에서 가장 위에 있는 정수를 빼고, 그 수를 출력한다. 만약 스택에 들어있는 정수가 없는 경우에는 -1을 출력한다.
func (s *Stack) Pop() int {
	if s.size == 0 {
		return -1
	}
	s.size--
	e := s.top
	s.top = e.next
	return e.value
}

// Size : 스택에 들어있는 정수의 개수를 출력한다.
func (s *Stack) Size() int {
	return s.size
}

// Empty : 스택이 비어있으면 1, 아니면 0을 출력한다.
func (s *Stack) Empty() int {
	empty := (s.size == 0)
	if empty {
		return 1
	}
	return 0
}

// Top : 스택의 가장 위에 있는 정수를 출력한다. 만약 스택에 들어있는 정수가 없는 경우에는 -1을 출력한다.
func (s *Stack) Top() int {
	if s.size == 0 {
		return -1
	}
	return s.top.value
}

func main() {

	defer writer.Flush()
	n := scan()

	var stack Stack

	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()

		switch str {
		case "push":
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			stack.Push(v)
		case "pop":
			fmt.Fprintf(writer, "%d\n", stack.Pop())
		case "size":
			fmt.Fprintf(writer, "%d\n", stack.Size())
		case "empty":
			fmt.Fprintf(writer, "%d\n", stack.Empty())
		case "top":
			fmt.Fprintf(writer, "%d\n", stack.Top())
		}
	}
}
