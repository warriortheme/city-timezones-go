#!/bin/bash

# Comprehensive Test Runner for city-timezones-go
# This script runs all types of tests: unit, integration, performance, and security

set -e

echo "🧪 Running Comprehensive Test Suite for city-timezones-go"
echo "========================================================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test counters
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Function to run tests and count results
run_test() {
    local test_name="$1"
    local test_command="$2"
    
    echo -e "\n${BLUE}Running: $test_name${NC}"
    echo "Command: $test_command"
    echo "----------------------------------------"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    if eval "$test_command"; then
        echo -e "${GREEN}✅ $test_name PASSED${NC}"
        PASSED_TESTS=$((PASSED_TESTS + 1))
    else
        echo -e "${RED}❌ $test_name FAILED${NC}"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
}

# Function to run tests with coverage
run_test_with_coverage() {
    local test_name="$1"
    local test_command="$2"
    local coverage_file="$3"
    
    echo -e "\n${BLUE}Running: $test_name (with coverage)${NC}"
    echo "Command: $test_command"
    echo "----------------------------------------"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    if eval "$test_command"; then
        echo -e "${GREEN}✅ $test_name PASSED${NC}"
        PASSED_TESTS=$((PASSED_TESTS + 1))
        
        # Show coverage if file exists
        if [ -f "$coverage_file" ]; then
            echo -e "${YELLOW}Coverage Report:${NC}"
            go tool cover -func="$coverage_file" | tail -1
        fi
    else
        echo -e "${RED}❌ $test_name FAILED${NC}"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go 1.21 or later.${NC}"
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | sed 's/go//')
REQUIRED_VERSION="1.21"
if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo -e "${RED}❌ Go version $GO_VERSION is not supported. Please install Go $REQUIRED_VERSION or later.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Go version $GO_VERSION is supported${NC}"

# Clean previous test artifacts
echo -e "\n${YELLOW}Cleaning previous test artifacts...${NC}"
rm -f coverage.out coverage.html
go clean -testcache

# 1. UNIT TESTS
echo -e "\n${BLUE}📋 UNIT TESTS${NC}"
echo "============="

run_test_with_coverage "Unit Tests" "go test -v -coverprofile=coverage.out ./internal/city" "coverage.out"
run_test "Unit Tests (Race Detection)" "go test -race ./internal/city"
run_test "Unit Tests (Short)" "go test -short ./internal/city"

# 2. INTEGRATION TESTS
echo -e "\n${BLUE}🔗 INTEGRATION TESTS${NC}"
echo "=================="

run_test "Integration Tests" "go test -v ./internal/city -run TestDataLoadingIntegration"
run_test "Integration Tests" "go test -v ./internal/city -run TestSearchIntegration"
run_test "Integration Tests" "go test -v ./internal/city -run TestCachingIntegration"
run_test "Integration Tests" "go test -v ./internal/city -run TestConcurrencyIntegration"

# 3. PUBLIC API TESTS
echo -e "\n${BLUE}🌐 PUBLIC API TESTS${NC}"
echo "=================="

run_test "Public API Tests" "go test -v ./pkg/citytimezones"

# 4. PERFORMANCE TESTS
echo -e "\n${BLUE}⚡ PERFORMANCE TESTS${NC}"
echo "===================="

run_test "Benchmark Tests" "go test -bench=. -benchmem ./internal/city"
run_test "Benchmark Tests (Parallel)" "go test -bench=. -benchmem -cpu=1,2,4,8 ./internal/city"

# 5. SECURITY TESTS
echo -e "\n${BLUE}🔒 SECURITY TESTS${NC}"
echo "=================="

run_test "Security Tests" "go test -v ./internal/city -run TestValidation"
run_test "Security Tests" "go test -v ./internal/city -run TestErrorHandlingIntegration"

# 6. EDGE CASE TESTS
echo -e "\n${BLUE}🎯 EDGE CASE TESTS${NC}"
echo "=================="

run_test "Edge Case Tests" "go test -v ./internal/city -run TestUnmarshalCityData_EdgeCases"
run_test "Edge Case Tests" "go test -v ./internal/city -run TestLimitsIntegration"

# 7. MEMORY TESTS
echo -e "\n${BLUE}💾 MEMORY TESTS${NC}"
echo "==============="

run_test "Memory Tests" "go test -v ./internal/city -run TestCacheMemoryUsage"
run_test "Memory Tests" "go test -v ./internal/city -run TestPerformanceIntegration"

# 8. CONCURRENCY TESTS
echo -e "\n${BLUE}🔄 CONCURRENCY TESTS${NC}"
echo "====================="

run_test "Concurrency Tests" "go test -v ./internal/city -run TestConcurrencyIntegration"
run_test "Concurrency Tests" "go test -race ./internal/city -run TestConcurrencyIntegration"

# 9. CODE QUALITY CHECKS
echo -e "\n${BLUE}🔍 CODE QUALITY CHECKS${NC}"
echo "======================="

run_test "Code Formatting" "go fmt ./..."
run_test "Code Linting" "go vet ./..."
run_test "Code Linting (Static Analysis)" "go test -v ./internal/city -run TestLimitsPerformance"

# 10. DOCUMENTATION TESTS
echo -e "\n${BLUE}📚 DOCUMENTATION TESTS${NC}"
echo "======================="

run_test "Documentation Tests" "go doc ./pkg/citytimezones"
run_test "Documentation Tests" "go doc ./internal/city"

# 11. BUILD TESTS
echo -e "\n${BLUE}🔨 BUILD TESTS${NC}"
echo "==============="

run_test "Build Test (Internal)" "go build ./internal/city"
run_test "Build Test (Public API)" "go build ./pkg/citytimezones"
run_test "Build Test (CLI)" "go build ./cmd/citytimezones"
run_test "Build Test (Examples)" "go build ./examples/basic"
run_test "Build Test (Examples)" "go build ./examples/advanced"

# 12. MODULE TESTS
echo -e "\n${BLUE}📦 MODULE TESTS${NC}"
echo "==============="

run_test "Module Dependencies" "go mod tidy"
run_test "Module Dependencies" "go mod verify"
run_test "Module Dependencies" "go mod download"

# Generate coverage report
if [ -f "coverage.out" ]; then
    echo -e "\n${YELLOW}📊 COVERAGE REPORT${NC}"
    echo "=================="
    go tool cover -func=coverage.out
    echo ""
    go tool cover -html=coverage.out -o coverage.html
    echo -e "${GREEN}Coverage HTML report generated: coverage.html${NC}"
fi

# Final Results
echo -e "\n${BLUE}📊 TEST RESULTS SUMMARY${NC}"
echo "========================"
echo -e "Total Tests: ${BLUE}$TOTAL_TESTS${NC}"
echo -e "Passed: ${GREEN}$PASSED_TESTS${NC}"
echo -e "Failed: ${RED}$FAILED_TESTS${NC}"

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "\n${GREEN}🎉 ALL TESTS PASSED! 🎉${NC}"
    echo -e "${GREEN}Your code is ready for production!${NC}"
    exit 0
else
    echo -e "\n${RED}❌ SOME TESTS FAILED${NC}"
    echo -e "${RED}Please fix the failing tests before proceeding.${NC}"
    exit 1
fi
