BINARY=web
BUILD_DIR=./build
SRC=./cmd/web/main.go

all: build

.PHONY: build
build:
	@go build -o $(BUILD_DIR)/${BINARY} $(SRC)

.PHONY: start
start: build
	@${BUILD_DIR}/$(BINARY)

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test ./...

.PHONY: clean
clean:
	@go clean
	@rm -rf $(BUILD_DIR)

