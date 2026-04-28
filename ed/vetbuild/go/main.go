package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	Root  []int
	Size  int
	Range int
}

func NewVector(val int) *Vector {
	return &Vector{
		Root:  make([]int, val),
		Size:  0,
		Range: val,
	}
}

func (v *Vector) PushBack(val int) {
	if v.Size == v.Range {
		newCap := v.Range * 2
		if newCap == 0 {
			newCap = 1
		}
		v.Reserve(newCap)
	}
	v.Root[v.Size] = val
	v.Size++
}

func (v *Vector) PopBack() error {
	if v.Size == 0 {
		return fmt.Errorf("vector is empty")
	}
	v.Size--
	return nil
}

func (v *Vector) Insert(index, value int) error {
	if index < 0 || index > v.Size {
		return fmt.Errorf("index out of range")
	}
	if v.Size == v.Range {
		newCap := v.Range * 2
		if newCap == 0 {
			newCap = 1
		}
		v.Reserve(newCap)
	}
	for i := v.Size; i > index; i-- {
		v.Root[i] = v.Root[i-1]
	}
	v.Root[index] = value
	v.Size++
	return nil
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.Size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < v.Size-1; i++ {
		v.Root[i] = v.Root[i+1]
	}
	v.Size--
	return nil
}

func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.Size; i++ {
		if v.Root[i] == value {
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1
}

func (v *Vector) Clear() {
	v.Size = 0
}

func (v *Vector) Capacity() int {
	return v.Range
}

func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.Size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.Root[index], nil
}

func (v *Vector) Set(index, value int) error {
	if index < 0 || index >= v.Size {
		return fmt.Errorf("index out of range")
	}
	v.Root[index] = value
	return nil
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity > v.Range {
		newData := make([]int, newCapacity)
		copy(newData, v.Root[:v.Size])
		v.Root = newData
		v.Range = newCapacity
	}
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.Size, v.Range)
}

func (v *Vector) String() string {
	if v.Size == 0 {
		return "[]"
	}
	return "[" + Join(v.Root[:v.Size], ", ") + "]"
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

	v := NewVector(0)
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
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
