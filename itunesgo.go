package itunesgo

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type SearchRequest struct {
	Term string
	Store string
	Media string
	Entity string
	Attribute string
	Limit int
	Lang string
}

type SearchResult struct {
	WrapperType string
	Explicitness string
	Kind string
	ArtistId int
	CollectionId int
	TrackId int
	ArtistName string
	CollectionName string
	TrackName string
	CollectionCensoredName string
	TrackCensoredName string
	ArtistViewUrl string
	ArtworkUrl500 string
	ArtworkUrl100 string
	ArtworkUrl60 string
	ArtworkUrl30 string
	ViewUrl string
	PreviewUrl string
	TrackTimeMillis float64
	DiscCount int
	DiscNumber int
	TrackCount int
	TrackNumber int
	Currency string
	PrimaryGenreName string
	PrimaryGenreId int
	ContentAdvisoryRating string
	ShortDescription string
	LongDescription string
	Features[] string
	SupportedDevices[] string
	IsGameCenterEnabled bool
	ScreenshotUrls[] string
	IPadScreenshotUrls[] string
	Genres[] string
	GenreIds[] string
	ReleaseDate string
	SellerName string
	BundleId string
	LanguageCodesISO2A[] string
	AverageUserRatingForCurrentVersion float64
	UserRatingCountForCurrentVersion int
	TrackContentRating string
	AverageUserRating float64
	UserRatingCount int
	Version string
	Description string
}

type SearchResults struct {
	ResultCount int
	Results[] SearchResult
}

func Search(request SearchRequest) (SearchResults, error) {
	url, urlErr := SearchUrl(request)
	if urlErr == nil {
		result, err := http.Get(url)
		if err == nil {
			defer result.Body.Close()
			contents, _ := ioutil.ReadAll(result.Body)
			var searchResults SearchResults
			jsonErr := json.Unmarshal(contents, &searchResults)
			return searchResults, jsonErr
		} else {
			return SearchResults{}, err
		}
	} else {
		return SearchResults{}, urlErr
	}
}

func SearchUrl(request SearchRequest) (string, error) {
	u, err := url.Parse("https://itunes.apple.com/search")
	if err != nil {
		return "", err
	}
	q := u.Query()
	if request.Term != "" {
		q.Set("term", request.Term)
	}
	if request.Store != "" {
		q.Set("country", request.Store)
	}
	if request.Media != "" {
		q.Set("media", request.Media)
	}
	if request.Entity != "" {
		q.Set("entity", request.Entity)
	}
	if request.Attribute != "" {
		q.Set("attribute", request.Attribute)
	}
	if request.Limit > 0 {
		q.Set("limit", string(request.Limit))
	}
	if request.Lang != "" {
		q.Set("lang", request.Lang)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}