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
	var sumPoints int
	cards := parser.ParseCards(file)
	for _, v := range cards {
		var i, j, points int
		winningNumbers := v.SortedWinningNumbers
		numbersYouHave := v.SortedNumbersYouHave
		for i < len(winningNumbers) && j < len(numbersYouHave) {
			// if the winning number is higher than the nu
			if winningNumbers[i] > numbersYouHave[j] {
				j++
			} else if winningNumbers[i] < numbersYouHave[j] {
				i++
			} else {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				i++
				j++
			}
		}
		fmt.Println("Point for card ", v.Id, ": ", points)
		sumPoints += points
	}
	fmt.Println("Total points: ", sumPoints)

}
