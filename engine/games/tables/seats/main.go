package seats

import (
	"errors"
	"github.com/luismasuelli/poker-go/engine/games/cards"
	"github.com/luismasuelli/poker-go/engine/players"
)

// Status of a seat during a game.
type Status uint32

const (
	// No player is occupying this seat.
	Free Status = iota
	// The player is waiting the next hand
	// in which they "will chose to play".
	Waiting
	// The player is playing the hand, and
	// has $X > 0.
	Active
	// The player is playing the hand, but
	// it has $0, which marks them as being
	// "all-in".
	AllIn
	// The player folded.
	Folded
)

// Additional flags to use (e.g. "sit out")
// to complement a status.
type Flags uint32

const (
	// No flags at all.
	Nothing Flags = 0
	// The seat is marked as "sit out".
	SitOut Flags = 1
)

// Interfaces for a seat. There are
// several different rules for the seats,
// depending mostly on the game rule (e.g.
// games with blinds, or not).
type Seat interface {
	// Data to be retrieved.
	// The player.
	Player() players.Player
	// The stack.
	Stack() uint64
	// The status.
	Status() Status
	// The flags.
	Flags() Flags

	// Sits a player, if empty.
	Sit(player players.Player, stack uint64) error
	// Pops a player, if occupied.
	Pop() (player players.Player, stack uint64)
	// Sets the status, if occupied.
	// It is not allowed to set the Free state.
	SetStatus(status Status) error
	// Sets a flag.
	SetFlag(flag Flags) error
	// Clears a flag.
	ClearFlag(flag Flags) error
	// Takes chips, if occupied and affordable.
	SubStack(chips uint64) error
	// Gives chips, if occupied and not overflowing.
	AddStack(chips uint64) error
	// Sets stack, if occupied.
	SetStack(chips uint64) error
	// Takes chips from the pot, if occupied and affordable.
	SubPot(chips uint64) error
	// Gives chips to the pot, if occupied and not overflowing.
	AddPot(chips uint64) error
	// Sets pot, if occupied.
	SetPot(chips uint64) error
	// Adds cards to the hand.
	AddCards(seatCards []*SeatCard) error
	// Takes cards (by indices) from the hand.
	RemoveCards(indices []int) error
	// Gets the actual cards.
	Cards(revealed bool) []cards.Card
}

// A card in a seat. Players will have
// cards (the number varies with the
// game), which at different points or
// according to different rules, will
// be covered or shown to the table.
type SeatCard struct {
	card  cards.Card
	shown bool
}

// Tells whether the card is shown or
// not.
func (seatCard *SeatCard) Shown() bool {
	return seatCard.shown
}

// Tells which card to show. If the card
// is already shown, return it. Also if
// the card is told to be revealed, then
// also return it. Otherwise, return nil.
func (seatCard *SeatCard) Card(revealed bool) cards.Card {
	if revealed || seatCard.Shown() {
		return seatCard.card
	} else {
		return nil
	}
}

// Creates a new seat card (hidden).
func NewSeatCard(card cards.Card) *SeatCard {
	return &SeatCard{card, false}
}

// Creates a new seat card (shown).
func NewSeatShownCard(card cards.Card) *SeatCard {
	return &SeatCard{card, true}
}

// Seats are managed by the tables. They
// keep a reference to the player and to
// each aspect of their participation in
// the table. This class could and should
// be enabled to support more flags that
// may be related to subsets of the rules,
// e.g. whether there are blinds being
// used, and even the underlying data
// related to the algorithm for blinds
// management (simple, moving, or dead).
type BaseSeat struct {
	player players.Player
	status Status
	flags  Flags
	stack  uint64
	pot    uint64
	cards  []*SeatCard
}

// Gets the underlying sit player.
func (seat *BaseSeat) Player() players.Player {
	return seat.player
}

// Gets the stack of this seat.
func (seat *BaseSeat) Stack() uint64 {
	return seat.stack
}

// Gets the current pot of this seat,
// for when bets are being made.
func (seat *BaseSeat) Pot() uint64 {
	return seat.pot
}

