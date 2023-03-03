package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var name string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		name = input.Text()
	} //команда захвата потока ввода и сохранение в буфер(захват идет до символа окончания строки)

	fmt.Println(name)
	fmt.Println("for loop started:")
	for {
		if input.Scan() {
			result := input.Text()

			if result == "" {
				break
			}
			fmt.Println("input string is:", result)
		}

	}

	//преобразование строкового литерала к чему-нибудь числовому

	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Println(numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Println(numInt64)

	numfloat32, _ := strconv.ParseFloat(numStr, 32)
	fmt.Println(numfloat32)
}
