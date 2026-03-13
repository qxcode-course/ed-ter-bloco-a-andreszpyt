package main

import "fmt"

func main() {
	var s string
	var n int

	fmt.Scan(&s, &n)

	if n < 12 {
		fmt.Println(s, "eh crianca")
	} else if n < 18 {
		fmt.Println(s, "eh jovem")
	} else if n < 65 {
		fmt.Println(s, "eh adulto")
	} else if n < 1000 {
		fmt.Println(s, "eh idoso")
	} else {
		fmt.Println(s, "eh mumia")
	}
}
