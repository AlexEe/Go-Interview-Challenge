package solutions

import "fmt"

/*
Ticket description:
Well done for solving the first ticket!
Unfortunately, now parameters of the Request coming into our application have changed.
Previously, the Request contained a full Battery instance, now it will only contain the name of the battery.
The remaining information (FullPower and UsedPower) will need to be queried from our database.

Acceptance criteria:
- Inside validation.go: In definition of the Request struct, rename the Battery field to BatteryName and change its type to string.
- Call the GetBattery function inside ValidateRequest to retrieve the full Battery instance.
- Fix unit tests and make sure they pass.
- In validation_test.go, add a unit test to Test_ValidateRequest which returns an error because the battery name is not in the database.

Bonus:
- Suggest ways in which the functions in validation.go or store.go could be improved.
- Imagine we had a postgreSQL database instead of the hardcoded one. How would you mock out/ test a call to a real database?
*/

var database = map[string][]int{"cool_battery": {500, 0}, "awesome_battery": {1000, 0}}

// GetBatteryInformation retrieves and returns an battery's data for a given battery name.
func GetBattery(batteryName string) (Battery, error) {
	var b Battery
	if data, ok := database[batteryName]; ok {
		b.FullPower = data[0]
		b.UsedPower = data[1]
		b.Name = batteryName
		return b, nil
	}
	return Battery{}, fmt.Errorf("battery %v does not exist", batteryName)
}
