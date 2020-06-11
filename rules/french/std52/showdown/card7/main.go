package card7

import "github.com/luismasuelli/poker-go/assets/cards"

// All the available 7C5 combinations.
var Combinations = [][]uint64{
	{1, 1, 1, 1, 1, 0, 0},
	{1, 1, 1, 1, 0, 1, 0},
	{1, 1, 1, 1, 0, 0, 1},
	{1, 1, 1, 0, 1, 1, 0},
	{1, 1, 1, 0, 1, 0, 1},
	{1, 1, 1, 0, 0, 1, 1},
	{1, 1, 0, 1, 1, 1, 0},
	{1, 1, 0, 1, 1, 0, 1},
	{1, 1, 0, 1, 0, 1, 1},
	{1, 1, 0, 0, 1, 1, 1},
	{1, 0, 1, 1, 1, 1, 0},
	{1, 0, 1, 1, 1, 0, 1},
	{1, 0, 1, 1, 0, 1, 1},
	{1, 0, 1, 0, 1, 1, 1},
	{1, 0, 0, 1, 1, 1, 1},
	{0, 1, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 0, 1},
	{0, 1, 1, 1, 0, 1, 1},
	{0, 1, 1, 0, 1, 1, 1},
	{0, 1, 0, 1, 1, 1, 1},
	{0, 0, 1, 1, 1, 1, 1},
}

// Combines the hand cards and community cards in a single
// array to be used by each player.
func AddCards(hand, community []cards.Card) []cards.Card {
	communityLen := len(community)
	if communityLen == 0 {
		return hand
	} else {
		handLen := len(hand)
		return append(append(make([]cards.Card, 0, handLen+communityLen), hand...), community...)
	}
}
