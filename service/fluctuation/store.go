package fluctuation

import (
	"database/sql"

	"github.com/Mattcazz/Fantasy.git/types"
)

type Store struct {
	db *sql.DB
}

 func NewFluctuationStore(db *sql.DB) Store {
	return Store{db: db}
}

func (s *Store) InsertFluctuation(tx *sql.Tx, fluctuation *types.Fluctuation) error {
	return nil
}

func (s *Store) GetFluctuationHistoryFromPlayer(player_id int) ([]types.Fluctuation, error) {
	return nil, nil
}

func (s *Store) GetLastFluctuationFromPlayer(player_id int) (*types.Fluctuation, error) {
	return nil, nil
}
