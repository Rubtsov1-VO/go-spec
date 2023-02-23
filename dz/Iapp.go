package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var i int
	var m int
	var k int
	var s int
	var str []string
	var arr []string
	fmt.Scan(&n)
	for i < n {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			arr = append(arr, scanner.Text())
			//	fmt.Println(arr)
			break
		}
		i++
		//	fmt.Println(arr[i-1])
	}
	fmt.Scan(&m)
	for k < m {
		fmt.Scan(&s)
		str = append(str, arr[s-1])
		k++
	}
	//fmt.Println(str)
	for _, l := range str {
		fmt.Printf("%s\n", l)
	}
}
