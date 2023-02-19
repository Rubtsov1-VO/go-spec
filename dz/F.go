package main

import (
	"fmt"
	"strconv"
)

func main() {

	var int int
	fmt.Scan(&int)
	if int > 0 {
		fmt.Printf(strconv.FormatInt(int64(int), 2))
	} else {
		fmt.Println("неверное число")
	}

}
