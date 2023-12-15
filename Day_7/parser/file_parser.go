package parser

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Card int

const (
	Two Card = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	T
	J
	Q
	K
	A
)

type Hand struct {
	Cards [5]Card
	Bid   int
}

func ParseHands(file io.Reader) []Hand {
	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		splitLine := strings.Split(line, " ")
		if len(splitLine) < 2 {
			log.Fatal("Invalid line:", line)
		}
		var hand Hand
		cards := splitLine[0]
		bid := splitLine[1]
		b, err := strconv.Atoi(bid)
		if err != nil {
			log.Fatal("Cannot parse number:", bid)
		}
		hand.Bid = b
		for i, v := range cards {
			switch v {
			case '2':
				hand.Cards[i] = Two
			case '3':
				hand.Cards[i] = Three
			case '4':
				hand.Cards[i] = Four
			case '5':
				hand.Cards[i] = Five
			case '6':
				hand.Cards[i] = Six
			case '7':
				hand.Cards[i] = Seven
			case '8':
				hand.Cards[i] = Eight
			case '9':
				hand.Cards[i] = Nine
			case 'T':
				hand.Cards[i] = T
			case 'J':
				hand.Cards[i] = J
			case 'Q':
				hand.Cards[i] = Q
			case 'K':
				hand.Cards[i] = K
			case 'A':
				hand.Cards[i] = A
			default:
				log.Fatal("Invalid card:", v)

			}
		}
		hands = append(hands, hand)
	}
	return hands
}
