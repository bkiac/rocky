package main

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
}
