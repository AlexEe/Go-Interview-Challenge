package main

import (
	"fmt"
	"testing"
	"time"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name         string
		request      Request
		start        time.Time
		end          time.Time
		desiredPower int
		battery      Battery
		wantErr      error
	}{
		{
			name: "valid request",
			request: Request{
				Start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 500,
				Battery: Battery{
					Name:           "cool_battery",
					AvailablePower: 500,
				},
			},

			wantErr: nil,
		},
		{
			name:         "invalid request: available power less than desired power",
			start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
			end:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
			desiredPower: 1000,
			battery: Battery{
				Name:           "cool_battery",
				AvailablePower: 500,
			},
			wantErr: fmt.Errorf("available power 500 less than desired power 1000"),
		},
		{
			name:         "invalid request: start time after end time",
			start:        time.Date(2020, 1, 1, 30, 0, 0, 0, time.UTC),
			end:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
			desiredPower: 1000,
			battery: Battery{
				Name:           "cool_battery",
				AvailablePower: 500,
			},
			wantErr: fmt.Errorf("start time after end time"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequest(test.start, test.end, test.desiredPower, test.battery)
			if err != test.wantErr {
				t.Errorf("got error '%v', want error '%v'", err, test.wantErr)
				return
			}
		})
	}
}
