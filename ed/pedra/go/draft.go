package main

import "fmt"

func main() {
	var n int
	menor := 1000
	var win bool
	var x int

	fmt.Scan(&n)

	d := make([]int, n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		if a < 10 || b < 10 {
			d[i] = -1
			continue
		}

		if a > b {
			d[i] = a - b
		}

		if b > a {
			d[i] = b - a
		}

		if d[i] != -1 && d[i] < menor {
			menor = d[i]
			x = i
			win = true
		}
	}

	if win == true {
		fmt.Println(x)
		return
	}

	fmt.Println("sem ganhador")
}
