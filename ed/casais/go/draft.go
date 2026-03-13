package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	descasados := make(map[int]int)
	pares := 0

	for i := 0; i < n; i++ {
		var animal int
		fmt.Scan(&animal)

		par := -animal

		if descasados[par] > 0 {
			pares++
			descasados[par]--
		} else {
			descasados[animal]++
		}
	}
	fmt.Println(pares)
}
