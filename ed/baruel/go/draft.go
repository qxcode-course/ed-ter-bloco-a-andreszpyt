package main

import "fmt"

func main() {
	var qt, qb int

	fmt.Scan(&qt)
	fmt.Scan(&qb)

	var fig []int
	var rep []int

	norep := 0

	for i := 0; i < qb; i++ {
		var actual int
		fmt.Scan(&actual)

		if i > 0 && actual == fig[len(fig)-1] {
			rep = append(rep, actual)
			continue
		}

		fig = append(fig, actual)
		norep++
	}

	faltando := []int{}
	j := 0

	for i := 1; i <= qt; i++ {
		if j < len(fig) && fig[j] == i {
			j++
		} else {
			faltando = append(faltando, i)
		}
	}

	if len(rep) > 0 {
		for i, val := range rep {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", val)
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("N")
	}

	if len(faltando) > 0 {
		for i, val := range faltando {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", val)
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("N")
	}

}
