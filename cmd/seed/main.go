package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Mattcazz/Fantasy.git/db"
	"github.com/Mattcazz/Fantasy.git/service/player"
	"github.com/Mattcazz/Fantasy.git/service/team"
	"github.com/Mattcazz/Fantasy.git/types"
)

var database *sql.DB

func main() {

	ScrapeResponse := getResponseFromScraper()

	database = db.ConnectDB()

	playerStore := player.NewPlayerStore(database)
	teamStore := team.NewTeamStore(database)

	err := seedDB(&playerStore, &teamStore, ScrapeResponse)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func getResponseFromScraper() *types.ScrapeResponse {
	file, err := os.ReadFile("team.json")

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	var res types.ScrapeResponse

	var root types.Root
	err = json.Unmarshal(file, &root)

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	for _, s := range root.Props.PageProps.Data.Clasificacion.Standings[0] {
		res.Teams = append(res.Teams, s.Team)
	}

	file, err = os.ReadFile("player.json")

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	var player_response types.ScrapePlayerResponse

	err = json.Unmarshal(file, &player_response)

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	res.Players = &player_response.Players

	return &res
}

func seedDB(ps types.PlayerStore, ts types.TeamStore, response *types.ScrapeResponse) error {

	tx, err := database.Begin()

	if err != nil {
		log.Fatal("Error opening the tx")
	}

	defer tx.Rollback()

	for _, team := range response.Teams {

		t := &types.Team{
			Name:     team.Name,
			Logo_url: team.Logo,
			Web_Id:   team.ID,
		}

		if err := ts.InsertTeamTx(tx, t); err != nil {
			return fmt.Errorf("Tx failed: %s", err.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Tx team commit failed: %s", err.Error())
	}

	tx, err = database.Begin()

	if err != nil {
		log.Fatal("Error opening the tx")
	}

	for _, player := range *response.Players {

		t, err := ts.GetTeamByName(player.Team)

		if err != nil {
			return fmt.Errorf("team by name errror: %s", err.Error())
		}

		p := &types.Player{
			Name:    player.Name,
			Team_id: t.Id,
			Points:  player.Points,
			Value:   player.Value,
			Goals:   player.Goals,
			Assists: player.Assists,
			WebID:   player.WebID,
		}

		if err := ps.InsertPlayerTx(tx, p); err != nil {
			return fmt.Errorf("Tx failed: %s", err.Error())
		}

	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Tx player commit failed: %s", err.Error())
	}

	return nil
}
