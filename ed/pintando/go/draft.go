package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float32
	fmt.Scan(&a, &b, &c)

	p := (a + b + c) / 2
	res := float32(math.Sqrt(float64(p * (p - a) * (p - b) * (p - c))))
	fmt.Printf("%.2f\n", res)
}
