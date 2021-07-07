package solutions

import (
	"fmt"
	"os"
	"time"
)

// Test scenario:
// At Limejump, we manage large batteries for our customers.
// Each battery can store a certain amount of energy.
// We can retrieve that power from the battery whenever we desire, for example to sell it on the Energy market.
// In the following example, we receive a request to have the battery 'cool_battery' deliver 500kW
// from 1:10 to 1:20 on 1.1.2020.
// We now need to perform a couple of checks to see if this request can be fulfilled by the selected battery.

func main() {
	// The following data is received, containing for a new request to 'cool_battery'
	// This is hard-coded, normally we would received this data via an API call.
	start := time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC)
	desiredPower := 500
	battery_name := "cool_battery"

	// Create PostgresSQL instance.
	var postgres PostgresStore

	// Create Validator.
	v := Validator{
		Store: &postgres,
	}

	// Validate a new request before sending it to the battery.
	request, err := v.ValidateRequest(start, end, desiredPower, battery_name)
	if err != nil {
		os.Exit(0)
	}

	// Send request onto owl.
	fmt.Printf("New request: %+v", request)
}
