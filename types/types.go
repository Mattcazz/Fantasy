package types

import "database/sql"

type TeamStore interface {
	GetTeamByName(string) (*Team, error)
	InsertTeamTx(*sql.Tx, *Team) error
	DeleteTeam(*Team) error
	AddPlayerToTeamTx(*sql.Tx, int, int) error
}

type PlayerStore interface {
	InsertPlayerTx(*sql.Tx, *Player) error
	DeletePlayer(*Player) error
	PlayerDailyUpdate(*Player) error
}
type Team struct {
	Id       int
	Name     string
	Logo_url string
}

type Player struct {
	Id                  int
	Name                string
	Nationality         string
	Position            string
	Points              int
	Price               float32
	Fluctuation         float32
	Fluctuation_History []float32
	Img_url             string
}

type APIResponse struct {
	Teams []APIResponseTeam `json:"teams"`
}

type APIResponsePlayer struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	Nationality string `json:"nationality"`
}

type APIResponseTeam struct {
	Name      string              `json:"name"`
	Crest_URL string              `json:"crest"`
	Squad     []APIResponsePlayer `json:"squad"`
}
