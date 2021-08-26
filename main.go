package main

<<<<<<< HEAD
import (
	"fmt"

	"github.com/eefret/gomdb"
)

func main() {
	cfg := GetConfig()
	api := gomdb.Init(cfg.omdbAPIKey)
	query := &gomdb.QueryData{Title: "There Will Be Blood", SearchType: gomdb.MovieSearch}
	res, err := api.Search(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Search)
=======
func main() {
	GetBook("https://www.goodreads.com/book/show/2175.Madame_Bovary")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
>>>>>>> 0037e3e (Start goodreads scraper)
}
