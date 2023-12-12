package main

import (
	"log"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_2/parser"
	"github.com/CamilaAlvarez/advent-of-code-2023/Day_2/part"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}

	games := parser.ParseInputFile(file)
	part.Part2(games)
}
