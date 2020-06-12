package low

import (
	"github.com/luismasuelli/poker-go/assets/cards"
	. "github.com/luismasuelli/poker-go/assets/cards/french"
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
	// testHandPower(t, 0b0000000000000000000000000000000000111101000, 0b11111, S7, S9, C6, H4, S8, D7, H7)
	// testHandPower(t, 0b0000000000000000000000000000000001010100101, 0b1011011, HA, D6, H6, D8, D3, H8, HT)
	// testHandPower(t, 0b0000000000000000000000000000000001001001101, 0b1111001, DT, HT, ST, CA, S3, H4, D7)
}
