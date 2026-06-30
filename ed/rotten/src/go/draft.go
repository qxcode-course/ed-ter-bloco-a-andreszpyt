package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	result := oranges(grid)
	fmt.Println(result)
}

func oranges(grid [][]int) interface{} {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	goods := 0
	queue := [][]int{}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			}
			if grid[r][c] == 1 {
				goods++
			}
		}
	}

	if goods == 0 {
		return 0
	}
	minutes := 0
	directions := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for len(queue) > 0 && goods > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			atual := queue[0]
			queue = queue[1:]
			r, c := atual[0], atual[1]

			for _, direction := range directions {
				nr, nc := r+direction[0], c+direction[1]

				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					goods--
					queue = append(queue, []int{nr, nc})
				}
			}
		}
		minutes++

	}

	if goods > 0 {
		return -1
	}
	return minutes
}
