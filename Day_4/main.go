package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_4/parser"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	cards := parser.ParseCards(file)
	cardAmounts := make([]int, len(cards))
	for k := range cards {
		cardAmounts[k] = 1
	}
	for k, v := range cards {
		var i, j, winningCount int
		winningNumbers := v.SortedWinningNumbers
		numbersYouHave := v.SortedNumbersYouHave
		for i < len(winningNumbers) && j < len(numbersYouHave) {
			// if the winning number is higher than the nu
			if winningNumbers[i] > numbersYouHave[j] {
				j++
			} else if winningNumbers[i] < numbersYouHave[j] {
				i++
			} else {
				winningCount++
				i++
				j++
			}
		}
		for i := 0; i < winningCount; i++ {
			cardAmounts[k+i+1] = cardAmounts[k+i+1] + cardAmounts[k]
		}
		fmt.Println("Total amount of ", v.Id, " cards: ", cardAmounts[k])
	}
	var sumCards int
	for _, v := range cardAmounts {
		sumCards += v
	}
	fmt.Println("Number of scratchcards: ", sumCards)

}
