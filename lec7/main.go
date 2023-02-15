package main

import "fmt"

func main() {
	//switch
	var price int
	fmt.Scan(&price)

	switch price {
	case 100:
		fmt.Println("first")
	case 110:
		fmt.Println("second")
	case 120:
		fmt.Println("third")
	default:
		fmt.Println("Error")
	}

	var ageGroup string = "b"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("Age 10-40")
	case "d", "e":
		fmt.Println("age 50-70")
	default:
		fmt.Println("you are young")

	}

	//case с выражениями

	var ages int
	fmt.Scan(&ages)

	switch {
	case ages <= 18:
		fmt.Println("young")
	case ages > 18 && ages <= 30:
		fmt.Println("norm")
		fallthrough
	case ages > 19 && ages <= 22:
		fmt.Println("norm2")
	case ages > 30:
		fmt.Println("too old")
	}

}
