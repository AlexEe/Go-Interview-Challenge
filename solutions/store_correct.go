package solutions

import (
	"database/sql"
	"fmt"
)

// Store represents a service for retrieving battery information.
type Store interface {
	GetBatteryInformation(batteryName string) (*Battery, error)
}

// PostgresStore is the PostgreSQL database manager.
type PostgresStore struct {
	DB *sql.DB
}

// GetBatteryInformation retrieves and returns a battery's data from the db.
func (p PostgresStore) GetBatteryInformation(batteryName string) (*Battery, error) {
	// Create database query.
	query := fmt.Sprintf(`
		SELECT
			b.name,
			c.available_power,
		FROM batteries AS b
		JOIN constraints AS c
			ON b.id = c.battery_id
		WHERE
			b.name = $1
		LIMIT (1)
	`)

	// Query database.
	b := &Battery{}
	rows, _ := p.DB.Query(query)
	for rows.Next() {
		err := rows.Scan(&b.Name, &b.AvailablePower)
		if err != nil {
			return nil, fmt.Errorf("could not fetch information for battery %s", batteryName)
		}
	}
	return b, nil
}
