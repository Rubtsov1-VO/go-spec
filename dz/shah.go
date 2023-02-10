package main

import (
	"fmt"
)

func main() {
	var (
		one   int
		two   int
		three int
		four  int
	)

	fmt.Scan(&one, &two, &three, &four)
	if one <= 6 && two <= 7 && three == one+2 && four == two+1 {
		fmt.Println("ДА")
	} else if one <= 6 && two <= 8 && three == one+2 && four == two-1 {
		fmt.Println("ДА")
	} else if one <= 7 && two <= 6 && three == one+1 && four == two+2 {
		fmt.Println("ДА")
	} else if one <= 8 && two <= 6 && three == one-1 && four == two+2 {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}

}
