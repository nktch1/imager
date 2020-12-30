service=manager

build:
	go build cmd/$(service)/main.go

run:
	go run cmd/$(service)/main.go

generate:
	protoc --gogofaster_out=plugins=grpc:api api/*.proto
