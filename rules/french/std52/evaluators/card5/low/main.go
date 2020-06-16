package low

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	"github.com/luismasuelli/poker-go/rules/french/std52/evaluators/common"
)

// Computes the power of a hand using the lowball
// metric. This means: the hand is converted to
// A-low ranks, suits are ignored, and straights
// are also ignored. The result value is the power
// of such hand under those conditions and then it
// is returned alongside a 0b11111 flag telling all
// the involved cards (in this case: just the hand
// cards) are needed.
func Power(hand []cards.Card, community []cards.Card) (best uint32, power uint64) {
	rankBits, _ := common.PickAll(hand, common.LowballRanks)
	power = common.Std52LowballPower(rankBits)
	best = 0b11111
	return
}
