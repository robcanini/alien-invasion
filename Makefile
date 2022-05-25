
BINARY_NAME=invade
OUTPUT_DIR=./bin
TEST_DIR=./test

.PHONY: build test clean

build:
	GOARCH=amd64 GOOS=darwin go build -o ${OUTPUT_DIR}/${BINARY_NAME}-darwin cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=linux go build -o ${OUTPUT_DIR}/${BINARY_NAME}-linux cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${OUTPUT_DIR}/${BINARY_NAME}-windows cmd/${BINARY_NAME}/main.go

clean:
	go clean
	rm -rf ${OUTPUT_DIR}

test:
	go test ${TEST_DIR}

test_coverage:
	go test ${TEST_DIR} -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all
