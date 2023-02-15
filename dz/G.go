package main

import (
	"fmt"
	"sort"
)

func main() {
	var count int
	var s []float64
	for {
		var candidate float64
		fmt.Scan(&candidate)
		if candidate >= 100.0 && candidate <= 140.0 {
			count++
			s = append(s, candidate)
			continue
		} else if candidate < 0 {
			fmt.Println(count)

			sort.Float64s(s)
			fmt.Printf("%.1f %.1f", s[0], s[count-1])
			break
		}
	}
}
