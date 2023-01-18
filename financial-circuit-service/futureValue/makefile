build:
	go build -o bin/futureValue ./main.go

run:
	go run ./main.go

test-circuit:
	go test -v ./internal/futureValue/test/futureValue_test.go

format: 
	go fmt ./...

vet: 	
	go vet ./...

lint: 
	staticcheck ./...
