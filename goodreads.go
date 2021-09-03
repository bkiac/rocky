package main

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var GoodreadsURLRegexp = regexp.MustCompile(
	`^(https?://)?(w{3}\.)?goodreads\.com/book/show/([0-9]+).(.*)(\?(.*))?$`,
)
var seriesRegexp = regexp.MustCompile(`^\((.*) #.*\)`)
var shelvedByUserRegexp = regexp.MustCompile("^[,0-9]* users?$")
var authorRegexp = regexp.MustCompile(
	`^([^\(\)]*)( \(Goodreads Author\))?( \((.*)\))?$`,
)
var publicationDateRegexp = regexp.MustCompile(
	`^\s*Published\s*?(([a-zA-Z]*)?( [0-9a-z]*)? ?[0-9]+)\s*.*\s*(\(first published (([a-zA-Z]*)?( [0-9a-z]*)? ?[0-9]+)\))?\s*$`,
)

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
	if !GoodreadsURLRegexp.MatchString(url) {
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

	genreSet := NewSet()
	c.OnHTML(".bookPageGenreLink", func(e *colly.HTMLElement) {
		g := strings.TrimSpace(e.Text)
		if !shelvedByUserRegexp.MatchString(g) {
			genreSet.Add(g)
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
		s := publicationDateRegexp.FindStringSubmatch(strings.TrimSpace(e.Text))
		edition := s[1]
		first := s[5]
		publicationDate = PublicationDate{
			edition,
			first,
		}
	})

	var coverImage string
	c.OnHTML("#coverImage", func(e *colly.HTMLElement) {
		coverImage = e.Attr("src")
	})

	c.OnScraped(func(r *colly.Response) {
		genres := genreSet.Values()
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