// Returns the status of a seat. There are
// only 5 statuses in poker games, although
// flags may condition when and how these
// states are changed.
func (seat *BaseSeat) Status() Status {
	return seat.status
}

// Returns the flags of this seat. So far,
// this only involves the "sit out" button.
func (seat *BaseSeat) Flags() Flags {
	return seat.flags
}

var ErrCannotSitNilPlayer = errors.New("cannot sit a nil player")
var ErrCannotSitPlayerOnOccupiedSeat = errors.New("cannot sit a player on an occupied seat")
var ErrCannotSitPlayerWithEmptyStack = errors.New("cannot sit a player with an empty stack")

// Tries to sit a player, if given and with
// a non-empty stack, and only if no player
// is already sit. It is recommended for
// this method to be redefined on composition
// if more stuff is to be changed on sit.
func (seat *BaseSeat) Sit(player players.Player, stack uint64) error {
	if seat.status != Free {
		return ErrCannotSitPlayerOnOccupiedSeat
	} else if player == nil {
		return ErrCannotSitNilPlayer
	} else if stack == 0 {
		return ErrCannotSitPlayerWithEmptyStack
	} else {
		seat.player = player
		seat.stack = stack
		seat.status = Waiting
		return nil
	}
}

// Pops a player and stack from the seat. It
// returns (nil, 0) if the seat was empty.
// It is recommended for this method to be
// redefined on composition if more stuff is
// to be changed on pop.
func (seat *BaseSeat) Pop() (player players.Player, stack uint64) {
	player = seat.player
	stack = seat.stack
	seat.player = nil
	seat.stack = 0
	seat.pot = 0
	seat.status = Free
	seat.flags = Nothing
	seat.cards = make([]*SeatCard, 0)
	return
}

var ErrCannotSetFreeStatus = errors.New("cannot manually set the 'free' status in a seat")
var ErrCannotChangeStatusOnEmptySeat = errors.New("cannot manually change status on empty seats")

// Sets the status of this seat. The "free"
// status cannot set, and cannot replaced,
// by this method.
func (seat *BaseSeat) SetStatus(status Status) error {
	if seat.player == nil {
		return ErrCannotChangeStatusOnEmptySeat
	} else if status == Free {
		return ErrCannotSetFreeStatus
	}
	seat.status = status
	return nil
}

var ErrCannotChangeFlagsOnEmptySeat = errors.New("cannot manually change flags on empty seats")

// Sets a flag on this seat. This can only
// be done to non-free seats.
func (seat *BaseSeat) SetFlag(flag Flags) error {
	if seat.player == nil {
		return ErrCannotChangeFlagsOnEmptySeat
	} else {
		seat.flags |= flag
		return nil
	}
}

// Clears a flag on this seat. This can only
// be done to non-free seats.
func (seat *BaseSeat) ClearFlag(flag Flags) error {
	if seat.player == nil {
		return ErrCannotChangeFlagsOnEmptySeat
	} else {
		seat.flags &= ^flag
		return nil
	}
}

var ErrCannotChangeStackOnEmptySeat = errors.New("cannot manually change stack on empty seats")
var ErrCannotAffordChipsSubtraction = errors.New("cannot afford the specified chips subtraction")
var ErrCannotContainChipsAddition = errors.New("cannot contain the specified chips addition")

// Takes chips from the seat. This can only
// be done to non-free seats that can afford
// the specified chips amount.
func (seat *BaseSeat) SubStack(chips uint64) error {
	if seat.player == nil {
		return ErrCannotChangeStackOnEmptySeat
	} else if seat.stack < chips {
		return ErrCannotAffordChipsSubtraction
	} else {
		seat.stack -= chips
		return nil
	}
}

// Gives chips to the seat. This can only be
// done to non-free seats that can receive
// the amount without overflowing. It is
// recommended that table levels and user
// accounts are defensively restricted to
// reach a case like this.
func (seat *BaseSeat) AddStack(chips uint64) error {
	if seat.player == nil {
		return ErrCannotChangeStackOnEmptySeat
	} else if seat.stack > (^(uint64(0) - chips)) {
		return ErrCannotContainChipsAddition
	} else {
		seat.stack += chips
		return nil
	}
}

