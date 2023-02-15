package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	for {

		var one string
		var two string
		fmt.Scanln(&one)
		fmt.Scanln(&two)
		if utf8.RuneCountInString(one) < 8 {
			fmt.Println("Слишком короткий пароль!")
			continue
		} else if utf8.RuneCountInString(one) >= 8 && strings.Contains(one, "123") == true || strings.Contains(one, "qwe") == true {
			fmt.Println("Слишком простой пароль!")
			continue
		} else if utf8.RuneCountInString(one) >= 8 && (strings.Contains(one, "123") || strings.Contains(one, "qwe")) == false && one != two {
			fmt.Println("Введенные пароли различаются!")
			continue
		} else if utf8.RuneCountInString(one) >= 8 && strings.Contains(one, "123") != true || strings.Contains(one, "qwe") != true && one == two {
			fmt.Println("Ну наконец-то!")
			break
		}
		break
	}
}
