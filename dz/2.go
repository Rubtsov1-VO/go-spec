package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a float64
		b float64
	)

	fmt.Scan(&a, &b)
	c := a + b
	quadrat := math.Pow(c, 2)

	fmt.Println(quadrat)

}
