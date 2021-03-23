package main

import (
	"database/sql"
	"fmt"
)

type Store interface {
	GetAssetByName(asset string) (*Asset, error)
}

type PostgresStore struct {
	DB *sql.DB
}

func (p *PostgresStore) Open(hostport int, hostname, username, password, dbname string) error {
	conn := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostport, hostname, username, password, dbname)
	var err error
	p.DB, err = sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer p.DB.Close()

	err = p.DB.Ping()
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *PostgresStore) GetAssetByName(assetName string) (*Asset, error) {
	// Create database query.
	query := fmt.Sprintf(`
		SELECT
			a.name,
			t.type AS technology,
			a.max_power,
		FROM assets AS a
		JOIN technologies AS t
			ON a.technology_id = t.id
		WHERE
			a.name = $1
		LIMIT (1)
	`)

	// Query database.
	a := &Asset{}
	rows, err := p.DB.Query(query)
	for rows.Next() {
		err = rows.Scan(&a.Name, &a.MaxPower, &a.Technology)
		if err != nil {
			return nil, fmt.Errorf("could not fetch asset by with name %s", assetName)
		}
	}
	return a, nil
}
