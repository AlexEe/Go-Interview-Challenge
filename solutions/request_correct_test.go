package solutions

import (
	"reflect"
	"testing"
	"time"
)

func TestBattery_HasSufficientPower(t *testing.T) {
	testBattery := Battery{
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
			name:         "battery has sufficient power",
			battery:      testBattery,
			desiredPower: 100,
			want:         true,
		},
		{
			name:         "battery has insufficient power",
			battery:      testBattery,
			desiredPower: 200,
			want:         false,
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

// MockStore mocks out a PostgresStore.
type MockStore struct {
	Battery *Battery
}

// GetBatteryByName is mocked out here to implement the Store interface.
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
		{
			name:         "valid Request",
			start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
			end:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
			desiredPower: 500,
			battery_name: "cool_battery",
			want: Request{
				Battery:      testBattery,
				Start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				End:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				DesiredPower: 500,
			},
			wantErr: false,
		},
		{
			name:         "invalid Request, insufficient power",
			start:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
			end:          time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
			desiredPower: 1500,
			battery_name: "cool_battery",
			want:         Request{},
			wantErr:      true,
		},
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
