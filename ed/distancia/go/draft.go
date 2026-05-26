package main

import (
	"fmt"
)

func main() {
	var seq string
	var l int
	fmt.Scan(&seq, &l)

	s := []byte(seq)

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(s) {
			return true
		}

		if s[idx] != '.' {
			return dfs(idx + 1)
		}

		for d := 0; d <= l; d++ {
			char := byte('0' + d)
			pode := true

			inicio := idx - l
			if inicio < 0 {
				inicio = 0
			}
			for i := inicio; i < idx; i++ {
				if s[i] == char {
					pode = false
					break
				}
			}

			if pode {
				fim := idx + l
				if fim > len(s)-1 {
					fim = len(s) - 1
				}
				for i := idx + 1; i <= fim; i++ {
					if s[i] == char {
						pode = false
						break
					}
				}
			}

			if pode {
				s[idx] = char
				if dfs(idx + 1) {
					return true
				}
				s[idx] = '.'
			}

		}

		return false
	}

	dfs(0)
	fmt.Println(string(s))
}
