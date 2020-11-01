package allmusic

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dctalbot/go-get-music/model"
	"github.com/gocolly/colly"
)

const editorsChoiceURL = "https://www.allmusic.com/newreleases/editorschoice"

// Do does the thing
func Do() []model.RowResult {
	var result []model.RowResult

	c := colly.NewCollector()

	c.OnHTML(".editors-choice-item", func(e *colly.HTMLElement) {

		var artists []string
		title := e.DOM.Find(".title > a").Text()
		var genres []string

		e.DOM.Find(".artist").Children().Each(func(a int, b *goquery.Selection) {
			artists = append(artists, b.Text())
		})

		e.DOM.Find(".genres").Children().Each(func(a int, b *goquery.Selection) {
			genres = append(genres, b.Text())
		})

		result = append(result, model.RowResult{
			Artists: strings.Join(artists, ", "),
			Name:    title,
			Source:  "All Music",
			Genres:  strings.Join(genres, ", "),
		})
	})

	c.Visit(editorsChoiceURL)

	return result
}
