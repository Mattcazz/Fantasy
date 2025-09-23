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
	Id       int
	Team_id  int
	Name     string
	Points   int
	Value    float32
	Avg      float32
	Img_url  string
	WebID    int
	Status   string
	Position string
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

type ScrapeTeamResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type TeamRoot struct {
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

type PlayerRoot struct {
	Players []ScrapePlayer `json:"players"`
}

type ScrapePlayer struct {
	Name   string          `json:"nm"`
	Team   PlayerTeamRes   `json:"t"`
	Points PlayerPointsRes `json:"pt"`
	WebID  int             `json:"id"`
	Value  PlayerValueRes  `json:"m"`
	PosId  int             `json:"pid"`
	Status PlayerStatusRes `json:"gs"`
}

type PlayerTeamRes struct {
	Id int `json:"id"`
}

type PlayerPointsRes struct {
	Points float32 `json:"p"`
	Avg    float32 `json:"a"`
}

type PlayerValueRes struct {
	Value float32 `json:"vm"`
}

type PlayerStatusRes struct {
	Status string `json:"st"`
}
