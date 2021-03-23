package solutions

import (
	"fmt"
	"time"
)

// Test scenario:
// The api receives a start and end time, power level and asset name sent by our trading team,
// who wish to have one of our assets turn on at a certain time at a certain power level.
// This instruction request must then be checked against the asset's power capability.
// We also check if the start and end time provided are sensible.
// We also make sure the asset name provided matches an existing asset in our database.
// Only then do we return an Instruction to be sent onto the asset.

type AssetInstructor struct {
	Store Store
}

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

// CreateAndValidateInstruction returns a validated instruction.
func (a AssetInstructor) CreateAndValidateInstruction(start, end time.Time, power int, asset_name string) (Instruction, error) {
	// Get asset from database.
	asset, _ := a.Store.GetAssetByName(asset_name)

	// Check asset power capacity.
	if !asset.HasSufficientPower(power) {
		return Instruction{}, fmt.Errorf("Instruction rejected: Asset max power is %v, instructed power is %v", asset.MaxPower, power)
	}

	// Create new Instruction instance.
	i := Instruction{
		Asset: asset,
		Start: start,
		End:   end,
		Power: power,
	}

	// Check start time is before end time.
	if !i.StartBeforeEnd() {
		return Instruction{}, fmt.Errorf("Instruction rejected: Instruction start %v is after instruction end %v", i.Start, i.End)
	}

	return i, nil
}

// StartBeforeEnd checks if instruction start is before instruction end.
func (i Instruction) StartBeforeEnd() bool {
	return i.Start.Before(i.End)
}

// HasSufficientPower compares instructed power to asset's max power capacity.
func (a Asset) HasSufficientPower(power int) bool {
	return a.MaxPower >= power
}
