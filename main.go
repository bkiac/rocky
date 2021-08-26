package main

<<<<<<< HEAD
<<<<<<< HEAD
import (
	"fmt"
)

func main() {
	res, err := GetMovie("https://www.imdb.com/title/tt0469494/")
	if err != nil {
		fmt.Println(err)
	}
<<<<<<< HEAD
	fmt.Println(res.Search)
=======
func main() {
	GetBook("https://www.goodreads.com/book/show/2175.Madame_Bovary")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
>>>>>>> 0037e3e (Start goodreads scraper)
=======
import "fmt"

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
<<<<<<< HEAD
	fmt.Println(book)
>>>>>>> 5e7b5f1 (Add return type to book getter)
=======
>>>>>>> 94b4151 (Add movie print)
=======
	fmt.Println(res)
>>>>>>> a7a89b4 (Add movie getter by URL)
}
