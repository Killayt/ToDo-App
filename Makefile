.SILENT:

build:
	go get -u github.com/gin-gonic/gin
	go mod tidy
	go mod vendor
	

run: build
	go run cmd/main.go