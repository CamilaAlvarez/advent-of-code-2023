package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_7/parser"
	"github.com/CamilaAlvarez/advent-of-code-2023/Day_7/sort"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	hands := parser.ParseHands(file)
	fmt.Println(hands)
	fmt.Println()
	sort.SortHands(hands)
	var totalWinnings int
	for k, v := range hands {
		rank := k + 1
		totalWinnings += (rank * v.Bid)
	}
	fmt.Println("Total winnings:", totalWinnings)
}
