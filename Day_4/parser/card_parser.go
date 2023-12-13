package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

const cardNameSeparator string = ":"
const cardNumbersSeparator string = "|"
const numbersSeparator string = " "

type Card struct {
	Id                   string
	SortedWinningNumbers []int
	SortedNumbersYouHave []int
}

func fillNumbersList(numbers string) []int {
	var winningNumbers []int
	// Use regex when spliting
	numbers = strings.Trim(numbers, numbersSeparator)
	for _, v := range strings.Split(numbers, numbersSeparator) {
		v = strings.Trim(v, numbersSeparator)
		fmt.Println(numbers)
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Invalid card number: ", v)
		}
		winningNumbers = append(winningNumbers, number)
	}
	sort.Slice(winningNumbers, func(i, j int) bool {
		return winningNumbers[i] < winningNumbers[j]
	})
	return winningNumbers
}

func ParseCards(file io.Reader) []Card {
	var cards []Card
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			fmt.Println("Skipping empty line")
			continue
		}
		splitLine := strings.Split(line, cardNameSeparator)
		if len(splitLine) < 2 {
			log.Fatal("Invalid card line: ", line)
		}
		splitCardNumbers := strings.Split(splitLine[1], cardNumbersSeparator)
		if len(splitCardNumbers) < 2 {
			log.Fatal("Invalid card numbers line: ", splitLine[1])
		}
		winningNumbers := fillNumbersList(splitCardNumbers[0])
		myNumbers := fillNumbersList(splitCardNumbers[1])
		cards = append(cards, Card{splitLine[0], winningNumbers, myNumbers})
	}
	return cards
}
