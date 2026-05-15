package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	capacity int
}

func NewMultiSet(cap int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, cap),
		capacity: cap,
	}
}

func (ms *MultiSet) search(value int) (bool, int) {
	low, high := 0, len(ms.data)-1
	found := false
	pos := 0

	for low <= high {
		mid := low + (high-low)/2
		if ms.data[mid] == value {
			found = true
			pos = mid
			low = mid + 1
		} else if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		return true, pos
	}
	return false, low
}

func (ms *MultiSet) insert(value int, index int) error {
	if index < 0 || index > len(ms.data) {
		return fmt.Errorf("fail: indice invalido")
	}
	ms.data = append(ms.data, 0)
	copy(ms.data[index+1:], ms.data[index:])
	ms.data[index] = value
	return nil
}

func (ms *MultiSet) Insert(value int) {
	found, idx := ms.search(value)
	if found {
		ms.insert(value, idx+1)
	} else {
		ms.insert(value, idx)
	}
}

func (ms *MultiSet) erase(index int) error {
	if index < 0 || index >= len(ms.data) {
		return fmt.Errorf("fail: indice invalido")
	}
	ms.data = append(ms.data[:index], ms.data[index+1:]...)
	return nil
}

func (ms *MultiSet) Erase(value int) error {
	found, idx := ms.search(value)
	if !found {
		return fmt.Errorf("fail: nao encontrado")
	}
	return ms.erase(idx)
}

func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}

func (ms *MultiSet) Count(value int) int {
	found, lastIdx := ms.search(value)
	if !found {
		return 0
	}

	count := 0
	for i := lastIdx; i >= 0 && ms.data[i] == value; i-- {
		count++
	}
	return count
}

func (ms *MultiSet) Unique() int {
	if len(ms.data) == 0 {
		return 0
	}
	count := 1
	for i := 1; i < len(ms.data); i++ {
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Clear() {
	ms.data = make([]int, 0, ms.capacity)
}

func (ms *MultiSet) String() string {
	return "[" + Join(ms.data, ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value := 0
			if len(args) > 1 {
				value, _ = strconv.Atoi(args[1])
			}
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if err := ms.Erase(value); err != nil {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
