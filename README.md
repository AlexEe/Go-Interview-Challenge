# Pair programming challenge for junior and mid-level developers
The following challenge has been designed with junior to mid-level engineers in mind.
Moreover, the challenge is supposed to give candidates an idea about the work done by Limejump.
The test scenario mimicks a small part of Asset Control's functionality: a request to switch a battery on has been received and now needs to be validated before being sent to the battery.
The test as such focusses on Go features which are used frequently in the infrastructure team's codebase (structs, interfaces, unit tests) in favour of other features which we use less frequently (concurrency, channels, recursive functions).
You can find all files with correct unit tests and idiomatic go in the solutions folder.

## What is being tested?
- Basic Go concepts: structs, unit tests, refactoring, error checking and mocking.
- Optional questions for mocking and interfaces.

## What are we sending to interviewees and when?
We are sending them the validation.go, validation_test.go, store.go files at the start of the interview.

## Interview progression
Generally, be kind and helpful if they have questions.
For go-specific questions, remind them that they are allowed to use Google/ Stackoverflow etc.

1. First ticket: basic Go knowledge (structs, logic, error checking)
- Start with the `validation.go` file
- Candidates should read through the instructions themselves. After they are done, ask if they understand the task and have any questions.
- If they are stuck, suggest running the tests. There are unit tests for each of the three functions in the `validation_test.go` file.
- If they need help with the time.Time package, suggest googling before providing them with the answer.
- Once all tests pass, move on to the `store.go` file.
- IMPORTANT: If they are stuck with this ticket, make sure to solve the challenges here first and do NOT move on to (or let them know about) the second ticket. Junior candidates are not expected to be able to fix both tickets in time.

2. Second ticket: common Go knowledge (refactoring, unit tests)
- Have them read through the instructions in `store.go` for the second ticket, ask if they have questions.
- After all unit tests pass (and new unit test has been added), move on to the bonus questions.

3. Bonus questions: more advanced Go knowledge (improving own code, best practices, mocking, interfaces)
3.1 Suggest ways in which the functions in validation.go or store.go could be improved.
- Depends on how they approached the task, but some likely improvements could be:
- Make code cleaner: write `AvailablePower()` and `StartBeforeEnd()` in a single line.
- Make code easier: `AvailablePower` could be a field on the `Battery` struct.
- Make code more idiomatic: `StartBeforeEnd` could be a method on the `Request` struct.

3.2 Imagine we had a postgreSQL database in `store.go` instead of the hardcoded one. How would you mock out a call to a real database?
Most common way to mock out a database call is by using an interface, e.g. called `Store`. This interface would have the `GetBattery` function on it.
Every struct that has a method called `GetBattery` would then implement the `Store` interface.
For the unit tests, a `MockStore` struct could be created which has a function named `GetBattery` but returns mocked out results.

## Interview expectations
### Junior developers (<= 1 year experience)
- Ask questions of the interviewer.
- Use Google et al to close knowledge gaps.
- Run and understand unit tests.
- General understanding of structs and basic logic.
- Should be able to finish parts or all of Ticket 1, with help.

### Mid-level developers (2-3 year experience)
- All of the above.
- Ability to improve their own code, knowledge about best practices.
- Should be able to finish Ticket 1 and most of Ticket 2.
- If time permits, should have at least partial answers to the Bonus questions.

### Senior-level developers (+4 years experience)
- Interview has been designed for junior to mid-level. However, if no other interview challenge is available:
- All of the above.
- Should be able to confidently suggest ideas to improve code quality.
- Should understand the concept of mocking and interfaces in Go.
- Should confidently finish both tickets and answer Bonus questions.
