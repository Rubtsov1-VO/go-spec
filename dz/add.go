package main

import (
	"fmt"
)

func main() {
	var n int
	var i int
	var one int
	var two int
	var number int
	var arr []int
	fmt.Scan(&n)
	for i < n {
		fmt.Scan(&number)
		arr = append(arr, number)
		i++
	}
	fmt.Println(arr)
	fmt.Scan(&one)
	fmt.Scan(&two)
	c := 1
	var sum int
	for one >= c || c <= two {
		c++
		sum += arr[c-2]

	}
	fmt.Println(sum)
}