// Sets the chips to the seat. This can only
// be done to non-free seats, and care should
// be taken when calling this method, since it
// may cause overflow as well when receiving
// more money or popping the seat in a cash
// 1-table game.
func (seat *BaseSeat) SetStack(chips uint64) error {
	if seat.player == nil {
		return ErrCannotChangeStackOnEmptySeat
	} else {
		seat.stack = chips
		return nil
	}
}

var ErrCannotChangePotOnEmptySeat = errors.New("cannot manually change pot on empty seats")
var ErrCannotAffordPotChipsSubtraction = errors.New("cannot afford the specified pot chips subtraction")
var ErrCannotContainPotChipsAddition = errors.New("cannot contain the specified pot chips addition")

// Takes chips from the seat's pot. This can
// only be done to non-free seats that can
// afford the specified chips amount from their
// pots.
func (seat *BaseSeat) SubPot(chips uint64) error {
	if seat.player == nil {
		return ErrCannotChangePotOnEmptySeat
	} else if seat.pot < chips {
		return ErrCannotAffordPotChipsSubtraction
	} else {
		seat.pot -= chips
		return nil
	}
}

// Gives chips to the seat's pot. This can
// only be done to non-free seats that can
// receive the amount without overflowing.
// It is recommended that table levels and
// user accounts are defensively restricted
// to reach a case like this.
func (seat *BaseSeat) AddPot(chips uint64) error {
	if seat.player == nil {
		return ErrCannotChangePotOnEmptySeat
	} else if seat.pot > (^(uint64(0) - chips)) {
		return ErrCannotContainPotChipsAddition
	} else {
		seat.pot += chips
		return nil
	}
}

// Sets the pot to the seat. This can only
// be done to non-free seats, and care should
// be taken when calling this method, since it
// may cause overflow as well when receiving
// more money or popping the seat in a cash
// 1-table game.
func (seat *BaseSeat) SetPot(chips uint64) error {
	if seat.player == nil {
		return ErrCannotChangePotOnEmptySeat
	} else {
		seat.pot = chips
		return nil
	}
}

var ErrCannotAddCardsToEmptySeat = errors.New("cannot add cards to empty seats")
var ErrCannotAddNilOrEmptyCardToSeat = errors.New("cannot add a nil or empty card to a seat")
var ErrCannotRemoveCardsFromEmptySeat = errors.New("cannot remove cards from empty seats")

// Adds non-nil cards.
func (seat *BaseSeat) AddCards(seatCards []*SeatCard) error {
	if seat.player == nil {
		return ErrCannotAddCardsToEmptySeat
	} else if len(seatCards) == 0 {
		return nil
	} else {
		for _, seatCard := range seatCards {
			if seatCard == nil || seatCard.Card(true) == nil {
				return ErrCannotAddNilOrEmptyCardToSeat
			}
		}

		seat.cards = append(seat.cards, seatCards...)
		return nil
	}
}

// Removes cards by indices. A []int{-1} array
// clears the hand.
func (seat *BaseSeat) RemoveCards(indices []int) error {
	if seat.player == nil {
		return ErrCannotAddCardsToEmptySeat
	} else if len(indices) == 1 && indices[0] < 0 {
		// Test []int{-1} special case.
		seat.cards = make([]*SeatCard, 0)
	} else {
		cardsLen := len(seat.cards)
		indicesSet := map[int]bool{}
		for _, index := range indices {
			if index < cardsLen && index >= 0 {
				indicesSet[index] = true
			}
		}

		newCards := make([]*SeatCard, cardsLen-len(indicesSet))
		newIndex := 0
		for index, card := range seat.cards {
			if _, ok := indicesSet[index]; !ok {
				newCards[newIndex] = card
				newIndex++
			}
		}
		seat.cards = newCards
	}
	return nil
}

// Gets the actual cards from the seat.
func (seat *BaseSeat) Cards(revealed bool) []cards.Card {
	cards := make([]cards.Card, len(seat.cards))
	for index, seatCard := range seat.cards {
		cards[index] = seatCard.Card(revealed)
	}
	return cards
}
