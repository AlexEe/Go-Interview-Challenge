package solutions

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

// ValidateRequest returns a validated request.
func (v Validator) ValidateRequest(start, end time.Time, desiredPower int, battery_name string) (Request, error) {
	// Get battery from database.
	battery, err := v.Store.GetBatteryInformation(battery_name)
	if err != nil {
		return Request{}, fmt.Errorf("Request rejected: %v", err)
	}
	// Check the available power on the battery.
	if !battery.HasSufficientPower(desiredPower) {
		return Request{}, fmt.Errorf("Request rejected: battery max power is %v, instructed power is %v", battery.AvailablePower, desiredPower)
	}

	// Create new Request instance.
	request := Request{
		Battery:      battery,
		Start:        start,
		End:          end,
		DesiredPower: desiredPower,
	}

	// Check start time is before end time.
	if !request.StartBeforeEnd() {
		return Request{}, fmt.Errorf("Request rejected: Request start %v is after Request end %v", request.Start, request.End)
	}

	return request, nil
}

// StartBeforeEnd checks if Request start is before Request end.
func (inst Request) StartBeforeEnd() bool {
	return inst.Start.Before(inst.End)
}

// HasSufficientPower compares the desired power to battery's available power.
func (b Battery) HasSufficientPower(desiredPower int) bool {
	return b.AvailablePower >= desiredPower
}
