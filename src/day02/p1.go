package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var CubesForGame = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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
		gameNumber, gameSets := ParseGameInfoLine(line)
		possible := true
	outer:
		for _, gameSetStr := range gameSets {
			specificGameSet := strings.Split(gameSetStr, ", ")
			for _, cubeInfoStr := range specificGameSet {
				cubeInfos := strings.Split(cubeInfoStr, " ")
				count, color := ParseInt(cubeInfos[0]), cubeInfos[1]
				if count > CubesForGame[color] {
					possible = false
					break outer
				}
			}
		}

		if possible {
			sum += gameNumber
		}
	}
	return sum
}

func ParseGameInfoLine(line string) (int, []string) {
	// Game 1: 7 green, 4 blue, 3 red; 4 blue, 10 red, 1 green; 1 blue, 9 red
	gameFullInfos := strings.Split(line, ": ")
	gameInfoStr, allGameSetsStr := gameFullInfos[0], gameFullInfos[1]
	gameNumber := ParseInt(strings.Split(gameInfoStr, " ")[1])
	gameSets := strings.Split(allGameSetsStr, "; ")
	return gameNumber, gameSets
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
