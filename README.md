# Pair programming challenge for junior and mid-level developers
The following challenge has been designed with junior to mid-level engineers in mind. Moreover, the challenge is supposed to give candidates an idea about the work done by the infra-team and thus test very concretely how they would deal with working on a project such as Asset Control. Therefore, it tests knowledge of Go features which are used frequently in our codebase (structs, interfaces, unit tests, mocking) in favour of other features which we use less frequently (concurrency, channels, recursive functions).

## What is being tested?
- Basic Go concepts: structs, interfaces, unit tests, error checking and mocking.
- Go naming conventions: encapsulation, CamelCase.
- Writing unit tests for existing functions.
- Using interfaces to mock out a database call.
- Refactoring non-idiomatic code.
- Optional: SQL.

## What are we sending to interviewees and when?
We are sending them the instruction.go, instruction_test.go, main.go, store.go files at the start of the interview.

## Interview progression and possible questions
Generally, be kind and helpful. Pay attention to how they deal with being stuck: Are they able to ask specific questions?
- First explain general scenario by showing the main.go file: Trader input for potential instruction (with params asset name, start, end, power) needs to be validated.
- Go to `instructions.go` file:
- Have them explain code back to you and ask if they have questions.
- Ask them to create a new Instruction instance in the `CreateAndValidateInstruction()` function.
- Ask them to add the `Start_before_end()` check to the `CreateAndValidateInstruction()` function.
- Ask them to write a unit test for the `Start_before_end()` check.
- Ask them how this function could be made into more idiomatic Go code. (CamelCase)
- If they don't suggest it themselves, ask them to re-write the function to be a method on `Instruction`.
- Ask more generally how they might refactor the code in the `instructions.go` file (see list below).
- Move onto `store.go` file and let them read through code. Ask to have it explained it back to you.
- Ask how they would test the `GetAssetByName()` function.
- We want them to suggest creating a mock database using the `Store` interface to mock out the db return values.
- If they don't get there: Ask about `Store` interface and the purpose of an interface.
- Go to instruction_test.go file and have them explain how the MockStore works.
- Let them run the tests: First test fails because end is before start.
- Let them add a second test which fails because instructed power exceeds max asset power.
- Optional: SQL questions.
- What is a foreign key and where is it used in the SQL query? (technologies.id)
- What kind of join is taking place on line 54? (Inner join)

## All issues in the instructions.go file
We are not expecting them to find all of these, but this should provide ample opportunity to showcase knowledge about idiomatic Go.
- All functions could be unexported since they are all in the `main` package.
- All functions, since they are exported, should have docstring.
- `CreateAndValidateInstruction()` could be 2 functions.
- `CreateAndValidateInstruction()` should have documentation for separate steps.
- Error returned by `GetAssetByName()` is not checked.
- `CreateAndValidateInstruction()` check for StartBeforeEnd needs to be added.
- `Start_before_end()` function should be CamelCase.
- `Start_before_end()` function should be a method on Instruction instead of taking Instruction as a parameter.
- `HasSufficientPower()` could be a single line.

## All issues in the main.go file
- Postgres connection details should be retrieved from config file or env vars.
