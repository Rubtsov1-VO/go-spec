package main

import "fmt"

func main() {
	//cлайсы

	startArr := [4]int{10, 20, 30, 40}
	var NewSlice []int = startArr[0:2]
	fmt.Println("Slice[0:2]:", NewSlice)

	//слайс без массива
	secondSlice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)
	//изменение элемента среза

	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4]
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)

	//один массив и два производных среза
	fs := originArr[:]
	ss := originArr[2:4]
	fmt.Println(originArr, fs, ss)
	fs[3]++
	ss[1]++
	fmt.Println(originArr, fs, ss)

	//срез как встроенный тип
	//type slice struct {
	//length int
	//??? int
	//	zeroElement *byte
	//	}
	// длина и емкость слайса
	words := []string{"one", "two", "three", "four"}
	fmt.Println("slice:", words, "Length:", len(words), "Capacity:", cap(words))

	//как создавать слайсы наиболее эффективно
	//make() -позволяет наиболее детально создавать срезы
	s1 := make([]int, 10, 15)
	fmt.Println(s1)
	//добавление элементов в срезы
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords: ", myWords)
	myWords = append(myWords, "four")
	fmt.Println("myWords: ", myWords)
	another := []string{"five", "seven", "six"}
	myWords = append(myWords, another...)
	fmt.Println("myWords", myWords)

	//многомерный срез
	mSl := [][]int{
		{1, 2},
		{3, 4, 5, 6},
	}
	fmt.Println(mSl)

}
