package main

import (
	"fmt"
)

func main() {
	var n int
	var amount int
	var s []int
	var i int
	fmt.Scan(&amount)
	for i < amount && amount > 0 {
		i++
		fmt.Scan(&n)
		if n > 0 {
			s = append(s, n)
			continue

		}
	}
	b := 0

	for f, a := range s {
		f++
		for f%2 == 0 {
			b -= a
			break
		}
		for f%2 != 0 {
			b += a
			break
		}

	}
	fmt.Println(b)

}
