package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (q *Queue[T]) Enqueue(value T) {
	NewNode := &Node[T]{
		Value: value,
	}
	if q.tail != nil {
		q.tail.next = NewNode
		q.tail = q.tail.next
	} else {
		q.head = NewNode
		q.tail = NewNode
	}
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head == nil {
		var v T
		return v, false
	}
	dequeue := q.head
	q.head = q.head.next
	return dequeue.Value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var v T
		return v, false
	}
	return q.head.Value, true
}

func (q *Queue[T]) Size() int {
	head := q.head
	for i := 0; head != nil; i++ {
		if head.next == q.tail {
			return i
		}
		head = head.next
	}
	return -1
}

func (q *Queue[T]) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
}

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			for _, arg := range args[1:] {
				value, _ := strconv.Atoi(arg)
				queue.Enqueue(value)
			}
		case "pop":
			if _, ok := queue.Dequeue(); !ok {
				fmt.Println("falha: fila vazia")
			}
		case "peek":
			if value, ok := queue.Peek(); ok {
				fmt.Println(value)
			} else {
				fmt.Println("falha: fila vazia")
			}
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
