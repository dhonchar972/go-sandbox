.DEFAULT_GOAL := build
fmt:
	go fmt ./...
.PHONY:fmt
lint: fmt
	golint ./...
.PHONY:lint
vet: fmt
	go vet ./...
.PHONY:vet
build: vet
	go build main.go
.PHONY:build


#hello:
#	echo "hello"
#

#BINARY_NAME=hello-world
#
#build:
# GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
# GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
# GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go
#
#run:
#	./${BINARY_NAME}
#
#build_and_run: build run
#
#clean:
#	go clean
#	rm ${BINARY_NAME}-darwin
#	rm ${BINARY_NAME}-linux
#	rm ${BINARY_NAME}-windows
