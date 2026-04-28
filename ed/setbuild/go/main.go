package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Set struct {
	Root  []map[int]int
	Size  int
	Range int
}

func NewSet(val int) *Set {
	root := make([]map[int]int, val)
	for i := range root {
		root[i] = make(map[int]int)
	}
	return &Set{
		Root:  root,
		Size:  0,
		Range: val,
	}
}

func (s *Set) hash(val int) int {
	return val % s.Range
}

func (s *Set) Insert(val int) {
	h := s.hash(val)
	if _, exists := s.Root[h][val]; !exists {
		s.Root[h][val] = 1
		s.Size++
	}
}

func (s *Set) Erase(val int) error {
	h := s.hash(val)
	if _, exists := s.Root[h][val]; exists {
		delete(s.Root[h], val)
		s.Size--
		return nil
	}
	return fmt.Errorf("value not found")
}

func (s *Set) Contains(val int) bool {
	h := s.hash(val)
	_, exists := s.Root[h][val]
	return exists
}

func (s *Set) Clear() {
	for i := range s.Root {
		s.Root[i] = make(map[int]int)
	}
	s.Size = 0
}

func (s *Set) String() string {
	if s.Size == 0 {
		return "[]"
	}

	elements := make([]int, 0, s.Size)
	for _, bucket := range s.Root {
		for key := range bucket {
			elements = append(elements, key)
		}
	}

	sort.Ints(elements)

	strElements := make([]string, len(elements))
	for i, val := range elements {
		strElements[i] = strconv.Itoa(val)
	}

	return "[" + strings.Join(strElements, ", ") + "]"
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
