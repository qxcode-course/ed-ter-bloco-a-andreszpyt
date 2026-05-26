package main

import "fmt"

type Node struct {
	char rune
	prev *Node
	next *Node
}

func main() {
	var input string
	fmt.Scan(&input)

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
				prev := cursor.prev
				next := cursor.next
				prev.next = next
				if next != nil {
					next.prev = prev
				}
				cursor = prev
			}
		case 'D':
			if cursor.next != nil {
				del := cursor.next
				next := del.next
				cursor.next = next
				if next != nil {
					next.prev = cursor
				}
			}
		default:
			newNode := &Node{char: char, prev: cursor, next: cursor.next}

			if cursor.next != nil {
				cursor.next.prev = newNode
			}
			cursor.next = newNode
			cursor = newNode
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
