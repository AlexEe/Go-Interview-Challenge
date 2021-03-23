package solutions

import (
	"fmt"
	"os"
	"time"
)

// Should be retrieved from Env Vars or Config file
const (
	hostname = "localhost"
	hostport = 5432
	username = "postgres"
	password = "password"
	dbname   = "assets"
)

func main() {
	// Data received via http POST request.
	start := time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC)
	power := 500
	asset_name := "cool_assset"

	// Connect to PostgresSQL.
	var postgres PostgresStore
	err := postgres.Open(hostport, hostname, username, password, dbname)
	if err != nil {
		os.Exit(0)
	}
	defer postgres.DB.Close()

	// Create Asset Instructor
	a := AssetInstructor{
		Store: &postgres,
	}

	// Create and validate new instruction
	instruction, err := a.CreateAndValidateInstruction(start, end, power, asset_name)
	if err != nil {
		os.Exit(0)
	}
	fmt.Printf("New instruction: %+v", instruction)
}
