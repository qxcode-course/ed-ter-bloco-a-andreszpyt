package main

import "fmt"

func main() {
	var n int
	var partes = 0
	fmt.Scan(&n)
	peoples := make([]int, n)
	mapa := make(map[int]int, 0)

	for i := range peoples {
		fmt.Scan(&peoples[i])
		mapa[peoples[i]]++
		if mapa[-peoples[i]] == 1 {
			partes++
			mapa[-peoples[i]]--
			mapa[peoples[i]]--
		}
	}
	fmt.Println(partes)
}
