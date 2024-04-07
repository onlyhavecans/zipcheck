all: update lint fmt test

update:
	go get -u ./...
	go mod tidy

lint:
	golangci-lint run

fmt:
	go install mvdan.cc/gofumpt@latest
	go fmt ./...
	gofumpt -w .

test:
	go test -v ./...

testloop:
	fd --glob "*.go" | entr go test -v ./...
