package main

import "fmt"

func main() {
	book, err := GetBook("https://www.goodreads.com/book/show/2175.Madame_Bovary")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(book)
}
