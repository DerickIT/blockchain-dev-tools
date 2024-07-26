# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=bdt
BINARY_UNIX=$(BINARY_NAME)_unix

# Build parameters
BUILD_DIR=bin
MAIN_PATH=cmd/bdt/main.go

# Make parameters
MAKE=make

all: test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

test-integration:
	$(GOTEST) -v ./tests/...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v $(MAIN_PATH)
	./$(BUILD_DIR)/$(BINARY_NAME)

deps:
	$(GOGET) -v -t -d ./...
	$(GOMOD) tidy

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v $(MAIN_PATH)

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_mac -v $(MAIN_PATH)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME).exe -v $(MAIN_PATH)

# Installation
install: build
	mv $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

# Docker
docker-build:
	docker build -t $(BINARY_NAME):latest .

# Help
help:
	@echo "make - compile the source code and run tests"
	@echo "make build - compile the source code"
	@echo "make test - run tests"
	@echo "make clean - remove binary files and cached files"
	@echo "make run - build and run the binary"
	@echo "make deps - get dependencies and tidy go.mod"
	@echo "make build-linux - cross compile for linux"
	@echo "make build-mac - cross compile for mac"
	@echo "make build-windows - cross compile for windows"
	@echo "make install - install the binary to /usr/local/bin/"
	@echo "make docker-build - build docker image"

.PHONY: all build test clean run deps build-linux build-mac build-windows install docker-build help