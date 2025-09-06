package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	getTeamsFromAPI()
}

func getTeamsFromAPI() {

	liga_id := 140
	season := 2023

	url := fmt.Sprintf("https://v3.football.api-sports.io/teams?league=%d&season=%d", liga_id, season)

	getFromAPIurl(url)
}

func getFromAPIurl(url string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("x-rapidapi-key", "7129b5848d724ed4b4a489523a27c0e7")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
