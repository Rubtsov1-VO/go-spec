package main

import "fmt"

func main() {
	for {
		var number int
		fmt.Scan(&number)
		if number != 0 {
			fmt.Println(number)
			continue
		}
		break
	}
}
