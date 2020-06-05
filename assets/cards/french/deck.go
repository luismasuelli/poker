package french

import "github.com/luismasuelli/poker/assets/cards"

// A standard french deck contains a sequence of standard french
// cards, and will act on it, change it, and so.
type Deck struct {
	cards []Card
}

// This is just a helper function to create a custom deck given:
// - epochs: How many times will this algorithm run to add the
//           {ranks}x{suits} cards, and the wildcards.
// - ranks: The ranks to add. Valid values are 2 .. 14. Other
//          values have a high chance of generating invalid cards.
// - suits: The suits to add. Valid values are the given constants
//          Clubs, Hearts, Diamonds and Spades.
// - wildcards: The number of wildcards to add to the deck.
// - extra: Additional cards to add to the deck. These cards will
//          be added only once per argument (the same card may be
//          specified many times, and thus be added many times),
//          out of the epochs.
func CustomDeck(epochs uint8, ranks []uint8, suits []Suit, wildcards uint8, extras ...Card) Deck {
	newCards := make([]Card, int(epochs) * (len(ranks) * len(suits) + int(wildcards)) + len(extras))
	index := 0
	for epoch := uint8(0); epoch < epochs; epoch++ {
		for _, suit := range suits {
			for _, rank := range ranks {
				newCards[index] = MakeCard(suit, rank)
				index++
			}
		}
		for wildcard := uint8(0); wildcard < wildcards; wildcard++ {
			newCards[index] = Joker
			index++
		}
	}
	for _, extra := range extras {
		newCards[index] = extra
		index++
	}
	return Deck{cards: newCards}
}

// The length of a deck is the length of the underlying array.
func (deck *Deck) Len() int {
	return len(deck.cards)
}

// Returns a copy of the current deck.
func (deck *Deck) Copy() cards.Deck {
	currentCards := make([]Card, len(deck.cards))
	copy(currentCards, deck.cards)
	return &Deck{cards: currentCards}
}

// Swaps two cards inside the deck.
func (deck *Deck) Swap(i, j int) {
	deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
}

// Deals n cards from the top of the deck.
func (deck *Deck) Deal(n int) []cards.Card {
	peeked := deck.Peek(n)
	deck.cards = deck.cards[0:len(deck.cards) - len(peeked)]
	return peeked
}

// Peeks n cards from the top of the deck (a non-destructive
// way of dealing cards).
func (deck *Deck) Peek(n int) []cards.Card {
	baseLength := len(deck.cards)

	if n < 1 {
		panic(cards.ErrDealBadCount)
	} else if n > baseLength {
		panic(cards.ErrDealNotEnough)
	} else {
		newLength := baseLength - n
		source := deck.cards[newLength:baseLength]
		result := make([]cards.Card, len(source))
		for index := 0; index < n; index++ {
			result[index] = source[n - 1 - index]
		}
		return result
	}
}

// Stacks new cards onto the deck, in the order
// they are specified.
func (deck *Deck) Stack(cards []cards.Card) {
	newLength := len(cards)
	baseLength := len(deck.cards)

	if newLength == 0 {
		return
	}

	newCards := make([]Card, newLength + baseLength)
	copy(newCards, deck.cards)
	for index := 0; index < newLength; index++ {
		newCards[baseLength + index] = cards[index].(Card)
	}

	deck.cards = newCards
}

// Queues new cards under the deck, in the order
// they are specified.
func (deck *Deck) Queue(cards []cards.Card) {
	newLength := len(cards)
	baseLength := len(deck.cards)

	if newLength == 0 {
		return
	}

	newCards := make([]Card, newLength + baseLength)
	for index := 0; index < newLength; index++ {
		newCards[newLength - 1 - index] = cards[index].(Card)
	}
	copy(newCards[newLength:newLength + baseLength], deck.cards)

	deck.cards = newCards
}

// A standard deck, used in most poker games.
var StandardDeck = CustomDeck(1, []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, []Suit{Clubs, Hearts, Diamonds, Spades}, 0)

// A deck with standard cards and one wildcard.
var Wilcard1Deck = CustomDeck(1, []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, []Suit{Clubs, Hearts, Diamonds, Spades}, 1)

// A deck with standard cards and two wildcard.
var Wilcard2Deck = CustomDeck(1, []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, []Suit{Clubs, Hearts, Diamonds, Spades}, 2)

// A 13-cards deck consisting of clubs-suited cards only.
var ClubsDeck = CustomDeck(1, []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, []Suit{Clubs}, 0)

// A 13-cards deck consisting of hearts-suited cards only.
var HeartsDeck = CustomDeck(1, []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, []Suit{Hearts}, 0)

// A 13-cards deck consisting of diamonds-suited cards only.
var DiamondsDeck = CustomDeck(1, []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, []Suit{Diamonds}, 0)

// A 13-cards deck consisting of spades-suited cards only.
var SpadesDeck = CustomDeck(1, []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, []Suit{Spades}, 0)

// A Kuhn Poker's deck
var KuhnDeck = CustomDeck(1, []uint8{11, 12, 13}, []Suit{Spades}, 0)