package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type rank int

const (
	HIGHCARD rank = iota
	ONEPAIR
	TWOPAIRS
	THREE
	STRAIGHT
	FLUSH
	FULLHOUSE
	FOUR
	STRAIGHTFLUSH
	ROYALFLUSH
)

type suit byte

type card struct {
	value int
	Suit  suit
}

type result int

const (
	draw result = iota
	winnerFirst
	winnerSecond
)

type game struct {
	firstHand  hand
	secondHand hand
}

type hand []card

func (s suit) String() string {
	switch s {
	case 'H':
		return "\u2665"
	case 'D':
		return "\u2666"
	case 'S':
		return "\u2660"
	case 'C':
		return "\u2663"
	default:
		panic("unknown suit")
	}
}

func (h hand) Len() int {
	return len(h)
}

func (h hand) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (g *game) findResult() result {
	first := g.firstHand
	sort.Sort(first)

	second := g.secondHand
	sort.Sort(second)

	rankFirst := evaluateHand(first)
	rankSecond := evaluateHand(second)

	if rankFirst > rankSecond {
		return winnerFirst
	}

	if rankFirst < rankSecond {
		return winnerSecond
	}

	switch rankFirst {
	case ROYALFLUSH:
		return draw
	case STRAIGHTFLUSH:
		return highCard(first, second, 0)
	case FOUR:
		return four(first, second)
	case FULLHOUSE:
		return fullHouse(first, second)
	case FLUSH:
		return highCard(first, second, 0)
	case STRAIGHT:
		return highCard(first, second, 0)
	case THREE:
		return three(first, second)
	case TWOPAIRS:
		return twoPairs(first, second)
	case ONEPAIR:
		return onePair(first, second)
	default:
		return highCard(first, second, 0)
	}
}

