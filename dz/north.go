package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	var result string

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		name = input.Text()
	}
	runeSlice := []rune(name)
	for run, pum := range runeSlice {
		if run%2 == 0 && pum != ' ' {
			result += strings.Repeat(string(pum), 3)
		}

	}
	fmt.Println(result)
}
