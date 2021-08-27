package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var goodreadsURLRegexp = regexp.MustCompile(`^https?://(w{3}\.)?goodreads\.com/book/show/([0-9]+).(.*)(\?(.*))?$`)
var shelvedByUserRegexp = regexp.MustCompile("^[,0-9]* users?$")
var authorRole = regexp.MustCompile(`^(\((.*)\))$`)
var editionPublicationDateRegexp = regexp.MustCompile("^(([a-zA-Z]*) ([0-9]*[a-z]*) )?([0-9]*)$")
var firstPublicationDateRegexp = regexp.MustCompile(`^(\(first published )(.*)(\))$`)

type Author struct {
	Name string
	Role string
}

type PublicationDate struct {
	Edition string
	First   string
}

type Book struct {
	Title                string
	Authors              []Author
	Genres               []string
	PublicationDate      PublicationDate
	FirstPublicationDate string
	CoverImage           string
}

func GetBook(url string) (*Book, error) {
	if !goodreadsURLRegexp.MatchString(url) {
		return nil, errors.New("regexp: invalid URL")
	}

	c := colly.NewCollector()
	book := new(Book)

	var title string
	c.OnHTML("[property='og:title']", func(e *colly.HTMLElement) {
		title = e.Attr("content")
		fmt.Println(title)
	})

	var authors []Author
	c.OnHTML(".authorName__container", func(e *colly.HTMLElement) {
		t := strings.TrimRight(strings.TrimSpace(e.Text), ",")
		f := strings.Fields(t)
		var author Author
		role := f[len(f)-1]
		if authorRole.MatchString(role) {
			author.Name = strings.Join(f[:len(f)-1], " ")
			author.Role = authorRole.FindStringSubmatch(role)[2]
		} else {
			author.Name = t
			author.Role = "Writer"
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

	var editionPublicationDate string
	var firstPublicationDate string
	c.OnHTML("#details > .row:nth-child(2)", func(e *colly.HTMLElement) {
		s := strings.Split(strings.TrimSpace(e.Text), "\n")
		for i, e := range s {
			s[i] = strings.TrimSpace(e)
		}
		editionPublicationDate = s[1]
		firstPublicationDate = firstPublicationDateRegexp.FindStringSubmatch(s[len(s)-1])[2]
	})

	var coverImage string
	c.OnHTML("#coverImage", func(e *colly.HTMLElement) {
		coverImage = e.Attr("src")
	})

	c.OnScraped(func(r *colly.Response) {
		book.Title = title
		book.Authors = authors
		book.Genres = genres
		book.PublicationDate = PublicationDate{
			Edition: editionPublicationDate,
			First:   firstPublicationDate,
		}
		book.FirstPublicationDate = firstPublicationDate
		book.CoverImage = coverImage
	})

	c.Visit(url)

	return book, nil
}
