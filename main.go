package main

import (
	"fmt"
)

func main() {
	book, err := GetBook("https://www.goodreads.com/book/show/2175.Madame_Bovary")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book")
		fmt.Println(book)
	}

	movie, err := GetMovie("https://www.imdb.com/title/tt0469494/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Movie")
		fmt.Println(movie)
	}
}
