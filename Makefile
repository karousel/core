all: test

test:

	@# There's probably better ways to run a command without caring about
	@# the exit code, but Internet access is expensive 35000 feet in the air
	$(shell dropdb karousel_test)
	createdb karousel_test
	go test ./...

.PHONY: test
