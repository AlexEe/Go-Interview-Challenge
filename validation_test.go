package main

import (
	"fmt"
	"testing"
	"time"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name    string
		request Request
		wantErr error
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
			name: "invalid request: available power less than desired power",
			request: Request{
				Start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 1000,
				Battery: Battery{
					Name:           "cool_battery",
					AvailablePower: 500,
				},
			},
			wantErr: fmt.Errorf("available power 500 less than desired power 1000"),
		},
		{
			name: "invalid request: start time after end time",
			request: Request{
				Start:        time.Date(2020, 1, 1, 30, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 500,
				Battery: Battery{
					Name:           "cool_battery",
					AvailablePower: 500,
				},
			},
			wantErr: fmt.Errorf("start time after end time"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequest(test.request)
			if err != test.wantErr {
				t.Errorf("got error '%v', want error '%v'", err, test.wantErr)
				return
			}
		})
	}
}
