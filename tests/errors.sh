#!/usr/bin/env bash

# set -euo pipefail

go build ./cmd/wordsearch

expectFailure() {
    echo Exit code: "$1"
    ([ "$1" -ne "0" ] && echo "✅ Test passed.") || (echo "❌ Test did not pass." && exit 1)
}

echo "Test wrong arguments"
./wordsearch
expectFailure $?
echo PASS

echo "Test bad puzzle"
go build ./cmd/wordsearch
./wordsearch a tests/bad-puzzle.txt > test.out 2> test.err
code=$?
if [ "$code" -eq "3" ]; then
    echo "✅ Test passed."
else
    echo "❌ Expected status code of 3 but got $code."
    exit 1
fi
if grep 'inconsistent line length' test.err > /dev/null; then
    echo "✅ Test passed."
else
    echo "❌ Expected status code of 3."
    exit 1
fi


