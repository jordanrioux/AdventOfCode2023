package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, scanner := NewScannerFromFile("input.txt")
	defer file.Close()

	sum := Solve(*scanner)

	fmt.Println("Solution is: ", sum)
}

func Solve(scanner bufio.Scanner) int {
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	sum := 0
	SolveRecursive(lines, 0, len(lines), &sum)
	return sum
}

var cache = map[int]int{}

func SolveRecursive(lines []string, start int, end int, sum *int) {
	for i := start; i < end; i++ {
		matchingNumbers := 0
		if value, ok := cache[i+1]; ok {
			matchingNumbers = value
		} else {
			line := lines[i]
			winningNumbers, yourNumbers := ParseCardInfoLine(line)

			for _, winningNumber := range winningNumbers {
				if Contains(yourNumbers, winningNumber) {
					matchingNumbers++
				}
			}
			cache[i+1] = matchingNumbers
		}

		if matchingNumbers != 0 {
			SolveRecursive(lines, i+1, i+1+matchingNumbers, sum)
		}
		*sum += 1
	}
}

// 1 instance of card 1
// 2 instances of card 2
// 4 instances of card 3
// 8 instances of card 4
// 14 instances of card 5
// 1 instance of card 6.

func ParseCardInfoLine(line string) ([]int, []int) {
	cardNumbersStr := strings.Split(line, ": ")[1]
	cardNumbers := strings.Split(cardNumbersStr, " | ")
	winningNumbers, yourNumbers := SplitAsInt(cardNumbers[0]), SplitAsInt(cardNumbers[1])
	return winningNumbers, yourNumbers
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

func SplitAsInt(str string) (numbers []int) {
	values := strings.Split(str, " ")
	for _, value := range values {
		if value != "" {
			numbers = append(numbers, ParseInt(value))
		}
	}
	return numbers
}

func ParseInt(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
