package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	HIGHCARD = iota
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

type game struct {
	firstHand   hand
	secondHand  hand
	winnerFirst bool
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

func (g *game) findWinner() {
	first := g.firstHand
	sort.Sort(first)
	second := g.secondHand
	sort.Sort(second)

	rankFirst := evaluateHand(first)
	rankSecond := evaluateHand(second)

	if rankFirst > rankSecond {
		g.winnerFirst = true
		return
	}

	if rankFirst < rankSecond {
		g.winnerFirst = false
		return
	}

	switch rankFirst {
	case ROYALFLUSH:
		g.winnerFirst = false
	case STRAIGHTFLUSH:
		g.winnerFirst = highCard(first, second, 0)
	case FOUR:
		g.winnerFirst = four(first, second)
	case FULLHOUSE:
		g.winnerFirst = fullHouse(first, second)
	case FLUSH:
		g.winnerFirst = highCard(first, second, 0)
	case STRAIGHT:
		g.winnerFirst = highCard(first, second, 0)
	case THREE:
		g.winnerFirst = three(first, second)
	case TWOPAIRS:
		g.winnerFirst = twoPairs(first, second)
	case ONEPAIR:
		g.winnerFirst = onePair(first, second)
	default:
		g.winnerFirst = highCard(first, second, 0)
	}
}

func readGameFile(fileName string) ([]game, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("File opening failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make([]game, 0)

	for scanner.Scan() {
		g := readGame(scanner.Text())
		if err != nil {
			return nil, errors.New("Reading line failed")
		}

		games = append(games, g)
	}

	return games, nil
}

func readGame(line string) game {
	var cards []card

	for _, value := range strings.Split(line, " ") {

		cards = append(cards, readCard(value))
	}

	g := game{}

	g.firstHand = cards[0:5]
	g.secondHand = cards[5:10]

	return g
}

func readCard(s string) card {
	c := card{}

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

func sameSuit(a []bool, h hand) bool {
	c := []card{}

	for i, value := range a {
		if value {
			c = append(c, h[i])
		}
	}

	s := c[0].Suit

	for _, value := range c {
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

func evaluateHand(h hand) int {
	suits := []bool{true, true, true, true, true}

	if sameSuit(suits, h) && h[0].value == 14 && h[1].value == 13 && h[2].value == 12 && h[3].value == 11 && h[4].value == 10 {
		fmt.Println(h)
		fmt.Println("ROYALFLUSH")
		return ROYALFLUSH
	}

	if sameSuit(suits, h) && h[0].value == h[1].value+1 && h[1].value == h[2].value+1 && h[2].value == h[3].value+1 && h[3].value == h[4].value+1 {
		fmt.Println(h)
		fmt.Println("FLUSH")
		return FLUSH
	}

	if h[1].value == h[2].value && h[2].value == h[3].value && (h[0].value == h[1].value || h[3].value == h[4].value) {
		fmt.Println(h)
		fmt.Println("FOUR")
		return FOUR
	}

	if h[0].value == h[1].value && h[3].value == h[4].value && (h[2].value == h[1].value || h[2].value == h[3].value) {
		fmt.Println(h)
		fmt.Println("FULLHOUSE")
		return FULLHOUSE
	}

	if sameSuit(suits, h) {
		fmt.Println(h)
		fmt.Println("FLUSH")
		return FLUSH
	}

	if h[0].value == h[1].value+1 && h[1].value == h[2].value+1 && h[2].value == h[3].value+1 && h[3].value == h[4].value+1 {
		fmt.Println(h)
		fmt.Println("STRAIGHT")
		return STRAIGHT
	}

	if (h[0].value == h[1].value && h[1].value == h[2].value) || (h[1].value == h[2].value && h[2].value == h[3].value) || (h[2].value == h[3].value && h[3].value == h[4].value) {
		fmt.Println(h)
		fmt.Println("THREE")
		return THREE
	}

	if findPair(h) == 2 {
		fmt.Println(h)
		fmt.Println("TWOPAIRS")
		return TWOPAIRS
	}

	if findPair(h) == 1 {
		fmt.Println(h)
		fmt.Println("ONEPAIR")
		return ONEPAIR
	}

	fmt.Println(h)
	fmt.Println("HIGHCARD")
	return HIGHCARD
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

func onePair(first, second hand) bool {
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

func twoPairs(first, second hand) bool {
	return highCard(twoPairsComparables(first), twoPairsComparables(second), 0)
}

func highCard(first, second []card, index int) bool {
	for i := index; i < len(first); i++ {
		if first[i].value > second[i].value {
			return true
		}

		if first[i].value < second[i].value {
			return false
		}
	}

	return false
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

func three(first, second hand) bool {
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

func four(first, second hand) bool {
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

func fullHouse(first, second hand) bool {
	return highCard(fullHouseComparables(first), fullHouseComparables(second), 0)
}

func main() {
	games, err := readGameFile("input.txt")

	if err != nil {
		panic(err)
	}

	wins := 0

	for _, g := range games {
		g.findWinner()
		fmt.Println(g.winnerFirst)
		fmt.Println()
		if g.winnerFirst {
			wins++
		}
	}

	fmt.Println(wins)
}
