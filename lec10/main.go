package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "hello"
	fmt.Println(name)

	// строка -это байтовый слайс со своими особенностями
	word := "тестовая строка"
	fmt.Printf("String %s\n", word)

	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - формат представления шестнадцатиричного числа

	}
	fmt.Println()
	// каким образом получать доступ к отдельно стоящим символам?
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) //%c - формат представления шестнадцатиричного числа

	}
	fmt.Println()

	// строки в go хранятся как набор гtf-8 cимволов
	//руна - это стандартный встроеный тип в пo(alias int32 позволяющий хранить единый неделимый UTF символ вне зависимости от того сколько байт он занимает)
	fmt.Printf("Runes: ")
	runeSlice := []rune(word) //преобразования слайса байт к слайсу рун
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()

	//итерирование по строке с использованием рун

	for id, runeVal := range word { //id - это индекс байта, с которого начинается руна runeVal
		fmt.Printf("%c starts at position %d\n", runeVal, id)

	}
	// создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42}
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDec := []byte{100, 101, 102, 103}
	MyDecStr := string(myDec)
	fmt.Println(MyDecStr)

	//создание строки из слайса рун
	//руны как hex
	runeHexSlice := []rune{0x45, 0x46, 0x47}
	myStrFromRune := string(runeHexSlice)
	fmt.Println(myStrFromRune)

	runeLit := []rune{'v', 'a', 's', 'y', 'a'}
	myStrRuneLit := string(runeLit)
	fmt.Println(myStrRuneLit)

	// длина len() - ко-во байт в слайсе
	fmt.Println(utf8.RuneCountInString(myStrRuneLit))

	//cравнение строк ==  и !=
	word1, word2 := "вася", "петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	// cтроитель строк
	firstName := "Alex"
	secondName := "Johnson"
	total := fmt.Sprintf("%s #### %s", firstName, secondName)
	fmt.Println("Full:", total)
	//строки не изменяемые

	// а слайсы изменяемые
	fullName := []rune(total)
	fullName[0] = 'F'
	fullNames := string(fullName)
	fmt.Println(fullNames)

	//import "strings" - много полезного со строками

}
