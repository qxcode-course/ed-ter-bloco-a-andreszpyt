package main

import "fmt"

func div(v int) {
	if v == 0 {
		return
	}

	value := v / 2
	resto := v % 2

	div(value)
	fmt.Printf("%d %d\n", value, resto)
}

func main() {
	var s int
	fmt.Scan(&s)
	div(s)
}
