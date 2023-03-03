package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var m int
	var i int
	var k int
	var l int
	var name []string
	var my []string
	var answer []string
	fmt.Scan(&n)
	fmt.Scan(&m)

	for n >= 1 && i < n {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			name = append(name, scanner.Text())
			break
		}
		i++
		//fmt.Println(name)
	}
	for m >= 1 && k < m {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			my = append(my, scanner.Text())
			break
		}
		k++
		//	fmt.Println(my)
		for l < len(my) {
			if strings.Contains(strings.Join(name, " "), my[l]) == true {
				answer = append(answer, "ЕСТЬ")
			} else {
				answer = append(answer, "НЕТ В НАЛИЧИИ")
			}
			l++

		}

	}
	for _, s := range answer {
		fmt.Println(s)
	}

}
