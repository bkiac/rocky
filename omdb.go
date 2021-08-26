package main

import (
	"regexp"

	"github.com/eefret/gomdb"
)

const (
	rImdbURL = "https?://(w{3}.)?imdb.com/title/(tt[0-9]*)/?"
)

func extractID(url string) string {
	re := regexp.MustCompile(rImdbURL)
	m := re.FindStringSubmatch(url)
	return m[len(m)-1]
}

func GetMovie(url string) (*gomdb.MovieResult, error) {
	api := gomdb.Init(GetConfig().omdbAPIKey)
	return api.MovieByImdbID(extractID(url))
}
