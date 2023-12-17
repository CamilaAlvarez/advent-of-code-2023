package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_9/parser"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	report := parser.ParseReport(file)
	predictions := report.ComputePredictions()
	var sum int
	for _, p := range predictions {
		sum += p
	}
	fmt.Println("Prediction sum:", sum)

}
