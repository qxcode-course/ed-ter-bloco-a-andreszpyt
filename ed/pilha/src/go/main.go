package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	data     []T
	capacity int
	size     int
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		data:     make([]T, 0, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
	s.size++
}

func (s *Stack[T]) Pop() error {
	if s.size == 0 {
		return fmt.Errorf("stack is empty")
	}
	s.data = s.data[:s.size-1]
	s.size--
	return nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.size == 0 {
		return zero, fmt.Errorf("stack is empty")
	}
	return s.data[s.size-1], nil
}

func (s *Stack[T]) String() string {
	output := ""
	for i := range cap(s.data) {
		if i != 0 {
			output += ", "
		}
		if i < len(s.data) {
			output += fmt.Sprintf("%v", s.data[i])
		} else {
			output += "_"
		}
	}
	return output
}

func (s *Stack[T]) Size() any {
	return s.size
}

func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
	s.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewStack[int](10)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			cap, _ := strconv.Atoi(parts[1])
			v = NewStack[int](cap)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Push(value)
			}
		case "debug":
			fmt.Println(v)
		case "top":
			top, err := v.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(v.Size())
		case "pop":
			err := v.Pop()
			if err != nil {
				fmt.Println(err)
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
