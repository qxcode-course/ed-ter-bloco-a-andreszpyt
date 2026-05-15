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

func (ll *LinkedList) Front() *Node {
	return ll.root
}

func (ll *LinkedList) Back() *Node {
	return ll.tail
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (ll *LinkedList) Search(value int) *Node {
	for node := ll.root; node != nil; node = node.next {
		if node.value == value {
			return node
		}
	}

	return nil
}

func (ll *LinkedList) Insert(node *Node, value int) {
	newNode := &Node{value: value}

	newNode.next = node
	newNode.prev = node.prev

	if node.next != nil {
		node.prev.next = newNode
	} else {
		ll.tail = newNode
	}

	node.next = newNode
}

func (ll *LinkedList) Remove(node *Node) {
	if node == nil {
		return
	}

	if node == ll.root {
		ll.PopFront()
		return
	}

	if node == ll.tail {
		ll.PopBack()
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
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

		case "walk":
			fmt.Print("[ ")

			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}

			fmt.Print("]\n[ ")

			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}

			fmt.Println("]")

		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])

			node := ll.Search(oldvalue)

			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}

		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])

			node := ll.Search(oldvalue)

			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}

		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])

			node := ll.Search(oldvalue)

			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}

		case "end":
			return

		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
