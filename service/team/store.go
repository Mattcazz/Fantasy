package team

import (
	"database/sql"

	"github.com/Mattcazz/Fantasy.git/types"
)

type Store struct {
	db *sql.DB
}

func NewTeamStore(db *sql.DB) Store {
	return Store{
		db: db,
	}
}

func (s *Store) GetTeamByName(name string) (*types.Team, error) {
	return nil, nil
}

func (s *Store) InsertTeamTx(tx *sql.Tx, team *types.Team) error {
	return nil
}

func (s *Store) DeleteTeam(team *types.Team) error {
	return nil
}

func (s *Store) AddPlayerToTeamTx(tx *sql.Tx, player_id, team_id int) error {
	return nil
}
