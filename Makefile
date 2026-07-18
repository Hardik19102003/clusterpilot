APP=clusterpilot

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -v

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

build:
	go build -o bin/$(APP) ./cmd/cli

run:
	go run ./cmd/cli

clean:
	rm -rf bin coverage.out

ci: fmt vet test build
