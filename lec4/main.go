package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//bool default false
	var firstBool bool

	fmt.Println(firstBool)
	aBool, bBool := true, false
	fmt.Println("AND: ", aBool && bBool)
	fmt.Println("OR: ", aBool || bBool)
	fmt.Println(!aBool)

	//Numerics
	a := 92
	fmt.Printf("Type is %T\n", a)
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	var c int32 = 12
	var d int64 = 12
	fmt.Println(c + int32(d))

	//	var third64 int64 = 16123413
	//	var fourint int = 156234
	//	fmt.Println(third64 + fourint)

	//complex

	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)

	//string

	name := "Fedya"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string: ", name, len(name))
	fmt.Println("Amount: ", utf8.RuneCountInString(name))

	total, sub := "ABCD", "BC"
	fmt.Println(strings.Contains(total, sub))

	var sampleRune rune
	fmt.Println(sampleRune)
	fmt.Printf("Rune %c\n", sampleRune)

}
