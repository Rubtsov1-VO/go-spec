package main

import "fmt"

func main() {
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {

		fmt.Println("The number", value, "is odd")
	}

	//инициализация в блоке условного оператора

}
