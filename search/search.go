package search

import "streaming/catalog"

func FilterByGenre(contents []catalog.Content, genre string) []catalog.Content {
	var result []catalog.Content
	for _, c := range contents {
		if c.Genre == genre {
			result = append(result, c)
		}
	}
	return result
}
