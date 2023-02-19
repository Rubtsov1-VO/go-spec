package main

import (
	"fmt"
)

func main() {
	var n int

	var s []int

	for {
		fmt.Scan(&n)
		if n > 0 {
			s = append(s, n)
			continue
		} else {
			b := 0
			for f, i := range s {
				f++
				fmt.Println(i)
				fmt.Println(f)
				fmt.Println(s)
				for f%2 == 0 {
					b -= i
					break
				}
				for f%2 != 0 {
					b += i
					break
				}

			}
			fmt.Println(b)
			break
		}
	}
}
