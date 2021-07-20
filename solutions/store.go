package solutions

import "fmt"

/*
Ticket description:
The parameters of the request coming into our application have changed!
We no longer receive a complete Battery struct, including the available power, but only the battery's name.
The available power for each battery must instead be retrieved with a call to our database.

Acceptance criteria:
- Inside validation.go: In the Request struct, rename the Battery field to BatteryName and change its type to string.
- Call the GetBattery function inside ValidateRequest to retrieve the battery with its available power.
- Fix unit tests and make sure they pass.
- In validation_test.go, add a unit test to Test_ValidateRequest which returns an error because the battery name is not in the database.

Bonus:
- Suggest ways in which the functions in validation.go or store.go could be improved.
- Imagine we had a postgreSQL database instead of the hardcoded one. How would you mock out/ test a call to a real database?
*/

var database = map[string]int{"cool_battery": 500, "awesome_battery": 1000}

// GetBatteryInformation retrieves and returns an battery's data for a given battery name.
func GetBattery(batteryName string) (Battery, error) {
	var b Battery
	if availablePower, ok := database[batteryName]; ok {
		b.AvailablePower = availablePower
		b.Name = batteryName
		return b, nil
	}
	return Battery{}, fmt.Errorf("battery %v does not exist", batteryName)
}
