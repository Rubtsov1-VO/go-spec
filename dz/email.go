package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		login string
		email string
	)
	fmt.Scan(&login, &email)
	if strings.Contains(login, "@") == true || len(login) < 10 {
		fmt.Println("Некорректный логин")

	} else if strings.Contains(email, "@") == false || strings.Contains(email, ".") == false {
		fmt.Println("Некорректная почта")
	} else {
		fmt.Println("ОК")
	}

}
