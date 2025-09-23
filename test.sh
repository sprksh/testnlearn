#!/bin/bash
cd /Users/sp/rudderstack/codes/testnlearn
go mod tidy
go build -o testnlearn .
echo "Testing version command:"
./testnlearn version
echo ""
echo "Testing run command (will sleep for 3 seconds):"
./testnlearn run
