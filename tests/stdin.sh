#!/usr/bin/env bash
set -euo pipefail

go build ./cmd/wordsearch

echo "Test 1 (stdin)"
./wordsearch de - < puzzle1.txt > test.out
diff test.out tests/out1.txt
echo PASS
