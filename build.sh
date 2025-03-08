#!/bin/bash

# Create bin directory if it doesn't exist
mkdir -p bin

echo "Building Arachnid tools..."

# Build main crawler
echo "Building main crawler..."
go build -o bin/arachnid cmd/arachnid/main.go

# Build Cogni
echo "Building Cogni..."
go build -o bin/cogni cmd/cogni/cogni.go

# Build PDF Bandit
echo "Building PDF Bandit..."
go build -o bin/pdf-bandit cmd/pdf-bandit/pdf_bandit.go

echo "Build complete! Binaries are in the bin/ directory" 