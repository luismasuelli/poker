package card7

import "github.com/luismasuelli/poker-go/assets/cards"

// All the available 7C5 combinations.
var Combinations = [][]uint64{
	{0b0011111, 1, 1, 1, 1, 1, 0, 0},
	{0b0101111, 1, 1, 1, 1, 0, 1, 0},
	{0b1001111, 1, 1, 1, 1, 0, 0, 1},
	{0b0110111, 1, 1, 1, 0, 1, 1, 0},
	{0b1010111, 1, 1, 1, 0, 1, 0, 1},
	{0b1100111, 1, 1, 1, 0, 0, 1, 1},
	{0b0111011, 1, 1, 0, 1, 1, 1, 0},
	{0b1011011, 1, 1, 0, 1, 1, 0, 1},
	{0b1101011, 1, 1, 0, 1, 0, 1, 1},
	{0b1110011, 1, 1, 0, 0, 1, 1, 1},
	{0b0111101, 1, 0, 1, 1, 1, 1, 0},
	{0b1011101, 1, 0, 1, 1, 1, 0, 1},
	{0b1101101, 1, 0, 1, 1, 0, 1, 1},
	{0b1110101, 1, 0, 1, 0, 1, 1, 1},
	{0b1111001, 1, 0, 0, 1, 1, 1, 1},
	{0b0111110, 0, 1, 1, 1, 1, 1, 0},
	{0b1011110, 0, 1, 1, 1, 1, 0, 1},
	{0b1101110, 0, 1, 1, 1, 0, 1, 1},
	{0b1110110, 0, 1, 1, 0, 1, 1, 1},
	{0b1111010, 0, 1, 0, 1, 1, 1, 1},
	{0b1111100, 0, 0, 1, 1, 1, 1, 1},
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
