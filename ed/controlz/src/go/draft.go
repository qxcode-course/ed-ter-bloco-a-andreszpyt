package main

import "fmt"

type Node struct {
	char rune
	prev *Node
	next *Node
}

type Command struct {
	op   rune
	node *Node
}

func main() {
	var input string
	fmt.Scan(&input)
	var undoStack []Command
	var redoStack []Command

	head := &Node{}
	cursor := head

	for _, char := range input {

		if char == 'R' {
			char = '\n'
		}

		switch char {
		case '>':
			if cursor.next != nil {
				cursor = cursor.next
			}
		case '<':
			if cursor != head {
				cursor = cursor.prev
			}
		case 'B':
			if cursor != head {
				delNode := cursor
				prev := cursor.prev
				next := cursor.next
				prev.next = next
				if next != nil {
					next.prev = prev
				}
				cursor = prev

				undoStack = append(undoStack, Command{'B', delNode})
				redoStack = nil
			}
		case 'D':
			if cursor.next != nil {
				del := cursor.next
				next := del.next
				cursor.next = next
				if next != nil {
					next.prev = cursor
				}
				undoStack = append(undoStack, Command{'D', del})
				redoStack = nil
			}
		case 'Z':
			if len(undoStack) > 0 {
				action := undoStack[len(undoStack)-1]
				undoStack = undoStack[:len(undoStack)-1]

				switch action.op {
				case 'I':
					node := action.node
					node.prev.next = node.next
					if node.next != nil {
						node.next.prev = node.prev
					}
					cursor = node.prev
				case 'B':
					node := action.node
					node.prev.next = node
					if node.next != nil {
						node.next.prev = node
					}
					cursor = node
				case 'D':
					node := action.node
					node.prev.next = node
					if node.next != nil {
						node.next.prev = node
					}
				}
				redoStack = append(redoStack, action)
			}
		case 'Y':
			if len(redoStack) > 0 {
				action := redoStack[len(redoStack)-1]
				redoStack = redoStack[:len(redoStack)-1]
				switch action.op {
				case 'I':
					node := action.node
					node.prev.next = node
					if node.next != nil {
						node.next.prev = node
					}
					cursor = node
				case 'B':
					node := action.node
					node.prev.next = node.next
					if node.next != nil {
						node.next.prev = node.prev
					}
					cursor = node.prev
				case 'D':
					node := action.node
					node.prev.next = node.next
					if node.next != nil {
						node.next.prev = node.prev
					}
				}
				undoStack = append(undoStack, action)
			}
		default:
			newNode := &Node{char: char, prev: cursor, next: cursor.next}
			if cursor.next != nil {
				cursor.next.prev = newNode
			}
			cursor.next = newNode
			cursor = newNode
			undoStack = append(undoStack, Command{'I', newNode})
			redoStack = nil
		}
	}

	var output []rune
	actual := head
	for actual != nil {
		if actual != head {
			output = append(output, actual.char)
		}
		if actual == cursor {
			output = append(output, '|')
		}
		actual = actual.next
	}

	fmt.Println(string(output))
}
