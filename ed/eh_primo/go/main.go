package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	_, _ = x, div
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%div-1 != 0 {
		return false
	}

	eh_primo(x/div, div)
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
