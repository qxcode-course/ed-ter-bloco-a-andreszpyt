package main

import "fmt"

func dfs(n int, c float64) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += c
		sum /= 2.0
	}
	return sum
}

func main() {
	var n int
	var c float64
	fmt.Scan(&n, &c)
	fmt.Printf("%.2f\n", dfs(n, c))
}
