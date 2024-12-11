// Package main implements a word search puzzle solver.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode/utf8"
)

// Directions with their corresponding Unicode arrows.
var directions = []struct {
	dx, dy int
	arrow  string
}{
	{0, 1, "→"},   // Right
	{0, -1, "←"},  // Left
	{1, 0, "↓"},   // Down
	{-1, 0, "↑"},  // Up
	{1, 1, "↘"},   // Diagonal down-right
	{-1, 1, "↗"},  // Diagonal up-right
	{1, -1, "↙"},  // Diagonal down-left
	{-1, -1, "↖"}, // Diagonal up-left
}

func main() {
	if exitCode := run(); exitCode != 0 {
		os.Exit(exitCode)
	}
}

// `run` performs the main logic and returns an exit code.
func run() int {
	if len(os.Args) != 3 {
		logError("Usage: wordsearch <word> <file path or '-'>")
		return 1
	}

	word := os.Args[1]
	puzzle, err := readPuzzle(os.Args[2])
	if err != nil {
		logError(err.Error())
		return 3
	}

	results := searchWord(puzzle, word)
	for _, result := range results {
		fmt.Println(result)
	}
	return 0
}

// Reads the puzzle from a file or standard input.
func readPuzzle(path string) ([][]rune, error) {
	var scanner *bufio.Scanner
	var file *os.File
	var err error

	if path == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err = os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer func() {
			if cerr := file.Close(); cerr != nil {
				logError(fmt.Sprintf("failed to close file: %v", cerr)) // Fixed usage
			}
		}()
		scanner = bufio.NewScanner(file)
	}

	var puzzle [][]rune
	lineLength := -1
	for scanner.Scan() {
		line := scanner.Text()
		if lineLength == -1 {
			lineLength = utf8.RuneCountInString(line)
		} else if utf8.RuneCountInString(line) != lineLength {
			return nil, fmt.Errorf("inconsistent line length")
		}
		puzzle = append(puzzle, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return puzzle, nil
}

// Logs errors to standard error.
func logError(msg string) {
	_, _ = fmt.Fprintln(os.Stderr, msg)
}

// Searches for the word in all 8 possible directions.
func searchWord(puzzle [][]rune, word string) []string {
	rows := len(puzzle)
	cols := len(puzzle[0])
	wordRunes := []rune(word)
	var results []string

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if match(puzzle, wordRunes, r, c, dir.dx, dir.dy) {
					result := fmt.Sprintf("(%d, %d) %s", r+1, c+1, dir.arrow)
					results = append(results, result)
				}
			}
		}
	}

	// Sort results lexicographically by row, column, and direction.
	sort.Strings(results)

	return results
}

// Checks if the word matches in the given direction.
func match(puzzle [][]rune, word []rune, r, c, dx, dy int) bool {
	for i := 0; i < len(word); i++ {
		nr, nc := r+i*dx, c+i*dy
		if nr < 0 || nr >= len(puzzle) || nc < 0 || nc >= len(puzzle[0]) {
			return false
		}
		if puzzle[nr][nc] != word[i] {
			return false
		}
	}
	return true
}
