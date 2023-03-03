package main

import "fmt"

func main() {
	//массивы
	var arr [5]int //при инициализации важно сколько будет элементов
	fmt.Println(arr)
	//определение элементов
	arr[0] = 10
	arr[3] = 39
	fmt.Println(arr)

	newArr := [5]int{10, 20, 30, 50, 39}
	fmt.Println(newArr)

	//создание массива через переменные
	arrNews := [...]int{10, 20, 30, 40}
	fmt.Println(arrNews)

	first := [...]int{1, 2, 3}
	second := first

	fmt.Println(second)

	//массив и его размер - это две составляющие одного типа

	flotArr := [...]float64{12.5, 13, 5, 13.2}
	for i := 0; i < len(flotArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, flotArr[i])
	}

	//8.итеррирование по массиву с range
	var sum float64
	for id, val := range flotArr {
		fmt.Println(id, val)
		sum += val
	}
	fmt.Println(sum)

	// многомерные массивы

	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor", "Jo"},
	}

	fmt.Println(words)

	//иттерирование по многомерному массиву

	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Println(val2)
		}

	}
	//12. общение со срезом
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	slice = append(slice, 200)
	for i, v := range slice {
		fmt.Println(i, v)
	}
	
}
