
BINARY_NAME=invade
OUTPUT_DIR=./bin

build:
	GOARCH=amd64 GOOS=darwin go build -o ${OUTPUT_DIR}/${BINARY_NAME}-darwin cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=linux go build -o ${OUTPUT_DIR}/${BINARY_NAME}-linux cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${OUTPUT_DIR}/${BINARY_NAME}-windows cmd/${BINARY_NAME}/main.go

clean:
	go clean
	rm ${OUTPUT_DIR}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all