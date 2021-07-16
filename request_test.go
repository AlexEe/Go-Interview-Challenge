package main

import (
	"reflect"
	"testing"
	"time"
)

func TestBattery_HasSufficientPower(t *testing.T) {
	tests := []struct {
		name         string
		battery      Battery
		desiredPower int
		want         bool
	}{
		{
			// 1, To Do: Add test case where the battery has sufficient power to deliver the desired Power.
		},
		{
			// 2, To Do: Add test case where the battery has insufficient power to deliver the desired Power.
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.battery
			if got := b.HasSufficientPower(tc.desiredPower); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// 3, To Do: Create a MockStore struct to be able to mock out calls to the database.
// The MockStore should have a field called Battery which is a pointer to a Battery struct.

// 4, To Do:
// Ensure the MockStore implements the Store interface.

func TestValidator_ValidateRequest(t *testing.T) {
	testBattery := &Battery{
		Name:           "cool_battery",
		AvailablePower: 1000,
	}

	tests := []struct {
		name         string
		start        time.Time
		end          time.Time
		desiredPower int
		battery_name string
		want         Request
		wantErr      bool
	}{
		{
			name:         "valid Request",
			start:        time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
			end:          time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
			desiredPower: 500,
			battery_name: "cool_battery",
			want: Request{
				Battery:      testBattery,
				Start:        time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				DesiredPower: 500,
			},
			wantErr: false,
		},
		// 8, To Do: Add new test case which will return an error due to insufficient power.
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 5, To Do:
			// Create an instance of MockStore which takes in the testBattery.

			// 6, To Do:
			// Create an instance v of Validator which takes in the created MockStore instance.

			// 7, To Do:
			// Use v to call the ValidateRequest function with the test params and check
			// a, if an error occured
			// b, if the result from the function equals the "want" in the test case.
		})
	}
}
