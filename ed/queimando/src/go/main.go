package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	r int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()

	nl := len(grid)
	nr := len(grid[0])

	stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		current := stack.Pop()
		cl, cc := current.r, current.c

		if cl < 0 || cl >= nl || cc < 0 || cc >= nr {
			continue
		}
		if grid[cl][cc] == '#' {
			grid[cl][cc] = 'o'

			stack.Push(Pos{cl + 1, cc})
			stack.Push(Pos{cl - 1, cc})
			stack.Push(Pos{cl, cc + 1})
			stack.Push(Pos{cl, cc - 1})
		}
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
