#!/bin/bash
set -e

# ========================================
# Satudata API - Deployment Script
# ========================================

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  Satudata API - Production Deployment  ${NC}"
echo -e "${GREEN}========================================${NC}"

# Check required tools
echo -e "${YELLOW}[1/5] Checking prerequisites...${NC}"
command -v go >/dev/null 2>&1 || { echo -e "${RED}Go is required but not installed.${NC}"; exit 1; }
command -v docker >/dev/null 2>&1 || { echo -e "${RED}Docker is required but not installed.${NC}"; exit 1; }

# Load environment
if [ -f .env ]; then
    echo -e "${GREEN}  Loading environment from .env file${NC}"
    export $(grep -v '^#' .env | xargs)
else
    echo -e "${YELLOW}  Warning: No .env file found, using defaults${NC}"
fi

# Build
echo -e "${YELLOW}[2/5] Building production binary...${NC}"
go build -ldflags="-s -w" -o bin/satudata-api ./cmd/api/main.go
echo -e "${GREEN}  Build complete: bin/satudata-api${NC}"

# Run tests
echo -e "${YELLOW}[3/5] Running tests...${NC}"
go test -v -count=1 ./... || {
    echo -e "${RED}Tests failed! Aborting deployment.${NC}"
    exit 1
}
echo -e "${GREEN}  All tests passed${NC}"

# Build Docker image
echo -e "${YELLOW}[4/5] Building Docker image...${NC}"
docker build -t satudata-api:latest .
echo -e "${GREEN}  Docker image built: satudata-api:latest${NC}"

# Deploy
echo -e "${YELLOW}[5/5] Deploying...${NC}"
echo -e "${GREEN}"
echo "========================================"
echo "  Deployment Complete!"
echo "========================================"
echo ""
echo "  To start the service:"
echo "    docker compose up -d"
echo ""
echo "  To view logs:"
echo "    docker compose logs -f"
echo ""
echo "  To stop:"
echo "    docker compose down"
echo -e "${NC}"
