test:
	go test ./test

build:
	go build -o bin/main cmd/main.go

build-run:
	./bin/main

run:
	go run cmd/main.go
