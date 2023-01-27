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

build-docker: build
	docker build . -t api-rest

run-docker: build-docker
	docker run -p 3000:3000 api-rest

test-docker:
	go test -v ./internal/tests/container_test.go
