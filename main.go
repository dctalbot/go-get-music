package main

import (
	"fmt"
	"os"

	"github.com/dctalbot/go-get-music/model"
	"github.com/dctalbot/go-get-music/pitchfork"
	"github.com/jedib0t/go-pretty/table"
)

func main() {

	fmt.Println("running")

	var results []model.RowResult

	results = append(results, pitchfork.Do()...)
	// result = append(result, allmusic.do())
	// result = append(result, npr.do())

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Artist", "Name", "Genre"})

	for line := range results {
		t.AppendRow([]interface{}{line + 1, results[line].Artists, results[line].Name, results[line].Genres})
	}

	t.Render()

}
