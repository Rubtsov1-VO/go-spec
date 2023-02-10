package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var x string

	fmt.Scan(&x)
	if utf8.RuneCountInString(x)*23 < 100 {
		fmt.Println(utf8.RuneCountInString(x)*23, "коп.")
	} else if utf8.RuneCountInString(x)*23 > 100 {
		var y float64
		var z int
		y = float64(utf8.RuneCountInString(x)) * 0.23
		z = (int(utf8.RuneCountInString(x)) * 23) % 100
		fmt.Printf("%d р. %d коп.", int(y), z)
	} else {
		fmt.Println("не верно")
	}

}
