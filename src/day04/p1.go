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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers, yourNumbers := ParseCardInfoLine(line)
		points := 0
		for _, winningNumber := range winningNumbers {
			if Contains(yourNumbers, winningNumber) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}
	return sum
}

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
