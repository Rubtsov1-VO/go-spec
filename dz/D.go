package main

import (
	"fmt"
	"strings"
)

func main() {
outer:
	for {

		var string string
		fmt.Scanln(&string)
		if strings.TrimSpace(string) == "" {

			break outer
		}
		fmt.Println(string)
		continue

	}
}
