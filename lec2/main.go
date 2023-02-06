package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello line")
	fmt.Println("print 2 line")
	fmt.Print("first")
	fmt.Print("second")
	fmt.Printf("Hello, my name is %s\n", "Vova")

	////декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("are after:", age)

	var height int = 183
	fmt.Println("my height is:", height)

	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	var firstVar, secondVar = 1, 3
	fmt.Printf("first: %d, second: %d", firstVar, secondVar)

	var (
		name      string = "Vova"
		personAge int    = 27
		personUid int
	)

	fmt.Printf("Name: %s\nAge %d\nUID: %d\n", name, personAge, personUid)

	var a, b = 30, "Vova"
	fmt.Println(a, b)

	count := 10
	fmt.Println(count)

	arg, barg := 10, "Vova"

	fmt.Println(arg, barg)

	width, length := 20.5, 30.4
	fmt.Printf("Min dimensional of rectangle is:%.3f\n ", math.Min(width, length))

}
