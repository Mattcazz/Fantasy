package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	getTeamsFromAPI()
}

func getTeamsFromAPI() {

	liga_id := 2014
	season := 2025

	url := fmt.Sprintf("http://api.football-data.org/v4/competitions/%d/teams?season=%d", liga_id, season)
	getFromAPIurl(url)
}

func getFromAPIurl(url string) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal("Error getting a new request")
		return
	}

	req.Header.Add("X-Auth-Token", "2ef5b7af58674f508a684607709b316f")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("Error getting a the response")
		return
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Error reading the body of the response")
		return
	}

	fmt.Print(string(body))
}
