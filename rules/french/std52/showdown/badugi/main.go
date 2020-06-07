package badugi

import (
	"github.com/luismasuelli/poker/assets/cards"
	"github.com/luismasuelli/poker/assets/cards/french"
)

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

const fullCardMask = 0x1111111111111111100

// A Badugi evaluator tries to return the best hand
// with the LEAST possible power. "A234" completely
// off-suit is the best hand, while "KKKK" the worst
// hand. As an expected precondition, which will not
// be tested, the hand must have 4 french cards in a
// "standard 52" set.
func Power(hand []cards.Card, community []cards.Card) (best []cards.Card, power uint64) {
	var handMask = [4]uint32{}
	// encoding each card as (1, mask, index).
	for index, card := range hand {
		handMask[index] = 1<<19 | cardMasks[card.(french.Card)] | uint32(index&3)
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
	for i := 0; i < 3; i++ {
		count := 0
		for j := i + 1; j < 4; j++ {
			if handMask[i]&handMask[j]&fullCardMask != 0 {
				count++
			}
		}
		if count > 1 {
			handMask[i] = 0
		}
	}
	// Then check cards that match ONE other card, and also
	// remove them.
	for i := 0; i < 3; i++ {
		count := 0
		for j := i + 1; j < 4; j++ {
			if handMask[i]&handMask[j]&fullCardMask != 0 {
				count++
			}
		}
		if count > 0 {
			handMask[i] = 0
		}
	}
	// Now, compute the power of the hand.
	// TODO.
}
