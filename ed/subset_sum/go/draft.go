package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var dfs func(idx, sum int) bool
	dfs = func(idx, sum int) bool {
		if sum == k {
			return true
		}
		if idx >= n || sum > k {
			return false
		}

		return dfs(idx+1, sum+arr[idx]) || dfs(idx+1, sum)
	}

	if dfs(0, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
