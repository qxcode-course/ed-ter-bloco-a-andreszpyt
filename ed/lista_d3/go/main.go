package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func merge(lla *LList, llb *LList) interface{} {
	merged := NewLList()
	p1 := lla.root.next
	p2 := llb.root.next

	for p1 != lla.root && p2 != llb.root {
		if p1.Value <= p2.Value {
			merged.PushBack(p1.Value)
			p1 = p1.next
		} else {
			merged.PushBack(p2.Value)
			p2 = p2.next
		}
	}

	for p1 != lla.root {
		merged.PushBack(p1.Value)
		p1 = p1.next
	}

	for p2 != llb.root {
		merged.PushBack(p2.Value)
		p2 = p2.next
	}

	return merged
}

func reverse(lla *LList) {
	var prev *Node = nil
	atual := lla.root

	for atual.next != nil {
		next := atual.next
		atual.next = atual.prev
		prev = atual
		atual = next
	}
	lla.root = prev
}

func addsorted(lla *LList, value int) {
	NewNode := &Node{Value: value}

	if lla.root == nil || lla.root.Value > value {
		NewNode.next = lla.root
		NewNode.prev = lla.root
		lla.root = NewNode
		return
	}

	ActualNode := lla.root.next

	for ActualNode != nil && ActualNode.Value < value {
		ActualNode = ActualNode.next
	}

	NewNode.next = ActualNode.next
	ActualNode.next = NewNode
}

func equals(lla *LList, llb *LList) bool {
	if lla.size != llb.size || lla.root != llb.root {
		return false
	}

	for i := 0; i != lla.size; i++ {
		lla.root = lla.root.next
		llb.root = llb.root.next
		if lla.root.Value != llb.root.Value {
			return false
		}
	}
	return true
}
