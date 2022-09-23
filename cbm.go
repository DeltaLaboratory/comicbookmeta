package cbm

import "time"

type CBM struct {
	// The title of the comic book
	Title string `json:"title"`
	// The identifiers of the comic book
	Identifiers []Identifier `json:"identifiers"`
	// The authors of the comic book
	Authors []Author `json:"authors,omitempty"`
	// "The series data of the comic book"
	Series *Series `json:"series,omitempty"`
	// The tags of the comic book
	Tags []string `json:"tags,omitempty"`
	// The language of the comic book in ISO 639-1 format
	Language *string `json:"language,omitempty"`
	// The date of publication of the comic book
	PublishedDate *time.Time `json:"published_date,omitempty"`
	// The date of distribution of the comic book
	DistributedDate *time.Time `json:"distributed_date,omitempty"`
	// The read direction of the comic book
	ReadDirection string `json:"read_direction"`
}

type Identifier struct {
	// The type of the identifier
	Type string `json:"type"`
	// The value of the identifier
	Value string `json:"value"`
}

type Author struct {
	// The identifier of the author
	Identifier string `json:"identifier,omitempty"`
	// The name of the author
	Name string `json:"name"`
	// The role of the author
	Role string `json:"role,omitempty"`
}

type Series struct {
	// The identifier of the series
	Identifier string `json:"identifier,omitempty"`
	// The title of the series
	Title string `json:"name"`
	// The number of the comic book in the series
	Index int `json:"index"`
	// The parent series data of the comic book
	Parent *SeriesParent `json:"parent,omitempty"`
}

type SeriesParent struct {
	// The identifier of the parent series
	Identifier string `json:"identifier"`
	// The title of the parent series
	Title string `json:"name"`
}
