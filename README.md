Word Search

![Wind Chill Chart](https://upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Wordsearch.svg/632px-Wordsearch.svg.png)

## Overview

[Word search](https://en.wikipedia.org/wiki/Word_search) is a game children play.  You find words in a rectangular grid.  A word can occur from right to left, left to right, and so on.  More precisely a matching word can start at any character in the puzzle and extend in the direction of any of its 8-connected neighbors.  A word can occur multiple times (in many different directions and locations).  The game is often used to help train the human brain to see patterns in character sequences.  In this assignment we will allow unicode characters encoded as UTF-8 in the puzzle.

## Learning Objectives

- Parse standard input
- Reading for a file
- Output formatting
- Unicode
- Looping and searching

## Requirements

1. Write a Go program and put your file(s) in `cmd/wordsearch/`.

1. The program's first argument is the word to search for in the puzzle.  The second argument is either the file path to the puzzle or "-" which means the program will read the puzzle from standard input.  The puzzle is a UTF-8 encoded text file with one line per row in the puzzle.  The puzzle is rectangular so all rows must have the same number of columns.  If they do not then write "inconsistent line length" to standard error and exit with exit code 3.

1. Output the matches one per line in the following format.  `(R, C) D` where `R` is the row (1-based), `C` is the column (1-based), and `D` is the unicode direction arrow from the set `→←↓↑↘↗↙↖`.  The matches must be output in lexicographic order on the vector `RCD`.

1. Be sure to close of file descriptor.  You will loose a point for each file descriptor not closed.

1. Do not use any library other than the Go standard library.

1. The source code must compile with the most recent version of the Go compiler.

1. The program must not panic under any circumstances.

1. Make sure your code is "gofmt'd".  See `gofmt` or better use `goimports` or better yet configure IDE to do this formatting on file save.

## Hints

- The `bufio`, `fmt`, and `strings` packages might be useful for this assignment.
- Use the `Makefile`.  For example, you can run tests with `make test -B`.  To make sure your code is properly formatted run `make fix`.

## Example Output

```shell
$ cat puzzle1.txt
abcdefg
hiŮklmn
pqœstuv

$ cat puzzle2.txt
cxxxefg
bxfxhi0
cxxxc10

$ go run ./cmd/wordsearch de puzzle1.txt 
(1, 4) →

$ go run ./cmd/wordsearch œke puzzle1.txt 
(3, 3) ↗

$ go run ./cmd/wordsearch cx puzzle2.txt 
(1, 1) →
(1, 1) ↘
(3, 1) →
(3, 1) ↗
(3, 5) ←
(3, 5) ↖

$ go run ./cmd/wordsearch i puzzle2.txt 
(2, 6) ←
(2, 6) ↑
(2, 6) →
(2, 6) ↓
(2, 6) ↖
(2, 6) ↗
(2, 6) ↘
(2, 6) ↙
```

You may use `./gen-puzzle.sh` to create random puzzles of any size.


