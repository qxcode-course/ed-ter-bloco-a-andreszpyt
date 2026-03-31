package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x <= 0 {
		return -x
	}
	return x
}

func getMen(vet []int) []int {
	mans := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		if vet[i] >= 0 {
			mans = append(mans, vet[i])
		}
	}
	return mans
}

func getCalmWomen(vet []int) []int {
	calmWomen := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 && vet[i] > -10 {
			calmWomen = append(calmWomen, vet[i])
		}
	}
	return calmWomen
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func sortStress(vet []int) []int {
	sort.Slice(vet, func(i, j int) bool {
		return abs(vet[i]) < abs(vet[j])
	})
	return vet
}

func reverse(vet []int) []int {
	if len(vet) < 2 {
		return vet
	}
	newvet := make([]int, len(vet))
	copy(newvet, vet)
	for i, j := 0, len(newvet)-1; i < j; i, j = i+1, j-1 {
		newvet[i], newvet[j] = newvet[j], newvet[i]
	}
	return newvet
}

func unique(vet []int) []int {
	seen := make(map[int]struct{})
	finalVet := make([]int, 0)

	for _, k := range vet {
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			finalVet = append(finalVet, k)
		}
	}
	return finalVet
}

func repeated(vet []int) []int {
	count := make(map[int]int)
	result := make([]int, 0)
	for _, v := range vet {
		count[v]++
		if count[v] > 1 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
