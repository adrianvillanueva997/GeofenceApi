build:
	go build -o app .

docker-build:
	docker build -t geofence-api .

run-dev:
	gow run main.go

run:
	go run main.go

lint:
	golangci-lint run 

test:
	go test -cover

test-dev:
	gow test -cover