package main

import (
	"fmt"
)

func main() {
	var (
		a float32
		b float32
	)

	fmt.Scan(&a, &b)
	var c int = int(a + b)
	if c%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else if c%2 != 0 {
		fmt.Println("НЕЧЁТНОЕ")
	} else {
		fmt.Println("не пойми что")
	}

}
