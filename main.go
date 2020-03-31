package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	searchQuery := os.Args[1:][0]
	var song Hit

	res, err := searchRequest(searchQuery)
	checkErr(err)

	hits := res.Response.Hits

	if len(hits) == 1 {

		song = hits[0]

	} else if len(hits) == 0 {

		checkErr(errors.New("song not found("))

	} else {

		stringBuilder := ""
		for i, elem := range hits {
			if elem.Type == "song" {
				stringBuilder += fmt.Sprintf("%d. %s – %s\n", i+1, elem.Result.PrimaryArtist.Name, elem.Result.TitleWithFeatured)
			}
		}
		fmt.Print(stringBuilder + "Choose song(index): ")

		var text string
		_, err = fmt.Scan(&text)
		checkErr(err)

		songIndex, err := strconv.Atoi(text)
		checkErr(err)

		if songIndex > len(hits) || songIndex < 1 {
			checkErr(errors.New("invalid song index"))

		}

		song = hits[songIndex-1]
	}

	text, err := scrapeText(song.Result.Url)
	checkErr(err)

	fmt.Println(text)
	fmt.Println("Lyrics from: " + song.Result.Url)
}

func checkErr(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}

func searchRequest(songName string) (*Search, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.genius.com/search?q="+songName, nil)
	if checkErr(err) {
		return nil, err
	}

	token, ok := token()
	if !ok {
		panic("handle me please")
	}

	req.Header.Add("Authorization", "Bearer - "+token)

	resp, err := client.Do(req)
	if checkErr(err) {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if checkErr(err) {
		return nil, err
	}

	var search Search
	err = json.Unmarshal(body, &search)
	if checkErr(err) {
		return nil, err
	}

	return &search, nil
}

func scrapperErr(err error) bool {
	if err != nil {
		log.Print("ScrapperError: " + err.Error())
		return true
	}
	return false
}

func scrapeText(url string) (string, error) {
	resp, err := http.Get(url)
	if scrapperErr(err) {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if scrapperErr(err) {
		return "", err
	}

	result := doc.Find(".lyrics").Text()
	return result, nil
}
