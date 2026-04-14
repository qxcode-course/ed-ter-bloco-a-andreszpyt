package main

import "fmt"

func main() {
	var qt, qb int
	fmt.Scan(&qt, &qb)
	fig := make([]int, qb)
	mapa := make(map[int]int, 0)
	rep := make([]int, 0)
	faltam := make([]int, 0)

	for i := 0; i < qb; i++ {
		fmt.Scan(&fig[i])
		mapa[fig[i]]++
	}

	for i := 1; i <= qb-1; i++ {
		if mapa[qt] == 0 {
			faltam = append(faltam, fig[i])
		}
	}

	for i := range fig {
		if mapa[fig[i]] > 1 {
			rep = append(rep, fig[i])
			mapa[fig[i]]--
		}
	}

	if len(rep) > 0 {
		for i := range rep {
			if i < len(rep)-1 {
				fmt.Printf("%d ", rep[i])
			} else {
				fmt.Printf("%d\n", rep[i])
			}
		}
	} else {
		fmt.Printf("N\n")
	}

	if len(faltam) > 0 {
		for i := range faltam {
			if i < len(faltam)-1 {
				fmt.Printf("%d ", faltam[i])
			} else {
				fmt.Printf("%d\n", faltam[i])
			}
		}
	} else {
		fmt.Printf("N\n")
	}
}
