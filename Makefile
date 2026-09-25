dev:
	air
run:
	go run .
build:
	go build
test:
	go test -json -cover ./... | gotestfmt
fmt:
	golangci-lint fmt
lint:
	golangci-lint run
fix:
	golangci-lint run --fix
clean:
	go clean
	git clean -dfX
