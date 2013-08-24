package main

import (
	"fmt"
	"github.com/programmingthomas/itunesgo"
)

func main() {
	fmt.Println("Welcome to iTunes Go")
	request := itunesgo.SearchRequest{Term:"Godfather", Store:"US", Media:"movie"}
	results, err := itunesgo.Search(request)
	if err == nil {
		fmt.Println("There are", results.ResultCount, "results")
		for _, result := range results.Results {
			fmt.Println("Found", result.TrackName, "type", result.Kind, "id", result.TrackId)
		}
	} else {
		fmt.Println(err)
	}
}