package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type primaryRank int
type secondaryRank []card

func compareSecondaryRanks(r1, r2 secondaryRank, index int) result {
	for i := index; i < len(r1); i++ {
		if r1[i].value > r2[i].value {
			return winnerFirst
		}

		if r1[i].value < r2[i].value {
			return winnerSecond
		}
	}

	return draw
}

type rank struct {
	primary   primaryRank
	secondary secondaryRank
}

func (hand *rank) getSecondaryRank() []card {
	return hand.secondary
}

func (hand *rank) getPrimaryRank() primaryRank {
	return hand.primary
}

func comparison(hand1, hand2 rank) result {
	if hand1.getPrimaryRank() > hand2.getPrimaryRank() {
		return winnerFirst
	}

	if hand1.getPrimaryRank() < hand2.getPrimaryRank() {
		return winnerSecond
	}

	return compareSecondaryRanks(hand1.getSecondaryRank(), hand2.getSecondaryRank(), 0)
}

const (
	HIGHCARD primaryRank = iota
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
	rankedFirst := evaluateHand(g.firstHand)
	rankedSecond := evaluateHand(g.secondHand)

	return comparison(rankedFirst, rankedSecond)
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

	firstHand := hand(cards[0:5])
	sort.Sort(firstHand)

	secondHand := hand(cards[5:10])
	sort.Sort(secondHand)

	return game{firstHand, secondHand}
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

func findPair(h hand) (int, []card) {
	pair := 0
	var pairCard card
	var comparables []card

	for i := 0; i < h.Len(); i++ {
		for j := i; j < h.Len(); j++ {
			if i != j && h[i].value == h[j].value {
				pair++
				pairCard = h[i]
			}
		}
	}

	if pair == 0 {
		return 0, nil
	}

	comparables = append(comparables, pairCard)

	for i := 0; i < h.Len(); i++ {
		if h[i].value != pairCard.value {
			comparables = append(comparables, h[i])
		}
	}

	return pair, comparables
}

func (h hand) isRoyalFlush() (bool, []card) {
	return sameSuit(h) && h[0].value == 14 && h[1].value == 13 &&
		h[2].value == 12 && h[3].value == 11 && h[4].value == 10, nil
}

func (h hand) isStraightFlush() (bool, []card) {
	if sameSuit(h) && h[0].value == h[1].value+1 && h[1].value == h[2].value+1 &&
		h[2].value == h[3].value+1 && h[3].value == h[4].value+1 {
		return true, h
	}

	return false, nil
}

func (h hand) isFour() (bool, []card) {
	if h[1].value == h[2].value && h[2].value == h[3].value {
		if h[0].value == h[1].value {
			return true, []card{h[0], h[4]}
		} else if h[3].value == h[4].value {
			return true, []card{h[4], h[0]}
		}
	}

	return false, nil
}

func (h hand) isFullHouse() (bool, []card) {
	if h[0].value == h[1].value && h[3].value == h[4].value {
		if h[2].value == h[1].value {
			return true, []card{h[0], h[4]}
		} else if h[2].value == h[3].value {
			return true, []card{h[4], h[0]}
		}
	}

	return false, nil
}

func (h hand) isFlush() (bool, []card) {
	if sameSuit(h) {
		return true, h
	}

	return false, nil
}

func (h hand) isStraight() (bool, []card) {
	if h[0].value == h[1].value+1 && h[1].value == h[2].value+1 &&
		h[2].value == h[3].value+1 && h[3].value == h[4].value+1 {
		return true, h
	}

	return false, nil
}

func (h hand) isThree() (bool, []card) {
	if h[0].value == h[1].value && h[1].value == h[2].value {
		return true, []card{h[0], h[3], h[4]}
	}

	if h[1].value == h[2].value && h[2].value == h[3].value {
		return true, []card{h[1], h[0], h[4]}
	}

	if h[2].value == h[3].value && h[3].value == h[4].value {
		return true, []card{h[2], h[0], h[1]}
	}

	return false, nil
}

func (h hand) isTwoPairs() (bool, []card) {
	pairs, comparables := findPair(h)

	if pairs == 2 {
		return true, comparables
	}

	return false, nil
}

func (h hand) isOnePair() (bool, []card) {
	pairs, comparables := findPair(h)

	if pairs == 1 {
		return true, comparables
	}

	return false, nil
}

func evaluateHand(h hand) rank {
	if temp, comparables := h.isRoyalFlush(); temp {
		return rank{ROYALFLUSH, comparables}
	}

	if temp, comparables := h.isStraightFlush(); temp {
		return rank{STRAIGHTFLUSH, comparables}
	}

	if temp, comparables := h.isFour(); temp {
		return rank{FOUR, comparables}
	}

	if temp, comparables := h.isFullHouse(); temp {
		return rank{FULLHOUSE, comparables}
	}
	if temp, comparables := h.isFlush(); temp {
		return rank{FLUSH, comparables}
	}
	if temp, comparables := h.isStraight(); temp {
		return rank{STRAIGHT, comparables}
	}
	if temp, comparables := h.isThree(); temp {
		return rank{THREE, comparables}
	}
	if temp, comparables := h.isTwoPairs(); temp {
		return rank{TWOPAIRS, comparables}
	}
	if temp, comparables := h.isOnePair(); temp {
		return rank{ONEPAIR, comparables}
	}

	return rank{HIGHCARD, secondaryRank(h)}
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
