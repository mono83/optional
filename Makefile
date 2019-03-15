.PHONY: test bench

test:
	@go test ./...

deps:
	@go get github.com/stretchr/testify/assert
	@go get gopkg.in/yaml.v2

bench:
	@go test -bench=. -benchmem

travis: deps test
