package main

import (
	"errors"
	"regexp"

	"github.com/eefret/gomdb"
)

var imdbURLRegexp = regexp.MustCompile(`^https?://(w{3}.)?imdb.com/title/(tt[0-9]*)/?.*$`)

func extractID(url string) (string, error) {
	m := imdbURLRegexp.FindStringSubmatch(url)
	l := len(m)
	if l != 3 {
		return "", errors.New("regexp: invalid URL")
	}
	return m[l-1], nil
}

func GetMovie(url string) (*gomdb.MovieResult, error) {
	api := gomdb.Init(GetConfig().omdbAPIKey)
	id, err := extractID(url)
	if err != nil {
		return nil, err
	}
	return api.MovieByImdbID(id)
}
