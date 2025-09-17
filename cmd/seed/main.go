package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mattcazz/Fantasy.git/db"
	"github.com/Mattcazz/Fantasy.git/service/player"
	"github.com/Mattcazz/Fantasy.git/service/team"
	"github.com/Mattcazz/Fantasy.git/types"
)

var database *sql.DB

func main() {

	APIResponse := getTeamsFromAPI()

	database = db.ConnectDB()

	playerStore := player.NewPlayerStore(database)
	teamStore := team.NewTeamStore(database)

	err := seedDB(&playerStore, &teamStore, APIResponse)

	if err != nil {
		log.Fatal(err.Error())
	}
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

func seedDB(ps types.PlayerStore, ts types.TeamStore, response *types.APIResponse) error {

	tx, err := database.Begin()

	if err != nil {
		log.Fatal("Error opening the tx")
	}

	defer tx.Rollback()

	for _, team := range response.Teams {

		t := &types.Team{
			Name:     team.Name,
			Logo_url: team.Crest_URL,
		}

		if err := ts.InsertTeamTx(tx, t); err != nil {
			return fmt.Errorf("Tx failed: %s", err.Error())
		}

		for _, player := range team.Squad {

			p := &types.Player{
				Name:        player.Name,
				Nationality: player.Nationality,
				Position:    player.Position,
			}

			if err := ps.InsertPlayerTx(tx, p); err != nil {
				return fmt.Errorf("Tx failed: %s", err.Error())
			}

			if err := ts.AddPlayerToTeamTx(tx, p.Id, t.Id); err != nil {
				return fmt.Errorf("Tx failed: %s", err.Error())
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Tx commit failed: %s", err.Error())
	}

	return nil
}
