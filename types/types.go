package types

import (
	"database/sql"
	"time"
)

type TeamStore interface {
	GetTeamByName(string) (*Team, error)
	InsertTeamTx(*sql.Tx, *Team) error
	DeleteTeam(*Team) error
}

type PlayerStore interface {
	InsertPlayerTx(*sql.Tx, *Player) error
	DeletePlayer(*Player) error
	PlayerDailyUpdate(*Player) error
}

type FluctuationStore interface {
	InsertFluctuation(*Fluctuation) error
	GetFluctuationHistoryFromPlayer(int) ([]Fluctuation, error)
	GetLastFluctuationFromPlayer(int) (*Fluctuation, error)
}

type Team struct {
	Id       int
	Name     string
	Logo_url string
	Web_Id   int
}

type Player struct {
	Id      int
	Team_id int
	Name    string
	Points  int
	Value   float32
	Avg     float32
	Goals   int
	Assists int
	Img_url string
	WebID   int
}

type Fluctuation struct {
	Player_Id  int
	Value      float32
	Created_at time.Time
}

type ScrapeResponse struct {
	Teams   []ScrapeTeamResponse
	Players *[]ScrapePlayer
}

type ScrapePlayerResponse struct {
	Players []ScrapePlayer `json:"players"`
}

type ScrapePlayer struct {
	Name    string  `json:"nn"`
	Team    string  `json:"tn"`
	Points  int     `json:"tp"`
	WebID   int     `json:"id"`
	Value   float32 `json:"mv"`
	Avg     float32 `json:"avg"`
	Goals   int     `json:"g"`
	Assists int     `json:"ga"`
}

type ScrapeTeamResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Root struct {
	Props struct {
		PageProps struct {
			Data struct {
				Clasificacion struct {
					Standings [][]struct {
						Team ScrapeTeamResponse `json:"team"`
					} `json:"standings"`
				} `json:"clasificacion"`
			} `json:"data"`
		} `json:"pageProps"`
	} `json:"props"`
}
