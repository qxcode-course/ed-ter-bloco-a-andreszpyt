package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
}

func dfs(board [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	board[i][j] = 'X'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
