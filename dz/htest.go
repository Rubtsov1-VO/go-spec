package main

import (
	"fmt"
)

func main() {
	var dro int
	var z []int
	var up []int
	fmt.Scan(&dro)
	for i := 0; i < dro; i++ {
		var x int
		var y int

		fmt.Scan(&x)
		fmt.Scan(&y)
		z = append(z, x)
		up = append(up, y)
	}
	fmt.Println(z)
	fmt.Println(up)

}
