package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	fmt.Scan(&a, &b)

	perimetr := (2 * a) + (2 * b)
	ploshad := a * b

	fmt.Printf("Периметр: %d\nПлощадь: %d", perimetr, ploshad)

}
