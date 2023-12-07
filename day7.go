package main

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type HandRank int

func (handRank HandRank) String() string {
	return [...]string{
		"HIGH_CARD",
		"ONE_PAIR",
		"TWO_PAIR",
		"THREE_OF_A_KIND",
		"FULL_HOUSE",
		"FOUR_OF_A_KIND",
		"FIVE_OF_A_KIND",
	}[handRank]
}

type Hand struct {
	cards       map[int]int
	sortedRanks []int
	bid         int
	asInput     string
}

func (h Hand) Format(f fmt.State, c rune) {
	f.Write([]byte(
		h.GetHandRank().String() +
			" " +
			h.asInput))
}

func (h *Hand) GetHandRank() HandRank {
	if h.cards[h.sortedRanks[0]] == 5 {
		return FIVE_OF_A_KIND
	} else if h.cards[h.sortedRanks[0]] == 4 {
		return FOUR_OF_A_KIND
	} else if h.cards[h.sortedRanks[0]] == 3 &&
		h.cards[h.sortedRanks[1]] == 2 {
		return FULL_HOUSE
	} else if h.cards[h.sortedRanks[0]] == 3 {
		return THREE_OF_A_KIND
	} else if h.cards[h.sortedRanks[0]] == 2 &&
		h.cards[h.sortedRanks[1]] == 2 {
		return TWO_PAIR
	} else if h.cards[h.sortedRanks[0]] == 2 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

type ByHandRank []Hand

func (h ByHandRank) Len() int      { return len(h) }
func (h ByHandRank) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByHandRank) Less(i, j int) bool {
	if h[i].GetHandRank() == h[j].GetHandRank() {
		for c := 0; c < 5; c++ {
			if h[i].asInput[c] == h[j].asInput[c] {
				continue
			}
			return ConvertSymbolToValue(h[i].asInput[c]) >
				ConvertSymbolToValue(h[j].asInput[c])
		}
	}
	return h[i].GetHandRank() > h[j].GetHandRank()
}

func RunDay7Part1() {
	file, err := os.Open("./inputs/day7input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := []Hand{}
	for scanner.Scan() {
		lineText := scanner.Text()

		bid := helpers.ParseInts(lineText[6:])[0]
		cards := make(map[int]int)
		jokersToHandle := 0
		for i := 0; i < 5; i++ {
			value := ConvertSymbolToValue(lineText[i])
			if value > 1 {
				cards[value]++
			} else {
				jokersToHandle++
			}
		}

		sortedRanks := make([]int, 0, len(cards))
		for card := range cards {
			sortedRanks = append(sortedRanks, card)
		}
		sort.SliceStable(sortedRanks, func(i, j int) bool {
			return cards[sortedRanks[i]] > cards[sortedRanks[j]]
		})

		if jokersToHandle == 5 {
			cards[14] = 5
			sortedRanks = make([]int, 0, len(cards))
			for card := range cards {
				sortedRanks = append(sortedRanks, card)
			}
			sort.SliceStable(sortedRanks, func(i, j int) bool {
				return cards[sortedRanks[i]] > cards[sortedRanks[j]]
			})
		} else {
			cards[sortedRanks[0]] += jokersToHandle
		}

		hands = append(hands, Hand{
			cards:       cards,
			sortedRanks: sortedRanks,
			bid:         bid,
			asInput:     lineText[0:5]})
	}

	sort.Sort(ByHandRank(hands))

	totalWinnings := 0
	for i := 0; i < len(hands); i++ {
		totalWinnings += (len(hands) - i) * hands[i].bid
	}
	fmt.Println(totalWinnings)
}

func ConvertSymbolToValue(symbol byte) int {
	switch symbol {
	case 65: // A
		return 14
	case 75: // K
		return 13
	case 81: // Q
		return 12
	case 74: // J
		return 1
	case 84: // T
		return 10
	default:
		return int(symbol) - 48
	}
}
