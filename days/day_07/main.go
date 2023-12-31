package day_07

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// hand struct representing a set of cards
type hand struct {
	cards [5]uint8
	bid   uint16
}

// getType calculates the strength of the given hand
func (h hand) getType() uint8 {
	count := map[uint8]uint8{}
	for _, card := range h.cards {
		count[card]++
	}
	var counts []uint8
	for _, c := range count {
		counts = append(counts, c)
	}
	switch len(counts) {
	case 1:
		return 6 // five of a kind
	case 2:
		if slices.Max(counts)-slices.Min(counts) > 1 {
			return 5 // four of a kind
		}
		return 4 // full house
	case 3:
		if slices.Max(counts) == 3 {
			return 3 // three of a kind
		}
		return 2 // two pair
	case 4:
		return 1 // one pair
	}
	return 0 // high card
}

// getTypeWithJoker calculates the strength of the given hand while using jokers to make the hand stronger
func (h hand) getTypeWithJoker() uint8 {
	strongestHand := h.getType()
	for i, card := range h.cards {
		if card == 'J' {
			for _, otherCard := range h.cards {
				if otherCard != 'J' {
					otherHand := hand{
						cards: h.cards,
						bid:   h.bid,
					}
					otherHand.cards[i] = otherCard
					strongestHand = max(strongestHand, otherHand.getTypeWithJoker())
				}
			}
		}
	}

	return strongestHand
}

// isStronger compares the current hand with a given one
func (h hand) isStronger(other *hand, order [13]uint8, joker bool) bool {
	var handType uint8
	var otherType uint8

	if joker {
		handType = h.getTypeWithJoker()
		otherType = other.getTypeWithJoker()
	} else {
		handType = h.getType()
		otherType = other.getType()
	}

	if handType != otherType {
		return handType > otherType
	}

	for i := range h.cards {
		handIndex := indexOfCard(h.cards[i], order)
		otherIndex := indexOfCard(other.cards[i], order)

		if handIndex != otherIndex {
			return handIndex < otherIndex
		}
	}

	panic("identical cards")
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	var cardOrder = [13]uint8{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	hands := readHands(input)
	return strconv.Itoa(calculateTotalWinnings(hands, cardOrder, false))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	var cardOrder = [13]uint8{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	hands := readHands(input)
	return strconv.Itoa(calculateTotalWinnings(hands, cardOrder, true))
}

// readHands reads the input and converts the rows to hands
func readHands(input []string) []hand {
	var hands []hand
	for _, s := range input {
		split := strings.Split(s, " ")
		c := split[0]
		b := utils.Atoi(split[1])
		hands = append(hands, hand{
			cards: cardsToArray(c),
			bid:   uint16(b),
		})
	}
	return hands
}

// cardsToArray converts a string of cards into an array of ints
func cardsToArray(cards string) [5]uint8 {
	arr := [5]uint8{}
	for i, c := range cards {
		arr[i] = uint8(c)
	}
	return arr
}

// calculateTotalWinnings ranks the hands according to their type and multiplies the rank with the bidding value
func calculateTotalWinnings(hands []hand, order [13]uint8, joker bool) int {
	winnings := 0
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].isStronger(&hands[j], order, joker)
	})
	for i, h := range hands {
		rank := len(hands) - i
		winnings += rank * int(h.bid)
	}
	return winnings
}

// indexOfCard gets the index of the given card in the cardOrder slice
func indexOfCard(c uint8, cardOrder [13]uint8) uint8 {
	for i, u := range cardOrder {
		if c == u {
			return uint8(i)
		}
	}
	panic("card not found")
}
