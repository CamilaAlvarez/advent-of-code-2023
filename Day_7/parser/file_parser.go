package parser

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Card int
type Strength int

const numberCards int = 13
const HandSize int = 5
const (
	J Card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	T
	Q
	K
	A
)

const (
	HighCard Strength = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards    [HandSize]Card
	Bid      int
	Strength Strength
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
		var cardCounter [numberCards]int
		for i, v := range cards {
			switch v {
			case '2':
				hand.Cards[i] = Two
				cardCounter[Two]++
			case '3':
				hand.Cards[i] = Three
				cardCounter[Three]++
			case '4':
				hand.Cards[i] = Four
				cardCounter[Four]++
			case '5':
				hand.Cards[i] = Five
				cardCounter[Five]++
			case '6':
				hand.Cards[i] = Six
				cardCounter[Six]++
			case '7':
				hand.Cards[i] = Seven
				cardCounter[Seven]++
			case '8':
				hand.Cards[i] = Eight
				cardCounter[Eight]++
			case '9':
				hand.Cards[i] = Nine
				cardCounter[Nine]++
			case 'T':
				hand.Cards[i] = T
				cardCounter[T]++
			case 'J':
				hand.Cards[i] = J
				cardCounter[J]++
			case 'Q':
				hand.Cards[i] = Q
				cardCounter[Q]++
			case 'K':
				hand.Cards[i] = K
				cardCounter[K]++
			case 'A':
				hand.Cards[i] = A
				cardCounter[A]++
			default:
				log.Fatal("Invalid card:", v)

			}
		}
		var counter, sum, highestNumberCards int
		// We're using J in the way that gives the most advantage
		if cardCounter[0] == 5 {
			// We just need to check this individual case
			hand.Strength = FiveOfAKind
		} else {
			for _, v := range cardCounter[1:] {
				sum += v
				if v != 0 {
					counter++
				}
				if highestNumberCards < v {
					highestNumberCards = v
				}
				if sum > 0 && sum == 5-cardCounter[0] {
					switch counter {
					case 1:
						hand.Strength = FiveOfAKind
					case 2:
						if highestNumberCards == (4 - cardCounter[0]) {
							hand.Strength = FourOfAKind
						} else {
							hand.Strength = FullHouse
						}
					case 3:
						if highestNumberCards == (3 - cardCounter[0]) {
							hand.Strength = ThreeOfAKind
						} else {
							hand.Strength = TwoPair
						}
					case 4:
						hand.Strength = OnePair
					case 5:
						hand.Strength = HighCard
					}
				}
			}
		}

		hands = append(hands, hand)
	}
	return hands
}
