package main

import "fmt"

func main() {
	var n int
	var e int
	fmt.Scan(&n)
	fmt.Scan(&e)

	pessoa := make([]int, n)

	for i := range pessoa {
		pessoa[i] = i + 1
	}

	index := 0
	for i, p := range pessoa {
		if p == e {
			index = i
			break
		}
	}

	for len(pessoa) > 0 {
		fmt.Print("[ ")
		for i, j := range pessoa {
			if i == index {
				fmt.Printf("%d> ", j)
				continue
			}
			fmt.Printf("%d ", j)
		}
		fmt.Println("]")

		if len(pessoa) == 1 {
			break
		}

		kill := (index + 1) % len(pessoa)
		pessoa = append(pessoa[0:kill], pessoa[kill+1:]...)
		index = kill % len(pessoa)
	}
}
