package french

import "fmt"

var faces = (func() []string {
	result := make([]string, 256)
	result[0] = "??"
	for index := 1; index <= 15; index++ {
		result[index] = "!!"
	}
	suits := " chds "
	ranks := "  23456789TJQKA "
	for suit := 1; suit <= 4; suit++ {
		result[suit*16] = "!!"
		result[suit*16+1] = "!!"
		result[suit*16+15] = "!!"
		for rank := 2; rank <= 14; rank++ {
			result[suit*16+rank] = fmt.Sprintf("%c%c", ranks[rank], suits[suit])
		}
	}
	result[80] = "*w"
	for entry := 81; entry < 255; entry++ {
		result[entry] = "!!"
	}
	return result
})()

// A standard suit in a french deck.
// It considers a "Hidden" suit and
// also a "Wildcard" suit.
type Suit uint8

const (
	Hidden Suit = iota
	Clubs
	Hearts
	Diamonds
	Spades
	Wildcard
)

// Standard french cards will have suits:
// 0 (unknown card), 5 (wildcard), or the
// 1-4 range (clubs, hearts, diamonds and
// spaces). For the last cases, they will
// have values 2-10 (numbers), 11-14 (J,
// Q, K, A).
type Card uint8

// Returns the rank of this card, or 0 if
// it is a wildcard or unknown card. As a
// precondition, cards will have suits in
// 0-5 range, and for 1-4 they will have
// ranks in 2-14 (0, 1, and 15 will not
// be used).
func (card Card) Rank() uint8 {
	suit := Suit(card >> 4)
	if suit < 1 || suit >= 4 {
		return 0
	} else {
		return uint8(card) & 15
	}
}

// Returns the suit of this card.
func (card Card) Suit() Suit {
	return Suit(card >> 4)
}

// A standard card face will involve the
// standard notation {rank}{suit} with
// the ranks in 23456789TJQKA and the
// suits in chds. For the unknown case,
// 2 question marks (??) will be used.
// For the wilcard, "!w" will be used.
func (card Card) Face() string {
	return faces[int(card)]
}

// Makes a card given a suit and a rank.
// The rank is applied module-16. The
// suit is applied module-4 with an offset
// of 1.
func MakeCard(suit Suit, rank uint8) Card {
	suit -= 1
	suit %= 4
	suit += 1
	return Card(uint8(suit << 4) + (rank & 15))
}

// An unknown card. This card is not yet revealed
// and triggers an error if used on evaluators.
// Its only purpose is to be serialized to the
// client.
var Unknown = Card(0)

// A joker card. Standard french cards consider
// the joker as the only wildcard-typed card.
var Joker = Card(80)