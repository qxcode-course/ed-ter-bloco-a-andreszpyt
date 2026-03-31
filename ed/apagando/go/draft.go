package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	queue := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&queue[i])
	}
	fmt.Scan(&m)
	leftQueue := make(map[int]bool, n)
	for i := 0; i < m; i++ {
		var k int
		fmt.Scan(&k)
		leftQueue[k] = true
	}

	for i := 0; i < n; i++ {
		if !leftQueue[queue[i]] {
			if i == n-1 {
				fmt.Printf("%d \n", queue[i])
				return
			}
			fmt.Printf("%d ", queue[i])
		}
	}
	fmt.Println()
}
