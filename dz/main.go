package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		family string
	)

	fmt.Scan(&name, &age, &family)

	fmt.Printf("Имя: %s , Фамилия: %s, Возраст: %d . Студент BPS", name, family, age)

}
