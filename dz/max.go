package main

import (
	"fmt"
)

func main() {
	var x float64

	for {
		fmt.Scan(&x)
		if x > 10 && x < 20 {
			continue
		} else if x == 0 {
			break
		}
	}
}
