package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {

	cols, rows := len(matrix), len(matrix[0])

	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
	}

	max := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			atualMax := dfs(matrix, memo, i, j, -1)
			if atualMax > max {
				max = atualMax
			}
		}
	}
}

func dfs(board [][]int, memo [][]int, i, j, valPrevio int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] <= valPrevio {
		return
	}

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
