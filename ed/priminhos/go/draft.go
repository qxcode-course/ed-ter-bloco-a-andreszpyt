package main

import "fmt"

func eh_primo(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2-1 != 0 {
		return false
	}

	eh_primo(x / 2)
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
	fmt.Println(s)
	return
}
