package main

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)

	if w >= 1 && w <= 100 && w%2 == 0 {
		a := w / 2

		if a%2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	} else {
		fmt.Println("NO")
	}
}
