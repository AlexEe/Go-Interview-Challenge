package main

import (
	"database/sql"
	"fmt"
)

/*
Ticket description:
The parameters of of the ValidateRequest functions are changing!
Our application is now only receiving the name of the battery which is to be turned on, not the complete Battery struct.
The necessary information about this battery (its available power) need to be retrieved from a database.

Acceptance criteria:
- ValidateRequest takes in parameter called batteryName instead of battery which is a string.
- Call the GetBattery function inside ValidateRequest to retrieve the battery with its available power.
- Hint: To declare the PostgresStore simply write `p := PostgresStore{}`
*/

// PostgresStore is the PostgreSQL database manager.
type PostgresStore struct {
	DB *sql.DB
}

// GetBatteryInformation retrieves and returns an battery's data for a given battery name.
func (p PostgresStore) GetBattery(batteryName string) (Battery, error) {
	// Create database query.
	query := fmt.Sprintf(`
		SELECT
			b.name,
			c.max_power,
		FROM batteries AS b
		JOIN constraints AS c
			ON b.id = c.battery_id
		WHERE
			b.name = $1
		LIMIT (1)
	`)

	// Query database.
	b := Battery{}
	rows, err := p.DB.Query(query)
	for rows.Next() {
		err = rows.Scan(&b.Name, &b.AvailablePower)
		if err != nil {
			return Battery{}, fmt.Errorf("could not fetch battery by with name %s", batteryName)
		}
	}
	return b, nil
}
