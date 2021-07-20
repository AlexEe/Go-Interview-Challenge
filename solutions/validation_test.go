package solutions

import (
	"testing"
	"time"
)

func TestBattery_AvailablePower(t *testing.T) {
	tests := []struct {
		name    string
		battery Battery
		want    int
	}{
		{
			name: "Some power used",
			battery: Battery{
				Name:      "cool_battery",
				FullPower: 5000,
				UsedPower: 4000,
			},
			want: 1000,
		},
		{
			name: "No power used",
			battery: Battery{
				Name:      "cool_battery",
				FullPower: 2000,
				UsedPower: 0,
			},
			want: 2000,
		},
		{
			name: "All power used",
			battery: Battery{
				Name:      "cool_battery",
				FullPower: 2000,
				UsedPower: 2000,
			},
			want: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := test.battery
			if got := b.AvailablePower(); got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}

func TestStartBeforeEnd(t *testing.T) {
	tests := []struct {
		name  string
		start time.Time
		end   time.Time
		want  bool
	}{
		{
			name:  "start is before end",
			start: time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 1, 1, 30, 0, 0, 0, time.UTC),
			want:  true,
		},
		{
			name:  "start is after end",
			start: time.Date(2022, 1, 1, 40, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 1, 1, 30, 0, 0, 0, time.UTC),
			want:  false,
		},
		{
			name:  "start is same as end",
			start: time.Date(2022, 1, 1, 30, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 1, 1, 30, 0, 0, 0, time.UTC),
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := StartBeforeEnd(test.start, test.end); got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

/*
Test before adding the GetBattery call.

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name       string
		request    Request
		returnsErr bool
	}{
		{
			name: "valid request: desired power is the same as available power",
			request: Request{
				Start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 500,
				Battery: Battery{
					Name:           "cool_battery",
					AvailablePower: 500,
				},
			},
			returnsErr: false,
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
			returnsErr: true,
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
			returnsErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequest(test.request)
			if err != nil && !test.returnsErr {
				t.Errorf("got error '%v' want '%v'", err, test.returnsErr)
			}
		})
	}
}
*/

// Test after adding the GetBattery call.
func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name       string
		request    Request
		returnsErr bool
	}{
		{
			name: "valid request: desired power is the same as available power",
			request: Request{
				Start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 500,
				BatteryName:  "cool_battery",
			},
			returnsErr: false,
		},
		{
			name: "invalid request: available power less than desired power",
			request: Request{
				Start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 1000,
				BatteryName:  "cool_battery",
			},
			returnsErr: true,
		},
		{
			name: "invalid request: start time after end time",
			request: Request{
				Start:        time.Date(2020, 1, 1, 30, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 500,
				BatteryName:  "cool_battery",
			},
			returnsErr: true,
		},
		{
			name: "invalid request: unknown battery name",
			request: Request{
				Start:        time.Date(2020, 1, 1, 30, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 500,
				BatteryName:  "fake_battery",
			},
			returnsErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequest(test.request)
			if err != nil && !test.returnsErr {
				t.Errorf("got error '%v' want '%v'", err, test.returnsErr)
			}
		})
	}
}
