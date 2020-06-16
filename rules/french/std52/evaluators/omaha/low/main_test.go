package low

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	. "github.com/luismasuelli/poker-go/engine/games/cards/french"
	"testing"
)

func testHandPower(t *testing.T, expectedPower uint64, expectedBest uint32, cards ...cards.Card) {
	best, power := Power(cards, nil)
	if power != expectedPower || best != expectedBest {
		t.Errorf(
			"Testing hand: %v\nexpected power: %#064b\n     got power: %#064b\nexpected best: %#07b\n     got best: %#07b\n",
			cards, expectedPower, power, expectedBest, best,
		)
	}
}

func TestHandPowers(t *testing.T) {
	// 4 of a kind: are easily avoided in Omaha low.
	// Full house: are easily avoided in Omaha low.
	// 3 of a kind: are easily avoided in Omaha low.
	// Double Pair:
	testHandPower(t, 0b0010000000000000000010100000000100000000000, 0b111000011, CT, ST, D8, C8, HT, DT, H8, S8, CQ)
	testHandPower(t, 0b0010000000000000000000010000010001000000000, 0b001110101, DA, SA, HT, D7, CA, C7, H7, S7, HA)
	// Pair.
	testHandPower(t, 0b0001000000000000000000000000100000000011100, 0b001111001, C2, C3, C4, C5, H2, H3, H4, S3, S4)
	testHandPower(t, 0b0001000000000000000000000000100000010001100, 0b100110011, C2, H2, C4, H4, C8, S4, D4, S2, S3)
	// Bust / High Card.
	testHandPower(t, 0b0000000000000000000000000000000011100000011, 0b110100101, C9, CT, CJ, CK, HK, H2, H9, HA, HT)
	testHandPower(t, 0b0000000000000000000000000000000100110000101, 0b111000011, CA, C3, HA, H4, CK, HK, C9, C8, CQ)
	testHandPower(t, 0b0000000000000000000000000000000000110011001, 0b111000011, C9, C5, H5, H4, D5, D9, S8, S4, SA)
}
