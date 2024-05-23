# Project Name
PROJECT_NAME = quakes

# Directories
SRC_DIR = ./cmd/$(PROJECT_NAME)
BIN_DIR = ./bin

# Commands
GO = go
GOCMD = $(GO) mod tidy
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean
GOTEST = $(GO) test
GOVET = $(GO) vet
GOLINT = golangci-lint run
GOINSTALL = $(GO) install
BINARY_NAME = $(BIN_DIR)/$(PROJECT_NAME)

# All targets
.PHONY: all build clean run test lint vet tidy

all: tidy vet lint test build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SRC_DIR)

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run the project
run:
	$(GOCMD)
	$(GO) run $(SRC_DIR)

# Run tests
test:
	$(GOTEST) -v ./...

# Run linter
lint:
	$(GOLINT)

# Run vet
vet:
	$(GOVET) ./...

# Tidy up dependencies
tidy:
	$(GOCMD)
