package player

import (
	"database/sql"
	"fmt"

	"github.com/Mattcazz/Fantasy.git/types"
)

type Store struct {
	db *sql.DB
}

func NewPlayerStore(db *sql.DB) Store {
	return Store{
		db: db,
	}
}

func (s *Store) InsertPlayerTx(tx *sql.Tx, player *types.Player) error {
	query := `INSERT INTO player (team_id, name, points, price,avg,goals, assists, web_id )
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`

	row, err := tx.Query(query,
		player.Team_id, player.Name, player.Points, player.Value,
		player.Avg, player.Goals, player.Assists, player.WebID)

	if err != nil {
		return err
	}

	for row.Next() {
		err = scanPlayerRow(row, player)
	}

	if err != nil {
		return fmt.Errorf("error inserting player tx: %s", err.Error())
	}

	return nil
}

func (s *Store) DeletePlayer(player *types.Player) error {
	return nil
}

func (s *Store) PlayerDailyUpdate(player *types.Player) error {
	return nil
}

func scanPlayerRow(row *sql.Rows, player *types.Player) error {
	return row.Scan(
		&player.Id,
		&player.Team_id,
		&player.Name,
		&player.Points,
		&player.Value,
		&player.Avg,
		&player.Goals,
		&player.Assists,
		&player.Img_url,
		&player.WebID,
	)
}
