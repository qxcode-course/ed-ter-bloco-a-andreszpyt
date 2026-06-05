package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {

	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(board, i, cols-1)
		}
	}
	for i := 0; i < cols; i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[rows-1][i] == 'O' {
			dfs(board, cols-1, i)
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}

}

func dfs(board [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	board[i][j] = 'Y'
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
