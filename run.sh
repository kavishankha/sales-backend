#!/bin/bash
set -euo pipefail

# Colors for output
GREEN="\033[0;32m"
RED="\033[0;31m"
RESET="\033[0m"

# Check for Go installation
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go first.${RESET}"
    exit 1
fi

# Move to the root directory of the project (where this script is located)
cd "$(dirname "$0")"

# Install dependencies
echo -e "${GREEN}📦 Installing dependencies...${RESET}"
go mod tidy

# Optional: Uncomment to build the application
# echo -e "${GREEN}🔨 Building the application...${RESET}"
# go build -o app ./cmd

# Optional: Uncomment to run the application
# echo -e "${GREEN}🚀 Running the application...${RESET}"
# ./app

# Or just run directly without building if using `go run`
echo -e "${GREEN}🚀 Running the application with go run...${RESET}"
go run ./cmd