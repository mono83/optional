.PHONY: test bench

test:
	@go test ./...

deps:
	@go get github.com/stretchr/testify/assert

bench:
	@go test -bench=. -benchmem

travis: deps test
