all: test build run

test:
	go test -v -cover -tags test -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

build:
	go build

run:
	./bcgtest
