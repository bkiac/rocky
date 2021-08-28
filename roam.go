package main

import (
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

func MovieToRoamPage(m *Movie) string {
	return fmt.Sprintf(
		movieTemplate,
		m.Title,
		referenceList(m.Genres),
		m.Description,
		referenceList(m.Directors),
		referenceList(m.Writers),
		referenceList(m.Stars),
		m.Date,
	)
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
