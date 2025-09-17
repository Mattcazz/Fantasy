package team

import (
	"database/sql"
	"fmt"

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
	query := "SELECT * FROM team WHERE name = ?"

	row, err := s.db.Query(query, name)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		return scanTeamRow(row)
	}

	return nil, fmt.Errorf("the search came up with no results")

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

func scanTeamRow(row *sql.Rows) (*types.Team, error) {
	team := new(types.Team)

	err := row.Scan(
		&team.Id,
		&team.Name,
		&team.Logo_url,
	)

	return team, err
}
