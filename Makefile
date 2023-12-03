build:
	go build -o bin/main cmd/main.go

run: build
	./bin/main

clean:
	rm -rf bin

test:
	go test -v ./... -count=1
