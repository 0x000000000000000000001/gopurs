#!/usr/bin/env bash
set -e

echo "Building gopurs compiler..."
npm run build

echo "Compiling tests/passing/1110.purs with spago..."
# We need to compile the test file to CoreFn first.
# Usually spago build produces output.
spago build

echo "Running gopurs on the output..."
node bin/gopurs.js --main Test1110

echo "Running the generated Go code..."
cd output
rm -f go.mod
go mod init gopurs/output
go mod tidy
go build -o app
./app
