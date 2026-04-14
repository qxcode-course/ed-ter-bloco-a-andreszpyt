package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	id := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&id[i])
	}

	var m int
	fmt.Scan(&m)

	queue := make(map[int]bool, m)
	for i := 0; i < m; i++ {
		var id int
		fmt.Scan(&id)
		queue[id] = true
	}

	for i := 0; i < n; i++ {
		if !queue[id[i]] {
			fmt.Printf("%d ", id[i])
		}
	}
	fmt.Println()
}
