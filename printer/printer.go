package printer

import (
	"os"

	"github.com/dctalbot/go-get-music/model"
	"github.com/dctalbot/go-get-music/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

// Print will output rows to stdout in a table
func Print(rows []model.RowResult) {
	t := table.NewWriter()

	for i := range rows {
		t.AppendRow([]interface{}{
			i + 1,
			util.Ellipsis(rows[i].Artists, 45),
			util.Ellipsis(rows[i].Name, 45),
			rows[i].Genres})
	}

	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Artist", "Title", "Genre"})
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
