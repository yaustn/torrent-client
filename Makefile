build:
	go build -o bin/main

run: build
	./bin/app

test:
	go test -v ./... -count=1
