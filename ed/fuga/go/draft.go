package main

import "fmt"

func main() {
	var h, p, f, d int
	fmt.Scan(&h, &p, &f, &d)

	for {
		if f == p {
			fmt.Println("N")
			return
		}

		if f == h {
			fmt.Println("S")
			return
		}

		f += d

		if f > 15 {
			f = 0
		}
		if f < 0 {
			f = 15
		}
	}
}
