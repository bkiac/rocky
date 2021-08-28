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
