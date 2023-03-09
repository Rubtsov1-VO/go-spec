package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		suffix := strings.Repeat(" ", n-i)
		rep := strings.Repeat("@", i*2-1)
		fmt.Printf("%s%s\n", suffix, rep)
	}

}
