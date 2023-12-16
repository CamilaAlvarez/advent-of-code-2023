package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_8/parser"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	desertMap := parser.ParseToMap(file)
	fmt.Println(desertMap)
	fmt.Println()

	startingPoints := desertMap.StartingPoints()
	fmt.Println(startingPoints)
	fmt.Println()
	numberSteps := desertMap.NumberOfSteps(startingPoints)
	fmt.Println("Number of steps:", numberSteps)
}
