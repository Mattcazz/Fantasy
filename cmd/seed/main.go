package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mattcazz/Fantasy.git/db"
	"github.com/Mattcazz/Fantasy.git/service/player"
	"github.com/Mattcazz/Fantasy.git/service/team"
	"github.com/Mattcazz/Fantasy.git/types"
)

func main() {

	APIResponse := getTeamsFromAPI()

	db := db.ConnectDB()

	playerStore := player.NewPlayerStore(db)
	teamStore := team.NewTeamStore(db)

}

func getTeamsFromAPI() *types.APIResponse {

	liga_id := 2014
	season := 2025

	url := fmt.Sprintf("http://api.football-data.org/v4/competitions/%d/teams/?season=%d", liga_id, season)
	return getFromAPIurl(url)
}

func getFromAPIurl(url string) *types.APIResponse {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal("Error getting a new request")
		return nil
	}

	req.Header.Add("X-Auth-Token", "2ef5b7af58674f508a684607709b316f")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("Error getting a the response")
		return nil
	}

	var responseData types.APIResponse

	err = json.NewDecoder(res.Body).Decode(&responseData)

	if err != nil {
		log.Fatal("Error reading the body of the response")
		return nil
	}

	return &responseData
}
