package main

import "fmt"

// явная функция - локально-определенный блок кода имеющий имя
//функцию необходимо определить + функцию необходимо вызвать
// сигнатура функции и их определение
//func functionName(params type) typeReturnValue {
//body
//}

// простейшая функция

func add(a int, b int) int {
	result := a + b
	return result

}

func main() {
	fmt.Println("Hello world")
	//вызов простейшей функции
	res := add(10, 20)
	fmt.Println(res)
	multic := mult(10, 20, 30, 40)
	fmt.Println(multic)
	per, area := rectangle(10.5, 10.5)
	fmt.Println(per, area)

	per2, area2 := nameReturn(10, 20)
	fmt.Println(per2, area2)

	empty(10) //ok
	//	resEmpty := empty(10)//ничего не возращает
	helloVariadic(10, 20, 30, 40)

	sum1 := enterVariadic(10, 20, 30, 40)
	sliceNumber := []int{10, 20, 30}
	sum2 := enterVariadic(sliceNumber...)
	fmt.Println(sum1, sum2)
	//анонимная функция(cинтаксис)
	anon := func(a, b int) int {
		return a + b
	}
	fmt.Println(anon(20, 30))

	fmt.Println(bigFunc(10, 20))
}

//функция с однотипными параметрами

func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

//возврат больше чем одного значения

func rectangle(lenght, width float64) (float64, float64) {
	var perimeter = 2 * (lenght + width)
	var area = lenght * width
	return perimeter, area
}

//именованный возврат значений

func nameReturn(a, b int) (perimeter, area int) {

	perimeter = 2 * (a + b)
	area = a * b
	return //не нужно указывать возвращаемые переменные
}

/// при вызове оператора return функция  прекращает свое выполнение и возвращает результат

func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}
	if a == b {
		return a, true
	}

	return 0, false
}

//что вернется в случае если return в принципе пустой

func empty(a int) {
	fmt.Println("I'm empty")
	return
}

/*
multi
sentence
*/

// *args, **kwargs

//variadic parameters
func helloVariadic(a ...int) {
	fmt.Println(a)

}

//континуальный параметр всегда последний //(a int, b int, c ...int)
// вариадик параметр - на всю функцию один != (a ...int, b ...int) - нельзя

//передача слайса или использование вариадик параметерс

func enterVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

// анонимная функция внутри явной
func bigFunc(aArg, bArg int) int {
	return func(a, b int) int { return a + b }(aArg, bArg)
}
