.PHONY: test bench

test:
	@go test ./...

bench:
	@go test -bench=. -benchmem
