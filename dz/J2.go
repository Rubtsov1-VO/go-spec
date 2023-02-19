package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var i int
	var sl []int
	var k []int
	fmt.Scan(&n)
	var l = -n
	for l < 0 {
		sl = append(sl, l)
		l++
		t := math.Abs(float64(l))
		fmt.Printf("Квадрат числа %d равен %v\n", l-1, math.Pow(t+1, 2))

	}
	//	fmt.Println(k)
	for i <= n {
		k = append(k, i)
		i++
		fmt.Printf("Квадрат числа %d равен %v\n", i-1, math.Pow(float64(i-1), 2))
	}
	//fmt.Println(sl)

	//fmt.Printf("%d %d", sl, k)
}

//for l < 0 {
//k = append(k, l)
//l++
//fmt.Println(k)
