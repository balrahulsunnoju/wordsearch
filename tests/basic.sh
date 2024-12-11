#!/usr/bin/env bash
set -euo pipefail

go build ./cmd/wordsearch

echo "Test 1"
./wordsearch de puzzle1.txt > test.out
diff test.out tests/out1.txt
echo PASS

echo "Test 2"
./wordsearch œke puzzle1.txt > test.out
diff test.out tests/out2.txt
echo PASS

echo "Test 3"
./wordsearch cx puzzle2.txt > test.out
diff test.out tests/out3.txt
echo PASS

echo "Test 4"
./wordsearch i puzzle2.txt > test.out
diff test.out tests/out4.txt
echo PASS

echo "Test 5"
./wordsearch not-there puzzle1.txt > test.out
diff test.out tests/out5.txt
echo PASS

echo "Test 6"
./wordsearch HlJR puzzle3.txt > test.out
diff test.out tests/out6.txt
echo PASS
