package main

import (
	"fmt"
	"time"
)

// Test scenario:
// At Limejump, we manage large batteries for our customers.
// Each battery can store a certain amount of energy.
// We can retrieve that power from the battery whenever we desire, for example to sell it on the Energy market.
// In the following example, we receive a request to have the battery 'cool_battery' deliver 500kW
// from 1:10 to 1:20 on 1.1.2020.

// We now need to perform a couple of checks to see if this request can be fulfilled by the selected battery.
// 1, We check if the desired power can be delivered by the battery.
// 2, We check if the start and end time provided are sensible.
// We also make sure the battery name provided matches an existing battery in our database.
// Only then do we return a Request to be sent onto the battery.

// Validator represents a service for validating requests.
type Validator struct {
	Store Store
}

// Request contains all the information necessary to request power from a given battery.
type Request struct {
	Battery      *Battery
	Start        time.Time
	End          time.Time
	DesiredPower int
}

// Battery contains all the relevant information about a given battery.
type Battery struct {
	Name           string `json:"name" db:"name"`
	AvailablePower int    `json:"available_power" db:"available_power"`
}

// A function that takes in start, end, power and battery name, performs various checks and returns a validated request.
func (v Validator) ValidateRequest(start, end time.Time, desiredPower int, battery_name string) (Request, error) {
	// Retrieve battery information from our database.
	battery, _ := v.Store.GetBatteryInformation(battery_name)

	// 1, To Do: Define this function below.
	if !battery.HasSufficientPower(desiredPower) {
		return Request{}, fmt.Errorf("Request rejected: available power is %v, desired power is %v", battery.AvailablePower, desiredPower)
	}

	// 2, To Do: Create new Request instance here.
	var request Request

	// 3, To Do: Call the 'Start_before_end' function here.

	return request, nil
}

// 1, To Do: Write method HasSufficientPower here.
// The method should be on the Battery struct.
// This method should take in a parameter: the desired power of the request which the battery is asked to deliver.
// The method should return a boolean of true if the battery has sufficient power to deliver the desired power,
// and false if it cannot deliver the desired power.

func Start_before_end(r Request) bool {
	return r.Start.Before(r.End)
}
