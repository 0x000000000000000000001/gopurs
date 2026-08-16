#!/bin/bash
cd tests/runner
rm -rf output/purescript
spago build -q
node ../../bin/gopurs.js --main Main
