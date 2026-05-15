package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	root *Node
	tail *Node
}

func NewLList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Size() int {
	count := 0
	for n := ll.root; n != nil; n = n.next {
		count++
	}
	return count
}

func (ll *LinkedList) PushBack(value int) {
	newNode := &Node{value: value}
	if ll.root == nil {
		ll.root = newNode
		ll.tail = newNode
		return
	}
	newNode.prev = ll.tail
	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *LinkedList) PushFront(value int) {
	newNode := &Node{value: value}
	if ll.root == nil {
		ll.root = newNode
		ll.tail = newNode
		return
	}
	newNode.next = ll.root
	ll.root.prev = newNode
	ll.root = newNode
}

func (ll *LinkedList) PopBack() {
	if ll.tail == nil {
		return
	}
	if ll.root == ll.tail {
		ll.root = nil
		ll.tail = nil
		return
	}
	ll.tail = ll.tail.prev
	ll.tail.next = nil
}

func (ll *LinkedList) PopFront() {
	if ll.root == nil {
		return
	}
	if ll.root == ll.tail {
		ll.root = nil
		ll.tail = nil
		return
	}
	ll.root = ll.root.next
	ll.root.prev = nil
}

func (ll *LinkedList) Clear() {
	ll.root = nil
	ll.tail = nil
}

func (ll *LinkedList) String() string {
	var res []string
	for n := ll.root; n != nil; n = n.next {
		res = append(res, strconv.Itoa(n.value))
	}
	return "[" + strings.Join(res, ", ") + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		fmt.Println(line)

		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
