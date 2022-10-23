.SILENT:

build:
	go mod tidy
	go mod vendor
	

run: build
	go run cmd/main.go