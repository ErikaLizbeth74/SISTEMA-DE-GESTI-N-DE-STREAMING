package catalog

import "fmt"

type Content struct {
	Title  string
	Genre  string
	Year   int
	Length int
}

func NewContent(title, genre string, year, length int) Content {
	return Content{Title: title, Genre: genre, Year: year, Length: length}
}

func NewContentSafe(title, genre string, year, length int) (Content, error) {
	if title == "" || genre == "" || year <= 0 || length <= 0 {
		return Content{}, fmt.Errorf("datos inválidos")
	}
	return Content{Title: title, Genre: genre, Year: year, Length: length}, nil
}

func PrintContent(c Content) {
	fmt.Printf("- %s (%d) [%s] - %d min\n", c.Title, c.Year, c.Genre, c.Length)
}

func (c *Content) SetTitle(title string) {
	c.Title = title
}

func (c Content) GetTitle() string {
	return c.Title
}

func (c Content) Show() {
	PrintContent(c)
}

type Printable interface {
	Show()
}
