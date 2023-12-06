package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var spelledOutNumbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// Could be dynamically generated based on previous map
var reversedSpelledOutNumbers = map[string]string{
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func main() {
	file, scanner := NewScannerFromFile("input-p2.txt")
	defer file.Close()

	sum := Solve(*scanner)

	fmt.Println("Solution is: ", sum)
}

func Solve(scanner bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := FindFirstDigitInLine(line, spelledOutNumbers)
		lastDigit := FindFirstDigitInLine(Reverse(line), reversedSpelledOutNumbers)
		total, _ := strconv.Atoi(firstDigit + lastDigit)
		sum += total
	}
	return sum
}

func FindFirstDigitInLine(line string, spelledOutNumbers map[string]string) string {
	for i, length := 0, len(line); i < length; i++ {
		c := rune(line[i])
		if unicode.IsDigit(c) {
			return string(c)
		}
		index := strings.IndexAny(line, "0123456789")
		str := line[:index]

		var currentValue string
		currentLowestIndex := math.MaxInt8
		for spelledOutNumber, value := range spelledOutNumbers {
			index := strings.Index(str, spelledOutNumber)
			if index != -1 && index < currentLowestIndex {
				currentLowestIndex = index
				currentValue = value
			}
		}

		if currentValue != "" {
			return currentValue
		}
	}
	return ""
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

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
