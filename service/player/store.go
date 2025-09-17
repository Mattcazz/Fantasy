package player

import (
	"database/sql"

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
	return nil
}

func (s *Store) DeletePlayer(player *types.Player) error {
	return nil
}

func (s *Store) PlayerDailyUpdate(player *types.Player) error {
	return nil
}
