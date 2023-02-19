package main

import (
	"fmt"
)

func main() {
	var value int
	var s []int
	i := 0
	fmt.Scan(&value)
	for value < 10000000 && i < value && value >= 0 {
		i++
		if value > 0 && value%i == 0 {
			s = append(s, i)
			continue
		}

	}
	for f, _ := range s {
		f++
		if len(s) == 2 && f == 2 {
			fmt.Printf("%d %d\nACHTUNG", s[0], s[1])
		} else if len(s) > 2 && f < len(s)+1 {

			fmt.Printf("%d ", s[f-1])
		} else if len(s) == 1 {
			fmt.Println(1)
		}
	}
}
