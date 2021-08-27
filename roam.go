package main

import (
	"fmt"

	"github.com/eefret/gomdb"
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

func MovieToRoamPage(m *gomdb.MovieResult) string {
	return fmt.Sprintf(movieTemplate, m.Title, m.Genre, m.Plot, m.Director, m.Writer, m.Actors, m.Released)
}
