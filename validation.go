package main

import (
	"fmt"
	"time"
)

/*
At Limejump, we manage large batteries for our customers.
Each battery can store a certain amount of energy.
We can retrieve that power from the battery whenever we desire, for example to sell it on the Energy market.

The following code is part of an application which receives, validates and sends on a request to retrieve power from a given battery.
Each request is received via an API call and contains the following information:
The battery which we want to turn on, the desired power it is supposed to deliver and the start and
end time for when it should deliver power.

As a developer at Limejump, you have been tasked with improving the validation of an incoming request.
You have been asigned the following ticket:

Ticket description:
As Limejump, we want to ensure only sensible requests are being sent to our batteries.
Therefore, the following two checks should be performed on each incoming request:
- the start time of the request is before the end time.
- the desired power is less than the available power on the battery.
Only if both checks are successful, the request should be send on to the battery.
If not, an error should be returned.

Acceptance criteria:
- Both checks are performed inside the ValidateRequest function.
- The provided unit tests are passing.
*/

// Battery contains all the relevant information about a given battery.
type Battery struct {
	Name           string `json:"name" db:"name"`
	AvailablePower int    `json:"available_power" db:"available_power"`
}

// Request contains all the information necessary to request power from a given battery.
type Request struct {
	Battery      Battery
	Start        time.Time
	End          time.Time
	DesiredPower int
}

// ValidateRequest takes in start, end, power and battery name, performs various checks and returns an error if a check fails.
func ValidateRequest(start, end time.Time, desiredPower int, battery Battery) error {
	return fmt.Errorf("no checks implemented")
}

// HasSufficientPower checks if the desired power of the request is less or equal to the available power on the battery.
func (b Battery) HasSufficientPower(desiredPower int) bool {
	return false
}

// StartBeforeEnd checks if the start time of a request is before its end time.
func StartBeforeEnd(start, end time.Time) bool {
	return false
}
