package main

import (
	"fmt"
)

func eh_primo(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}

	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	aux := 0

	for i := 0; aux < n; i++ {
		if eh_primo(i) {
			s[aux] = i
			aux++
		}
	}

	fmt.Printf("[")
	for i, val := range s {
		if i == len(s)-1 {
			fmt.Printf("%d]\n", val)
		} else {
			fmt.Printf("%d, ", val)
		}
	}
}
