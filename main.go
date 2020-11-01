package main

import (
	"fmt"
	"os"

	"github.com/dctalbot/go-get-music/allmusic"
	"github.com/dctalbot/go-get-music/model"
	"github.com/dctalbot/go-get-music/pitchfork"
	"github.com/dctalbot/go-get-music/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func main() {

	fmt.Println("running")

	var results []model.RowResult

	results = append(results, pitchfork.Do()...)
	results = append(results, allmusic.Do()...)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Artist", "Title", "Genre"})

	for line := range results {
		t.AppendRow([]interface{}{
			line + 1,
			util.Ellipsis(results[line].Artists, 45),
			util.Ellipsis(results[line].Name, 45),
			results[line].Genres})
	}

	t.SetAllowedRowLength(150)
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:         "Artist",
			Align:        text.AlignLeft,
			AlignFooter:  text.AlignLeft,
			AlignHeader:  text.AlignLeft,
			VAlign:       text.VAlignMiddle,
			VAlignFooter: text.VAlignTop,
			VAlignHeader: text.VAlignBottom,
			WidthMax:     50,
		},
		{
			Name:         "Title",
			Align:        text.AlignLeft,
			AlignFooter:  text.AlignLeft,
			AlignHeader:  text.AlignLeft,
			VAlign:       text.VAlignMiddle,
			VAlignFooter: text.VAlignTop,
			VAlignHeader: text.VAlignBottom,
			WidthMax:     50,
		},
	})
	t.Render()

}
