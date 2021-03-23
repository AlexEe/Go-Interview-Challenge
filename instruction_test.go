package main

import (
	"reflect"
	"testing"
	"time"
)

type MockStore struct {
	Asset *Asset
}

func (m MockStore) GetAssetByName(assetName string) (*Asset, error) {
	return m.Asset, nil
}

func TestAssetInstructor_CreateAndValidateInstruction(t *testing.T) {
	test_asset := &Asset{
		Name:       "cool_asset",
		MaxPower:   1000,
		Technology: "battery",
	}

	type params struct {
		start      time.Time
		end        time.Time
		power      int
		asset_name string
	}

	tests := []struct {
		name    string
		params  params
		want    Instruction
		wantErr bool
	}{
		// Make this test pass.
		{
			name: "valid instruction",
			params: params{
				start:      time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				end:        time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				power:      500,
				asset_name: "cool_asset",
			},
			want: Instruction{
				Asset: test_asset,
				Start: time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
				End:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
				Power: 500,
			},
			wantErr: false,
		},
		// Add new test case which will return an error due to insufficient power.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MockStore{
				Asset: test_asset,
			}
			a := AssetInstructor{
				Store: s,
			}
			got, err := a.CreateAndValidateInstruction(tt.params.start, tt.params.end, tt.params.power, tt.params.asset_name)
			if (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
