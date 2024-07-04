# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOBUILD=$(GOCMD) build
BINARY_NAME=clearprice-client-sdk

# Directories
SRC_DIR=.
TEST_DIR=./...

# Default target executed when no arguments are provided to make
default: build

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(SRC_DIR)/main.go

# Run the tests
test:
	$(GOTEST) $(TEST_DIR) -v

# Format the code
fmt:
	$(GOFMT) $(SRC_DIR)/...

# Vet the code
vet:
	$(GOVET) $(SRC_DIR)/...

# Clean up the generated files
clean:
	if [ -f $(BINARY_NAME) ] ; then rm $(BINARY_NAME) ; fi

# Run the application
run: build
	./$(BINARY_NAME)

# Generate the binary and run the tests
all: fmt vet test build

.PHONY: all build test fmt vet clean run
