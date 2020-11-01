package pitchfork

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dctalbot/go-get-music/model"
	"github.com/gocolly/colly/v2"
)

const bestNewAlbumsURL = "https://pitchfork.com/reviews/best/albums/"
const pageCount = 1

// Do does the thing
func Do() []model.RowResult {
	var result []model.RowResult

	c := colly.NewCollector()

	c.OnHTML(".review", func(e *colly.HTMLElement) {

		var artists []string
		name := e.DOM.Find("h2").Text()
		var genres []string

		e.DOM.Find(".artist-list").Children().Each(func(a int, b *goquery.Selection) {
			artists = append(artists, b.Text())
		})

		e.DOM.Find(".genre-list").Children().Each(func(a int, b *goquery.Selection) {
			genres = append(genres, b.Text())
		})

		result = append(result, model.RowResult{
			Artists: strings.Join(artists, ", "),
			Name:    name,
			Source:  "Pitchfork",
			Genres:  strings.Join(genres, ", "),
		})
	})

	for i := 1; i <= pageCount; i++ {
		c.Visit(bestNewAlbumsURL + "?page=" + strconv.Itoa(i))
	}

	return result
}
