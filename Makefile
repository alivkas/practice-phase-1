.PHONY: build lint test clean
build:
	go build -o bin/tool ./cmd/tool
lint:
	golangci-lint run ./...
test:
	go test -race -v ./...
clean:
	rm -rf bin/ *.prof coverage.out