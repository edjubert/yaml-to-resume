# Makefile for the YAML to Resume project

# Go binary
BINARY = bin/yaml-to-resume

# Go source files
SRC = cmd/main.go

# Scripts folder
SCRIPT_FOLDER = scripts/

# Build the application
build:
	@echo Building project to $(BINARY)
	@go build -o $(BINARY) $(SRC)

# Run the application
all: build
	@echo Generating all resume
	@$(SCRIPT_FOLDER)/generate-all.sh

# Clean up generated files
clean:
	rm -f $(BINARY)
	rm -rf output/*
	rm -rf tmp/*

# Install dependencies
deps:
	go mod tidy

# Generate PDF resumes
generate:
	go run cmd/main.go --config=default --lang=en --template=default
	go run cmd/main.go --config=default --lang=fr --template=default

.PHONY: build run clean deps generate
