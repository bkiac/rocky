package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var goodreadsURLRegexp = regexp.MustCompile(`^https?://(w{3}\.)?goodreads\.com/book/show/([0-9]+).(.*)(\?(.*))?$`)
var seriesRegexp = regexp.MustCompile(`^\((.*) #.*\)`)
var shelvedByUserRegexp = regexp.MustCompile("^[,0-9]* users?$")
var authorRegexp = regexp.MustCompile(`^([^\(\)]*)( \(Goodreads Author\))?( \((.*)\))?$`)
var editionPublicationDateRegexp = regexp.MustCompile("^(([a-zA-Z]*) ([0-9]*[a-z]*) )?([0-9]*)$")
var firstPublicationDateRegexp = regexp.MustCompile(`^(\(first published )(.*)(\))$`)

type Author struct {
	Name string
	Role string
}

type Description struct {
	Text string
	HTML string
}

type PublicationDate struct {
	Edition string
	First   string
}

type Book struct {
	Title           string
	Series          string
	Authors         []Author
	Genres          []string
	Description     Description
	PublicationDate PublicationDate
	CoverImage      string
}

func GetBook(url string) (*Book, error) {
	if !goodreadsURLRegexp.MatchString(url) {
		return nil, errors.New("regexp: invalid URL")
	}

	c := colly.NewCollector()
	book := new(Book)

	var title string
	c.OnHTML("#bookTitle", func(e *colly.HTMLElement) {
		title = strings.TrimSpace(e.Text)
	})

	var series string
	c.OnHTML("#bookSeries", func(e *colly.HTMLElement) {
		t := strings.TrimSpace(e.Text)
		if t != "" {
			series = seriesRegexp.FindStringSubmatch(t)[1]
		}
		fmt.Println(series)
	})

	var authors []Author
	c.OnHTML(".authorName__container", func(e *colly.HTMLElement) {
		t := strings.TrimRight(strings.TrimSpace(e.Text), ",")
		s := authorRegexp.FindStringSubmatch(t)
		var author Author
		author.Name = s[1]
		role := s[4]
		if role == "" {
			author.Role = "Author"
		} else {
			author.Role = role
		}
		authors = append(authors, author)
	})

	var genres []string
	c.OnHTML(".bookPageGenreLink", func(e *colly.HTMLElement) {
		g := strings.TrimSpace(e.Text)
		if !shelvedByUserRegexp.MatchString(g) {
			genres = append(genres, g)
		}
	})

	var description Description
	c.OnHTML("#description > span:nth-child(2)", func(e *colly.HTMLElement) {
		text := e.Text
		html, _ := e.DOM.Html()
		description = Description{
			text,
			html,
		}
	})

	var publicationDate PublicationDate
	c.OnHTML("#details > .row:nth-child(2)", func(e *colly.HTMLElement) {
		s := strings.Split(strings.TrimSpace(e.Text), "\n")
		for i, e := range s {
			s[i] = strings.TrimSpace(e)
		}
		publicationDate = PublicationDate{
			Edition: s[1],
			First:   firstPublicationDateRegexp.FindStringSubmatch(s[len(s)-1])[2],
		}
	})

	var coverImage string
	c.OnHTML("#coverImage", func(e *colly.HTMLElement) {
		coverImage = e.Attr("src")
	})

	c.OnScraped(func(r *colly.Response) {
		book = &Book{
			title,
			series,
			authors,
			genres,
			description,
			publicationDate,
			coverImage,
		}
	})

	if err := c.Visit(url); err != nil {
		return nil, err
	}

	return book, nil
}
