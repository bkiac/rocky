package main

import (
	"errors"
	"regexp"

	"github.com/eefret/gomdb"
)

var IMDbURLRegexp = regexp.MustCompile(
	`^(https?://)?((w{3}.)|m.)?imdb.com/title/(tt[0-9]*)/?.*$`,
)

type MovieOrSeries struct {
	Type        string
	Title       string
	Genres      []string
	Description string
	Directors   []string
	Writers     []string
	Stars       []string
	Date        string
}

func extractID(url string) (string, error) {
	m := IMDbURLRegexp.FindStringSubmatch(url)
	l := len(m)
	if m == nil || l != 5 {
		return "", errors.New("regexp: invalid URL")
	}
	return m[l-1], nil
}

func GetMovieOrSeries(url string) (*MovieOrSeries, error) {
	api := gomdb.Init(OMDbAPIKey)
	id, err := extractID(url)
	if err != nil {
		return nil, err
	}
	m, err := api.MovieByImdbID(id)
	if err != nil {
		return nil, err
	}
	return &MovieOrSeries{
		Type:        m.Type,
		Title:       m.Title,
		Genres:      Split(m.Genre),
		Description: m.Plot,
		Directors:   Split(m.Director),
		Writers:     Split(m.Writer),
		Stars:       Split(m.Actors),
		Date:        m.Year,
	}, nil
}
