package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input-p1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	file.Close()
}
