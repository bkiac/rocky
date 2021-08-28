package main

import (
	"errors"
	"fmt"
)

const (
	movieTemplate = `
%s
- Reference:: [[Films]]
	- Status::
	- Genres:: %s
	- Description:: %s
	- Directors:: %s
	- Writers:: %s
	- Stars:: %s
	- Cinematographers::
	- Editors::
	- Music::
	- Date:: %s
	- Keywords::
- Rating::
- Highlights::
	- 
- Notes::
	- 
`

	seriesTemplate = `
%s
- Reference:: [[TV Series]]
	- Status:: 
	- Genres:: %s
	- Description:: %s
	- Creators:: %s
	- Stars:: %s
	- Date:: %s
	- Keywords::
- Rating::
- Highlights::
	- 
- Notes::
	- 
`

	bookTemplate = `
%s
- Reference:: [[Books]]
	- Status:: 
	- Authors:: %s
	- Description:: 
		- %s
	- Genres:: %s
	- Date:: %s
	- Rating::
	- Keywords::
- Highlights::
	- 
- Notes::
	- 
`
)

func referencePage(s string) string {
	return fmt.Sprintf("[[%s]]", s)
}

func referenceList(s []string) string {
	var rs []string
	for _, e := range s {
		rs = append(rs, referencePage(e))
	}
	return Join(rs)
}

func referenceAuthor(a Author) string {
	ref := referencePage(a.Name)
	if a.Role == "Writer" {
		return ref
	}
	return fmt.Sprintf("%s (%s)", ref, a.Role)
}

func referenceAuthors(as []Author) string {
	var ras []string
	for _, a := range as {
		ras = append(ras, referenceAuthor(a))
	}
	return Join(ras)
}

func MovieOrSeriesToRoamPage(mos *MovieOrSeries) (string, error) {
	if mos.Type == "movie" {
		return fmt.Sprintf(
			movieTemplate,
			mos.Title,
			referenceList(mos.Genres),
			mos.Description,
			referenceList(mos.Directors),
			referenceList(mos.Writers),
			referenceList(mos.Stars),
			mos.Date,
		), nil
	}
	if mos.Type == "series" {
		return fmt.Sprintf(
			seriesTemplate,
			mos.Title,
			referenceList(mos.Genres),
			mos.Description,
			referenceList(mos.Writers),
			referenceList(mos.Stars),
			mos.Date,
		), nil
	}
	return "", errors.New("type: unknown type")
}

func BookToRoamPage(b *Book, html bool) string {
	var d string
	if html {
		d = b.Description.HTML
	} else {
		d = b.Description.Text
	}
	return fmt.Sprintf(
		bookTemplate,
		b.Title,
		referenceAuthors(b.Authors),
		d,
		referenceList(b.Genres),
		b.PublicationDate.First,
	)
}
