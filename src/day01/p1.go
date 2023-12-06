package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, scanner := NewScannerFromFile("input.txt")
	defer file.Close()

	//sum := SolveLoop(*scanner)
	sum := SolveIndexAny(*scanner)

	fmt.Println("Solution is: ", sum)
}

func SolveIndexAny(scanner bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		start := strings.IndexAny(line, "0123456789")
		end := strings.LastIndexAny(line, "0123456789")

		str := string(line[start]) + string(line[end])
		total, _ := strconv.Atoi(str)
		sum += total
	}
	return sum
}

func SolveLoop(scanner bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		var firstDigit string
		var lastDigit string

		line := scanner.Text()
		length := len(line)
		for start, end := 0, length-1; start < length; start, end = start+1, end-1 {
			startRune := rune(line[start])
			endRune := rune(line[end])
			if firstDigit == "" && unicode.IsDigit(startRune) {
				firstDigit = string(startRune)
			}
			if lastDigit == "" && unicode.IsDigit(endRune) {
				lastDigit = string(endRune)
			}
			if firstDigit != "" && lastDigit != "" {
				break
			}
		}

		total, _ := strconv.Atoi(firstDigit + lastDigit)
		sum += total
	}
	return sum
}

// Helpers
func NewScannerFromFile(filepath string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return file, scanner
}
