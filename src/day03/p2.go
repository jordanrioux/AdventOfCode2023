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

	sum := Solve(*scanner)

	fmt.Println("Solution is: ", sum)
}

func RemoveNonDigits(r rune) bool {
	return !unicode.IsDigit(r)
}

func Solve(scanner bufio.Scanner) int {
	var lines []string
	var numbersPerLine [][]string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		numbersPerLine = append(numbersPerLine, strings.FieldsFunc(line, RemoveNonDigits))
	}

	matriceHeight := len(lines)
	matriceWidth := len(lines[0])

	type GearInfo struct {
		Count   int
		Numbers []int
	}

	gears := map[string]GearInfo{}
	for y, line := range lines {
		previousIndex := 0
		for _, number := range numbersPerLine[y] {
			numberSize := len(number)
			index := IndexAt(line, number, previousIndex)
			previousIndex = index + numberSize
			found, i, j := CheckIfRectangleTouchingSymbols(lines, index, y, numberSize, matriceWidth, matriceHeight, number)
			if found {
				key := strconv.Itoa(i) + "|" + strconv.Itoa(j)
				value, contains := gears[key]

				if !contains {
					info := new(GearInfo)
					info.Count = 1
					info.Numbers = append(info.Numbers, ParseInt(number))
					gears[key] = *info
				} else {
					value.Count++
					value.Numbers = append(value.Numbers, ParseInt(number))
					gears[key] = value
				}
			}
		}
	}

	sum := 0
	for _, info := range gears {
		if info.Count == 2 {
			sum += info.Numbers[0] * info.Numbers[1]
		}
	}
	return sum
}

func CheckIfRectangleTouchingSymbols(matrice []string, x int, y int, rectWidth int, matriceWidth int, matriceHeight int, number string) (bool, int, int) {
	if x < 0 {
		return false, -1, -1
	}

	x0 := max(x-1, 0)
	x1 := min(x+rectWidth, matriceWidth-1)
	y0 := max(y-1, 0)
	y1 := min(y+1, matriceHeight-1)

	for i := y0; i <= y1; i++ {
		for j := x0; j <= x1; j++ {
			r := rune(matrice[i][j])
			if !unicode.IsDigit(r) && r != '.' {
				return true, i, j
			}
		}
	}
	return false, -1, -1
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

func ParseInt(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}

func IndexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}
