package main

import "fmt"

func main() {
	//map - это набор пар ключ значение

	mapper := make(map[string]int)

	fmt.Println("empty map: ", mapper)
	//добавление пар в существующую мапу

	mapper["Alice"] = 24
	mapper["Bob"] = 25
	//инициализация мапы с указанием пар
	newMapper := map[string]int{
		"alice": 1000,
		"bob":   1200,
	}
	fmt.Println(newMapper)

	// что может быть ключом в мапе?
	// важно! мапа не упорядочена в GO
	// ключом в мапе может быть только сравнимый тип (==, !=)
	//нулевое значение для map

	//var mapZero map[string]int // == nil

	//mapZero["Alice"] = 12
	//получение элементов из map
	// получение представленного в мапе

	testPerson := "alice"
	fmt.Println("salary: ", newMapper[testPerson])

	//получение не представленного в мапе
	testPerson = "she"
	fmt.Println("salary: ", newMapper[testPerson]) //при обращение по несуществующему ключу, yовая пара не добавляется
	fmt.Println(newMapper)

	//проверка вхождения ключа
	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}
	fmt.Println("Den:", employee["Den"])

	//при обращении по ключу возвращаются 2 значения

	if value, ok := employee["Jo"]; ok {
		fmt.Println("Jo:", value)
	} else {
		fmt.Println("Jo does not exists in map")
	}

	//перебор элементов мапы

	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}

	//как удалять пары

	// удаление существующей пары
	fmt.Println("Before: ", employee)
	delete(employee, "Den")
	fmt.Println("After: ", employee)
	//удаление несуществующей пары

	delete(employee, "Pack")
	fmt.Println("After second: ", employee)

	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna") //очень дорогая операция в плане ресурсов
	}

	// кол-во пар == длина мапы

	fmt.Println("Pair amount in map:", len(employee))

	// мапа как и слайс ссылочный тип
	words := map[string]string{
		"one": "один",
		"two": "два",
	}

	newWords := words
	newWords["three"] = "три"
	delete(newWords, "one")
	fmt.Println("words map:", words)
	fmt.Println("newWords map:", newWords)

	//сравнение мап
	// сравнеие массивов(можно использовать как ключ в мапе)
	if [3]int{1, 2, 3} == [3]int{3, 4, 5} {
		fmt.Println("equal")
	} else {
		fmt.Println("Not equal")
	}
	// сравнение слайсов(из-за сылочного типа, jvжно сравнить только с nil)
	//	if []int{1,2,3} == []int{1,2,3}{
	//		fmt.Println()
	//	}
	// сравнение мап

	aMap := map[string]int{
		"a": 1,
	}
	//bMap := map[string]int{
	//		"a":1 ,
	//	}
	if aMap == nil {
		fmt.Println("Zero value map")
	}

	//aMap != bMap // нельзя из-за ссылочного типа, только сравнение с тil

	//грустное следствие
	//если мапа/слайс являются составляющими какой-либо структуры - структура несравнима

}
