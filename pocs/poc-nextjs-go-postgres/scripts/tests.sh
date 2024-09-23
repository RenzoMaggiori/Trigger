#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
BOLD='\033[1m'
RESET='\033[0m'
BOLDCYAN='\033[1;36m'
BOLDBLUE='\033[1;34m'
BOLDGREEN='\033[1;32m'
BOLDMAGENTA='\033[1;35m'

total_coverage=0
total_packages=0
total_failed=0
total_tests=0

disable_logger() {
    go run -exec "lib.DisableLogger()" > /dev/null 2>&1
}

enable_logger() {
    go run -exec "lib.EnableLogger()" > /dev/null 2>&1
}

packages=$(go list ./... | grep -v /vendor/)

run_tests() {
    package=$1

    echo -e "${BOLDCYAN}Testing package: ${BLUE}${package}${RESET}"

    output=$(go test -v -coverprofile=coverage.out "$package")
    failed=$(echo "$output" | grep -- '--- FAIL' | wc -l)
    tests_run=$(echo "$output" | grep '^=== RUN' | wc -l)

    total_failed=$((total_failed + failed))
    total_tests=$((total_tests + tests_run))

    echo -e "${BOLDMAGENTA}Test Results:${RESET}"
    echo "$output" | grep -E '^=== RUN|PASS|FAIL'

    if [ -f coverage.out ]; then
        coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
        echo -e "${GREEN}Coverage:${RESET} ${coverage}% with ${tests_run} tests, ${failed} failed."
        total_coverage=$(echo "$total_coverage + $coverage" | bc)
        total_packages=$((total_packages + 1))
        rm coverage.out
    else
        echo -e "${RED}No coverage data found for ${package}.${RESET}"
    fi
    echo -e "${BOLD}${YELLOW}-------------------------------------------------------${RESET}"
}

disable_logger

for package in $packages; do
    run_tests "$package"
done

enable_logger

if [ "$total_packages" -gt 0 ]; then
    average_coverage=$(echo "scale=2; $total_coverage / $total_packages" | bc)
else
    average_coverage=0
fi

echo -e "${BOLD}${BLUE}======================= Summary =======================${RESET}"
echo -e "${CYAN}Total Packages Tested:${RESET} $total_packages"
echo -e "${CYAN}Total Tests Run:${RESET} $total_tests"
echo -e "${CYAN}Total Tests Failed:${RESET} $total_failed"
echo -e "${CYAN}Overall Coverage:${RESET} ${average_coverage}%"
echo -e "${BOLD}${BLUE}=======================================================${RESET}"

exit $total_failed
