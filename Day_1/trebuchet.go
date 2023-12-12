package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Invalid number of arguments: %d\n", len(os.Args))
		os.Exit(1)
	}
	sum := 0
	inputFile := os.Args[1]
	contents, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Could not open input file")
		os.Exit(1)
	}
	defer contents.Close()
	scanner := bufio.NewScanner(contents)
	for scanner.Scan() {
		line := scanner.Text()
		numbersInLine := make([]int, 0, len(line))
		for _, v := range line {
			if unicode.IsDigit(v) {
				numbersInLine = append(numbersInLine, int(v-'0'))
			}
		}
		if len(numbersInLine) <= 0 {
			continue
		}
		firstDigit := numbersInLine[0]
		lastDigit := numbersInLine[len(numbersInLine)-1]
		sum += (firstDigit*10 + lastDigit)
	}
	fmt.Printf("Sum of calibration values: %d\n", sum)

}
