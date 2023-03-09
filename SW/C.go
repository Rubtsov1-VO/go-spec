package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	var s string
	var i int
	fmt.Scan(&n,
		&k,
		&s)

	for i = 1; i <= n; i++ {

		if i == 1 || i == k {
			m := strings.Repeat(s, n)
			fmt.Println(m)
			continue

		}
		fmt.Printf("%s%s%s\n", s, strings.Repeat(" ", k-2), s)
	}

}
