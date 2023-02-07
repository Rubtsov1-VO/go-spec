package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		one   string
		two   string
		three string
	)

	fmt.Scan(&one, &two, &three)
	if strings.Contains(one, "раз") == true && strings.Contains(two, "два") == true && strings.Contains(three, "три") == true {
		fmt.Println("ОК")
	} else if strings.Contains(one, "один") == true {
		fmt.Println("ОК")
	} else if strings.Contains(one, "1") == true && strings.Contains(two, "2") == true && strings.Contains(three, "3") == true {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}

}
