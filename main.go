package main

import (
	"fmt"
)

func main() {
	res, err := GetMovie("https://www.imdb.com/title/tt0469494/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
