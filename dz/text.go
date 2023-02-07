package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		text string
	)
	fmt.Scan(&text)
	if strings.Contains(text, "чат") == true {
		fmt.Println("БОТ")

	} else {
		fmt.Println("НЕ БОТ")
	}

}
