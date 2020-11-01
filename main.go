package main

import (
	"fmt"

	"github.com/dctalbot/go-get-music/model"
	"github.com/dctalbot/go-get-music/pitchfork"
)

func main() {

	fmt.Println("running")

	var result []model.RowResult

	result = append(result, pitchfork.Do()...)
	// result = append(result, allmusic.do())
	// result = append(result, npr.do())

	fmt.Println(result)

}
