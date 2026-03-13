package main

import "fmt"

type Pair struct {
	x, y int
}

func main() {
	var q int
	var d string
	var gomos []Pair

	fmt.Scan(&q)
	fmt.Scan(&d)

	for i := 0; i < q; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		gomos = append(gomos, Pair{x, y})
	}

	cabeca := gomos[0]

	for i := len(gomos) - 1; i > 0; i-- {
		gomos[i] = gomos[i-1]
	}

	switch d {
	case "L":
		gomos[0].x = gomos[0].x - 1
	case "R":
		gomos[0].x = gomos[0].x + 1
	case "D":
		gomos[0].y = gomos[0].y + 1
	case "U":
		gomos[0].y = gomos[0].y - 1
	}

	if len(gomos) > 1 {
		gomos[1] = cabeca
	}

	for _, v := range gomos {
		fmt.Printf("%d %d\n", v.x, v.y)
	}
}
