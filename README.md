# Pair programming challenge for junior and mid-level developers
The following challenge has been designed with junior to mid-level engineers in mind.
Moreover, the challenge is supposed to give candidates an idea about the work done by Limejump.
The test scenario mimicks a small part of Asset Control's functionality: a request to switch a battery on has been received and now needs to be validated before being sent to the battery.
The test as such focusses on Go features which are used frequently in the infrastructure team's codebase (structs, interfaces, unit tests, mocking) in favour of other features which we use less frequently (concurrency, channels, recursive functions).
You can find all files with correct unit tests and idiomatic go in the solutions folder.

## What is being tested?
- Basic Go concepts: structs, interfaces, unit tests, error checking and mocking.
- Go naming conventions: encapsulation, CamelCase.
- Writing unit tests for existing functions.
- Using interfaces to mock out a database call.
- Refactoring non-idiomatic code.
- Optional: SQL.

## What are we sending to interviewees and when?
We are sending them the request.go, request_test.go, main.go, store.go files at the start of the interview.

## Interview progression and possible questions
Generally, be kind and helpful. Pay attention to how they deal with being stuck: Are they able to ask specific questions?
The interview is designed in multiple steps which should be followed in order: Depending on how fast/ experienced the developer is, advance to the next step or spend more time on the first ones. The difficulty increases with every additional step.

1. Basic knowledge: idiomatic Go, implementing structs and writing methods/ functions:
- First explain general scenario by showing the main.go file: Trader input for potential Request (with params battery name, start, end, desired power) needs to be validated.
- Go to `request.go` file:
- Have them explain code back to you and ask if they have questions.
- Ask them to write the `Battery.HasSufficientPower()` method that is called in `ValidateRequest()`.
- Ask them to create a new Request instance in the `ValidateRequest()` function.
- Ask them to add the `Start_before_end()` check to the `ValidateRequest()` function.
- Ask them how this function could be made into more idiomatic Go code. (CamelCase)
- If they don't suggest it themselves, ask them to re-write the function to be a method on `Request`.
- Ask more generally how they might refactor the code in the `request.go` file (see list below).

2. Writing unit tests:
- Go to the `request_test.go` file.
- Ask them to complete the tests for the `Battery.HasSufficientPower` method.
- Write a test where the battery has sufficient power and one where the battery does not have sufficient power.

3. For more advanced candidates: Mocking a database call using interfaces
- Go to request_test.go file and have them explain how the MockStore works.
- Let them run the tests: First test fails because end is before start.
- Let them add a second test which fails because instructed power exceeds the available battery power.
- Move onto `store.go` file and let them read through code. Ask to have it explained it back to you.
- Ask how they would test the `GetBatteryInformation()` function.
- We want them to suggest creating a mock database using the `Store` interface to mock out the db return values.
- If they don't get there: Ask about `Store` interface and the purpose of an interface.

4. Optional: SQL questions.
- What is a foreign key and where is it used in the SQL query? (c.battery_id)
- What kind of join is taking place on line 26? (Inner join)

## All issues in the request.go file
We are not expecting them to find all of these, but this should provide ample opportunity to showcase knowledge about idiomatic Go.
- All functions could be unexported since they are all in the `main` package.
- All functions, since they are exported, should have docstring.
- Error returned by `GetBatteryInformation()` is not checked.
- `Start_before_end()` function should be CamelCase.
- `Start_before_end()` function should be a method on Request instead of taking Request as a parameter.
