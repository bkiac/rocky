package main

import (
	"regexp"

	"github.com/eefret/gomdb"
)

var imdbURLRegexp = regexp.MustCompile(`^https?://(w{3}.)?imdb.com/title/(tt[0-9]*)/?.*$`)

func extractID(url string) string {
	m := imdbURLRegexp.FindStringSubmatch(url)
	return m[len(m)-1]
}

func GetMovie(url string) (*gomdb.MovieResult, error) {
	api := gomdb.Init(GetConfig().omdbAPIKey)
	return api.MovieByImdbID(extractID(url))
}
