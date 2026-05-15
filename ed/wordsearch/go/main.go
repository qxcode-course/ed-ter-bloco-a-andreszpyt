package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	nl := len(grid)
	nc := len(grid[0])

	var dfs func(l, c, idx int) bool
	dfs = func(l, c, idx int) bool {
		if idx == len(word) {
			return true
		}

		if l < 0 || l >= nl || c < 0 || c >= nc || grid[l][c] != word[idx] {
			return false
		}

		aux := grid[l][c]
		grid[l][c] = '#'
		res := dfs(l-1, c, idx+1) || dfs(l+1, c, idx+1) || dfs(l, c-1, idx+1) || dfs(l, c+1, idx+1)
		grid[l][c] = aux

		return res
	}

	for l := 0; l < nl; l++ {
		for c := 0; c < nc; c++ {
			if grid[l][c] == word[0] && dfs(l, c, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
