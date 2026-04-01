build:
	CGO_ENABLED=0 go build -o grubstake ./cmd/grubstake/

run: build
	./grubstake

test:
	go test ./...

clean:
	rm -f grubstake

.PHONY: build run test clean
