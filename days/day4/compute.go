package day4

import (
	"fmt"
	"strings"

	"github.com/naruto678/aoc-go/globals"
)

type Card struct {
	cardNum    string
	winningSet map[string]bool
	handSet    map[string]bool
	score      int
}

func (c Card) GetScore() int {
	score := 0
	for values := range c.handSet {
		if c.winningSet[values] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func (c Card) GetMatches() int {
	score := 0
	for values := range c.handSet {
		if c.winningSet[values] {
			score++
		}
	}
	return score
}

func (c Card) String() string {
	return fmt.Sprintf("Card(cardNum=%s, winningSet = %T, handSet = %T, score = %d)", c.cardNum, c.winningSet, c.handSet, c.score)
}

func NewCard(input string) *Card {
	cardAttrs := strings.Split(input, ":")
	cardName, values := cardAttrs[0], cardAttrs[1]
	id, found := strings.CutPrefix(cardName, "Card ")

	card := &Card{
		cardNum:    id,
		winningSet: map[string]bool{},
		handSet:    map[string]bool{},
	}

	if !found {
		panic("card must have a card id")
	}
	sets := strings.Split(values, "|")
	winningSet, handSet := strings.TrimSpace(sets[0]), strings.TrimSpace(sets[1])

	for _, win := range strings.Fields(winningSet) {
		card.winningSet[win] = true
	}

	for _, hand := range strings.Fields(handSet) {
		card.handSet[hand] = true
	}

	return card
}

func computeFirst(content []byte) {
	strContent := string(content)
	cards := []*Card{}
	for _, line := range strings.Split(strContent, "\n") {
		cards = append(cards, NewCard(line))
	}
	totalSum := 0
	for _, card := range cards {
		totalSum += card.GetScore()
	}
	fmt.Println(fmt.Sprintf("computed the first part %d", totalSum))
}
func computeSecond(content []byte) {
	strContent := string(content)
	cards := []*Card{}
	for _, line := range strings.Split(strContent, "\n") {
		cards = append(cards, NewCard(line))
	}

	panic("not computed yet ")
}

func init() {
	globals.FuncMap["4-1"] = computeFirst
	globals.FuncMap["4-2"] = computeSecond
}
