package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func occurr(vet []int) []Pair {
	count := make(map[int]int)

	for i := 0; i < len(vet); i++ {
		abs := abs(vet[i])
		count[abs]++
	}
	var pairs []Pair

	for k, v := range count {
		pairs = append(pairs, Pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].One < pairs[j].One
	})
	return pairs
}

func teams(vet []int) []Pair {
	count := make(map[int]int)
	for i := 0; i < len(vet)-1; i++ {
		if vet[i+1] == vet[i] {
			count[vet[i]]++
		}
	}
	var pairs []Pair
	for k, v := range count {
		pairs = append(pairs, Pair{v, k})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].One < pairs[j].One
	})
	return pairs
}

func mnext(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		if vet[i] >= 0 && (vet[i-1] < 0 || vet[i+1] < 0) {
			vet[i] = 1
		}
	}
	return vet
}

func alone(vet []int) []int {
	alone := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		if vet[i] >= 0 && !(vet[i-1] < 0 || vet[i+1] < 0) {
			alone = append(alone, vet[i])
		}
	}
	return alone
}

func couple(vet []int) int {
	count := make(map[int]int)
	for _, v := range vet {
		count[v]++
	}

	couples := 0
	for k, v := range count {
		if k > 0 {
			if negV, exists := count[-k]; exists {
				if v < negV {
					couples += v
				} else {
					couples += negV
				}
			}
		}
	}
	return couples
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	newVet := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		isPresent := false
		for j := 0; j < len(posList); j++ {
			if posList[j] == i {
				isPresent = true
			}
		}
		if !isPresent {
			newVet = append(newVet, vet[i])
		}
	}
	return newVet
}

func clear(vet []int, value int) []int {
	newVet := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		if vet[i] != value {
			newVet = append(newVet, vet[i])
		}
	}
	return newVet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
