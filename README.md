#iTunes Go

This is a simple wrapper around the iTunes Search API for Go.

##Usage example

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