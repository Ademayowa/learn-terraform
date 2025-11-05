# Makefile
.PHONY: test test-coverage

test:
	go test ./tests/... -v

