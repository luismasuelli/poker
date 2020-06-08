package badugi

import (
	"github.com/luismasuelli/poker/assets/cards"
	"github.com/luismasuelli/poker/assets/cards/french"
)

// Masks of 1-hot rank and 1-hot suit.
var cardMasks = (func() [52]uint32 {
	result := [52]uint32{}
	index := 0
	for suit := 0; suit < 4; suit++ {
		suitMask := uint32(1 << (2 + suit))
		for rank := 0; rank < 12; rank++ {
			// Since this is a lowball metric, we consider
			// Aces the lowest rank, so 0..11 (i.e. 2..K)
			// will shift 1 << (r + 7).
			result[index] = uint32(1<<(7+rank) | suitMask)
			index++
		}
		// Ace is low. It will be shifted by suit and index only.
		result[index] = 1<<6 | suitMask
		index++
	}
	return result
})()

// Full mask telling the space the 1-hot bits are relevant.
const fullMask uint32 = 0b1111111111111111100
const rankMask uint32 = 0b1111111111111000000
const indxMask uint32 = 0b0000000000000000011

// A Badugi evaluator tries to return the best hand
// with the LEAST possible power. "A234" completely
// off-suit is the best hand, while "KKKK" the worst
// hand. As an expected precondition, which will not
// be tested, the hand must have 4 french cards in a
// "standard 52" set.
//
// The return value is: (best, power) where
func Power(hand []cards.Card, community []cards.Card) (best uint32, power uint64) {
	handMask := [4]uint32{}
	removedCards := 0
	// encoding each card as (1, mask, index).
	for index, card := range hand[0:4] {
		handMask[index] = cardMasks[card.(french.Card)] | uint32(index)
	}
	// a manual sort over these uint32 values.
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if handMask[i] < handMask[j] {
				handMask[i], handMask[j] = handMask[j], handMask[i]
			}
		}
	}
	// First check cards that match TWO other cards or more,
	// and remove them.
	for i := 0; i < 4; i++ {
		count := 0
		for j := 0; j < 4; j++ {
			if j != i && handMask[i]&handMask[j]&fullMask != 0 {
				count++
			}
		}
		if count > 1 {
			handMask[i] = 0
			removedCards++
		}
	}
	// Then check cards that match ONE other card, and also
	// remove them.
	for i := 0; i < 3; i++ {
		count := 0
		for j := i + 1; j < 4; j++ {
			if handMask[i]&handMask[j]&fullMask != 0 {
				count++
			}
		}
		if count > 0 {
			handMask[i] = 0
			removedCards++
		}
	}
	// Now, compute the power of the hand. The power will be
	// comparable to other hands (the greater, the worst) but
	// will not serve to tell the distribution of win/lose
	// hands with respect to it.
	// Also, the indices will be computed as well. The indices
	// are an uint32 flagset telling which indices (marked as
	// set bits) make the best hand.
	best = 0
	power = uint64(removedCards << 19)
	for _, card := range handMask {
		power |= uint64(card & rankMask)
		if card != 0 {
			best |= 1 << (card & indxMask)
		}
	}
	return
}
