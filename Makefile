all: lint fmt test

update:
	go get -u ./...
	go mod tidy

lint:
	go vet ./...
	golangci-lint run

fmt:
	go fmt ./...

test:
	go test -v ./...

testloop:
	fd --glob "*.go" | entr go test -v ./...
