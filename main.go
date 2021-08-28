package main

import (
	"fmt"
)

func main() {
	book, err := GetBook("https://www.goodreads.com/book/show/2175.Madame_Bovary")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(BookToRoamPage(book, true))
	}

	fmt.Println()

	mos, err := GetMovieOrSeries("https://www.imdb.com/title/tt6412452/?ref_=nv_sr_srsg_0")
	if err != nil {
		fmt.Println(err)
	} else {
		r, _ := MovieOrSeriesToRoamPage(mos)
		fmt.Println(r)
	}

	mos, err = GetMovieOrSeries("https://www.imdb.com/title/tt0141842/?ref_=fn_al_tt_1")
	if err != nil {
		fmt.Println(err)
	} else {
		r, _ := MovieOrSeriesToRoamPage(mos)
		fmt.Println(r)
	}
}