func readGameFile(fileName string) ([]game, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("File opening failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var games []game

	for scanner.Scan() {
		g := readGame(scanner.Text())
		games = append(games, g)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return games, nil
}

func readGame(line string) game {
	var cards []card

	for _, value := range strings.Split(line, " ") {
		cards = append(cards, readCard(value))
	}

	return game{
		firstHand:  cards[0:5],
		secondHand: cards[5:10],
	}
}

func readCard(s string) card {
	var c card

	switch s[0] {
	case 'A':
		c.value = 14
	case 'J':
		c.value = 11
	case 'Q':
		c.value = 12
	case 'K':
		c.value = 13
	case 'T':
		c.value = 10
	default:
		c.value = int(s[0] - '0')
	}

	c.Suit = suit(s[1])

	return c
}

func sameSuit(h hand) bool {
	s := h[0].Suit

	for _, value := range h {
		if value.Suit != s {
			return false
		}
	}

	return true
}

func findPair(h hand) int {
	pair := 0

	for i := 0; i < h.Len(); i++ {
		for j := i; j < h.Len(); j++ {
			if i != j && h[i].value == h[j].value {
				pair++
			}
		}
	}

	return pair
}

func (h hand) isRoyalFlush() bool {
	return sameSuit(h) && h[0].value == 14 && h[1].value == 13 &&
		h[2].value == 12 && h[3].value == 11 && h[4].value == 10
}

func (h hand) isStraightFlush() bool {
	return sameSuit(h) && h[0].value == h[1].value+1 && h[1].value == h[2].value+1 &&
		h[2].value == h[3].value+1 && h[3].value == h[4].value+1
}

func (h hand) isFour() bool {
	return h[1].value == h[2].value && h[2].value == h[3].value &&
		(h[0].value == h[1].value || h[3].value == h[4].value)
}

func (h hand) isFullHouse() bool {
	return h[0].value == h[1].value && h[3].value == h[4].value &&
		(h[2].value == h[1].value || h[2].value == h[3].value)
}

func (h hand) isFlush() bool {
	return sameSuit(h)
}

func (h hand) isStraight() bool {
	return h[0].value == h[1].value+1 && h[1].value == h[2].value+1 &&
		h[2].value == h[3].value+1 && h[3].value == h[4].value+1
}

func (h hand) isThree() bool {
	return (h[0].value == h[1].value && h[1].value == h[2].value) ||
		(h[1].value == h[2].value && h[2].value == h[3].value) ||
		(h[2].value == h[3].value && h[3].value == h[4].value)
}

func (h hand) isTwoPairs() bool {
	return findPair(h) == 2
}

func (h hand) isOnePair() bool {
	return findPair(h) == 1
}

func evaluateHand(h hand) rank {
	switch {
	case h.isRoyalFlush():
		return ROYALFLUSH
	case h.isStraightFlush():
		return STRAIGHTFLUSH
	case h.isFour():
		return FOUR
	case h.isFullHouse():
		return FULLHOUSE
	case h.isFlush():
		return FLUSH
	case h.isStraight():
		return STRAIGHT
	case h.isThree():
		return THREE
	case h.isTwoPairs():
		return TWOPAIRS
	case h.isOnePair():
		return ONEPAIR
	default:
		return HIGHCARD
	}
}

func onePairComparables(h hand) []card {
	comparables := []card{}
	found := false
	var pairValue card

	for i := 0; i < h.Len() && !found; i++ {
		for j := i; j < h.Len(); j++ {
			if i != j && h[i].value == h[j].value {
				found = true
				pairValue = h[i]
				break
			}
		}
	}

	comparables = append(comparables, pairValue)

	for _, value := range h {
		if pairValue.value != value.value {
			comparables = append(comparables, value)
		}
	}

	return comparables
}

func onePair(first, second hand) result {
	return highCard(onePairComparables(first), onePairComparables(second), 0)
}

func twoPairsComparables(h hand) []card {
	comparables := []card{}

	if h[0] == h[1] {
		comparables = append(comparables, h[1])

		if h[2] == h[3] {
			comparables = append(comparables, h[2])
			comparables = append(comparables, h[4])
		} else {
			comparables = append(comparables, h[4])
			comparables = append(comparables, h[2])
		}

	} else {
		comparables = append(comparables, h[1])
		comparables = append(comparables, h[4])
		comparables = append(comparables, h[0])
	}

	return comparables
}

func twoPairs(first, second hand) result {
	return highCard(twoPairsComparables(first), twoPairsComparables(second), 0)
}

func highCard(first, second []card, index int) result {
	for i := index; i < len(first); i++ {
		if first[i].value > second[i].value {
			return winnerFirst
		}

		if first[i].value < second[i].value {
			return winnerSecond
		}
	}

	return draw
}

func threeComparables(h hand) []card {
	comparables := []card{}

	comparables = append(comparables, h[2])
	for _, value := range h {
		if value.value != h[2].value {
			comparables = append(comparables, value)
		}
	}

	return comparables
}

func three(first, second hand) result {
	return highCard(threeComparables(first), threeComparables(second), 0)
}

func fourComparables(h hand) []card {
	comparables := []card{}

	if h[0].value == h[1].value {
		comparables = append(comparables, h[0])
		comparables = append(comparables, h[4])
	} else {
		comparables = append(comparables, h[4])
		comparables = append(comparables, h[0])
	}

	return comparables
}

func four(first, second hand) result {
	return highCard(fourComparables(first), fourComparables(second), 0)
}

func fullHouseComparables(h hand) []card {
	comparables := []card{}

	if h[1].value == h[2].value {
		comparables = append(comparables, h[0])
		comparables = append(comparables, h[4])
	} else {
		comparables = append(comparables, h[4])
		comparables = append(comparables, h[0])
	}

	return comparables
}

func fullHouse(first, second hand) result {
	return highCard(fullHouseComparables(first), fullHouseComparables(second), 0)
}

func main() {
	games, err := readGameFile("input.txt")

	if err != nil {
		panic(err)
	}

	wins := 0

	for _, g := range games {
		if g.findResult() == winnerFirst {
			wins++
		}
	}

	fmt.Println(wins)
}
