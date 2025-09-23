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

	query := "SELECT * FROM teams WHERE $1 LIKE '%' || name || '%'"
	row, err := s.db.Query(query, name)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		team := new(types.Team)
		err = scanTeamRow(row, team)
		return team, err
	}

	return nil, fmt.Errorf("the search came up with no results for %s", name)

}

func (s *Store) InsertTeamTx(tx *sql.Tx, team *types.Team) error {
	query := `INSERT INTO teams (name, logo_url, web_id)
				VALUES ($1, $2, $3) RETURNING *`

	row, err := tx.Query(query, team.Name, team.Logo_url, team.Web_Id)

	if err != nil {
		return fmt.Errorf("error executing query -> %s", err.Error())
	}

	for row.Next() {
		err = scanTeamRow(row, team)
	}

	if err != nil {
		return fmt.Errorf("error inserting team tx -> %s", err.Error())
	}

	return nil

}

func (s *Store) DeleteTeam(team *types.Team) error {
	query := `DELETE FROM teams WHERE id = $1`

	_, err := s.db.Query(query, team.Id)

	return err
}

func scanTeamRow(row *sql.Rows, team *types.Team) error {

	err := row.Scan(
		&team.Id,
		&team.Name,
		&team.Logo_url,
		&team.Web_Id,
	)

	return err
}
