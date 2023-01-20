test-circuit:
	go test -v ./private-location-and-membership-circuit/location/test/flatdistance_test.go

test-setup:
	go test -v ./private-location-and-membership-circuit/location/test/export_test.go

format: 
	go fmt ./...

vet: 	
	go vet ./...
