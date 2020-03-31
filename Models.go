package main

type Search struct {
	Meta struct{
		Status int `json:"status"`
	}	`json:"meta"`
	Response struct{
		Hits []Hit `json:"hits"`
	} 	`json:"response"`
}

type Hit struct {
	Type string		`json:"type"`
	Result Result 	`json:"result"`
}

type Result struct {
	PrimaryArtist Artist 	 `json:"primary_artist"`
	TitleWithFeatured string `json:"title_with_featured"`
	Url string				 `json:"url"`
}

type Artist struct {
	Name string `json:"name"`
}