build:
	go build -o bin/http/app cmd/http/main.go

run:
	go run cmd/http/main.go

clean:
	rm -rf bin

dev:
	~/go/bin/air
