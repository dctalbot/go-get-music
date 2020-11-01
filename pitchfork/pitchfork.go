package pitchfork

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dctalbot/go-get-music/model"
	"github.com/gocolly/colly/v2"
)

const bestNewAlbumsURL = "https://pitchfork.com/reviews/best/albums/?page=1"

// Do does the thing
func Do() []model.RowResult {
	var result []model.RowResult

	c := colly.NewCollector()

	c.OnHTML(".review", func(e *colly.HTMLElement) {

		var artists []string
		name := e.DOM.Find("h2").Text()

		e.DOM.Find(".artist-list").Children().Each(func(a int, b *goquery.Selection) {
			artists = append(artists, b.Text())
		})

		result = append(result, model.RowResult{Artist: strings.Join(artists, ", "), Name: name})
	})

	c.Visit(bestNewAlbumsURL)

	return result
}
