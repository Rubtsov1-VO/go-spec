package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a, &b, &c)
	var d = math.Pow(float64(b), 2) - float64(4*a*c)
	if d > 0 && (a != 0 || b != 0 || c != 0) {
		fmt.Println("два корня")
	} else if d == 0 && (a != 0 || b != 0 || c != 0) {
		fmt.Println("один корень")
	} else if d < 0 && a != 0 || b != 0 || c != 0 {
		fmt.Println("корней нет")
	} else {
		fmt.Println("неверные данные")
	}
}
