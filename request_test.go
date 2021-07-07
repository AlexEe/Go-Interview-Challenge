package main

import (
	"reflect"
	"testing"
	"time"
)

func TestBattery_HasSufficientPower(t *testing.T) {
	battery := Battery{
		Name:           "cool_battery",
		AvailablePower: 100,
	}

	tests := []struct {
		name         string
		battery      Battery
		desiredPower int
		want         bool
	}{
		{
			// Add test case where the battery has sufficient power to deliver the desired Power.
		},
		{
			// Add test case where the battery has insufficient power to deliver the desired Power.
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

type MockStore struct {
	Battery *Battery
}

func (m MockStore) GetBatteryInformation(batteryName string) (*Battery, error) {
	return m.Battery, nil
}

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
		// Make this test pass.
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
		// Add new test case which will return an error due to insufficient power.
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := MockStore{
				Battery: testBattery,
			}
			v := Validator{
				Store: s,
			}
			got, err := v.ValidateRequest(tc.start, tc.end, tc.desiredPower, tc.battery_name)
			if (err != nil) != tc.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
