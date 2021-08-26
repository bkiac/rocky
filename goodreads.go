package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	rGoodreadsURL = `https?://(w{3}\.)?goodreads\.com/book/show/([0-9]+).(.*)(\?(.*))?`
)

type Author struct {
	Name string
	Role string
}

type Book struct {
	Title           string
	Authors         []Author
	PublicationDate string
	Genres          []string
	CoverImage      string
}

func GetBook(url string) {
	c := colly.NewCollector()

	// var coverImage string
	// c.OnHTML("#coverImage", func(e *colly.HTMLElement) {
	// 	coverImage = e.Attr("src")
	// })

	// re := regexp.MustCompile("^[,0-9]* users?$")
	// var genres []string
	// c.OnHTML(".bookPageGenreLink", func(e *colly.HTMLElement) {
	// 	g := strings.TrimSpace(e.Text)
	// 	if !re.MatchString(g) {
	// 		genres = append(genres, g)
	// 	}
	// })

	// var title string
	// c.OnHTML("[property='og:title']", func(e *colly.HTMLElement) {
	// 	title = e.Attr("content")
	// 	fmt.Println(title)
	// })

	// re := regexp.MustCompile(`^(\((.*)\))$`)
	// var authors []Author
	// c.OnHTML(".authorName__container", func(e *colly.HTMLElement) {
	// 	t := strings.TrimRight(strings.TrimSpace(e.Text), ",")
	// 	f := strings.Fields(t)
	// 	var author Author
	// 	role := f[len(f)-1]
	// 	if re.MatchString(role) {
	// 		author.Name = strings.Join(f[:len(f)-1], " ")
	// 		author.Role = re.FindStringSubmatch(role)[2]
	// 	} else {
	// 		author.Name = t
	// 		author.Role = "Writer"
	// 	}
	// 	authors = append(authors, author)
	// })

	c.OnHTML("#details > .row:nth-child(2)", func(e *colly.HTMLElement) {
		ye := strings.TrimSpace(e.Text)
		fmt.Println(ye)
		yee := strings.Split(ye, "\n")
		fmt.Println(strings.Join(yee, ","))
		var yeee []string
		for _, v := range yee {
			yeee = append(yeee, strings.TrimSpace(v))
		}
		fmt.Println(yeee)
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("coverImage", coverImage)
		// fmt.Println("genres", genres)
		// fmt.Println("authors", authors)
	})

	c.Visit(url)
}
