package solutions

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
- Add the missing logic to the incomplete functions 'AvailablePower' and 'StartBeforeEnd'.
- Add the checks (available power sufficient and start before end) to the ValidateRequest function.
- If a check fails, a descriptive error should be returned.
- Make sure the provided unit tests are passing.

Tipp:
- You can use any means you normally would to look up information about a method or package.
- Any other questions, just ask!
*/

// Battery contains all the relevant information about a given battery.
type Battery struct {
	Name      string
	FullPower int // FullPower describes the amount a battery can provide at full power, its maximum capacity.
	UsedPower int // UsedPower describes the amount which has already been used or consumed.
}

// Request contains all the information necessary to request power from a given battery.
type Request struct {
	Battery      Battery
	Start        time.Time
	End          time.Time
	DesiredPower int
}

// ValidateRequest performs various checks on a given request and returns an error if a check fails.
func ValidateRequest(r Request) error {
	if r.Battery.AvailablePower() < r.DesiredPower {
		return fmt.Errorf("available power is less than desired power")
	}

	if !StartBeforeEnd(r.Start, r.End) {
		return fmt.Errorf("start time after end time")
	}

	return nil
}

// AvailablePower returns the available power of a given battery.
func (b Battery) AvailablePower() int {
	return b.FullPower - b.UsedPower
}

// StartBeforeEnd checks if the start time of a request is before its end time.
func StartBeforeEnd(start, end time.Time) bool {
	return start.Before(end)
}
