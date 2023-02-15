package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 0; i <= 2; i++ {
		fmt.Println(i)
	}

	i := "Vova"
	fmt.Println(i)

	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("current value: %d\n", i)

	}
	fmt.Println("After for loop with BREAK")

	for i := 0; i <= 15; i++ {
		if i > 10 && i <= 13 {
			continue
		}
		fmt.Printf("current value: %d\n", i)
	}
	fmt.Println("CONTINUE")

	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
outer:
	for i := 0; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer
			}
		}
	}

	var loop int = 0

	for loop < 10 {
		fmt.Printf("In while like loop %d\n", loop)
		loop++

	}
	var passwd string
	for {
		fmt.Printf("Insert passwd: ")
		fmt.Scan(&passwd)
		if strings.Contains(passwd, "1234") {
			fmt.Println("Weak passwd. Try again")
			continue
		}
		fmt.Println("Passwd accepted")
		break
	}

	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)

	}

}
