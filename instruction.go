package main

import (
	"fmt"
	"time"
)

// Test scenario:
// The api receives a start and end time, power level and asset name sent by our trading team,
// who wish to have one of our assets turn on at a certain time at a certain power level.
// This potential instruction must then be checked against the asset's capability (MaxPower).
// We also check if the start and end time provided are correct.
// We also make sure the asset name provided matches an existing asset in our database.
// Only then do we return an Instruction to be sent onto the asset.

type Instruction struct {
	Asset *Asset
	Start time.Time
	End   time.Time
	Power int
}

type Asset struct {
	Name       string `json:"name" db:"name"`
	MaxPower   int    `json:"max_power" db:"max_power"`
	Technology string `json:"technology" db:"technology"`
}

// A function that takes in start, end, power and asset name, performs various checks and returns a validated instruction.
func CreateAndValidateInstruction(start, end time.Time, power int, asset_name string) (Instruction, error) {
	store := PostgresStore{}
	asset, _ := store.GetAssetByName(asset_name)
	if !asset.HasSufficientPower(power) {
		return Instruction{}, fmt.Errorf("Instruction rejected: Asset max power is %v, instructed power is %v", asset.MaxPower, power)
	}

	instruction := Instruction{
		Asset: asset,
		Start: start,
		End:   end,
		Power: power,
	}

	// Add 'StartBeforeEnd' check here.

	return instruction, nil
}

func Start_before_end(i Instruction) bool {
	return i.Start.Before(i.End)
}

func (a Asset) HasSufficientPower(power int) bool {
	result := a.MaxPower >= power
	return result
}
