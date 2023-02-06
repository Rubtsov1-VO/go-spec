package main

import "fmt"

func main() {
	var (
		word1  string
		word2  string
		word3  string
		word4  string
		author string
	)

	fmt.Scan(&word1, &word2, &word3, &word4, &author)

	fmt.Printf("%s - %s\n%s - %s\n%s - %s\n%s - %s\n", word4, author, word3, author, word2, author, word1, author)

}
