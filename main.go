package main

import (
	"fmt"

	"github.com/dctalbot/go-get-music/allmusic"
	"github.com/dctalbot/go-get-music/model"
	"github.com/dctalbot/go-get-music/pitchfork"
	"github.com/dctalbot/go-get-music/printer"
)

func main() {
	fmt.Println("go getting music...")

	var results []model.RowResult
	results = append(results, pitchfork.Do()...)
	results = append(results, allmusic.Do()...)

	printer.Print(results)
}
