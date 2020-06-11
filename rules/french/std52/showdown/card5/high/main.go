package high

import (
	"github.com/luismasuelli/poker-go/assets/cards"
	"github.com/luismasuelli/poker-go/rules/french/std52/showdown/common"
)

// Computes the power of a hand using the standard
// high metric. This means: the hand is converted
// to A-high ranks, suits are kept, and straights are
// also considered. The result value is the power of
// such hand under those conditions and then it is
// returned alongside a 0b11111 flag telling all the
// involved cards (in this case: just the hand cards)
// are needed.
func Power(hand []cards.Card, community []cards.Card) (best uint32, power uint64) {
	rankBits, suitBits := common.PickAll(hand, common.HighRanks)
	power = common.Std52HighPower(rankBits, suitBits != 0)
	best = 0b11111
	return
}
