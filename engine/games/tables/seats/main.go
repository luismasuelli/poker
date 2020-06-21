package seats

import (
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
	Sit(player players.Player, stack uint64) bool
	// Pops a player, if occupied.
	Pop() (player players.Player, stack uint64)
	// Sets the status, if occupied.
	// It is not allowed to set the Free state.
	SetStatus(status Status) bool
	// Sets a flag.
	SetFlag(flag Flags) bool
	// Clears a flag.
	ClearFlag(flag Flags) bool
	// Takes chips, if occupied and affordable.
	TakeChips(chips uint64) bool
	// Gives chips, if occupied and not overflowing.
	GiveChips(chips uint64) bool
	// Sets stack, if occupied.
	SetStack(chips uint64) bool
	// Adds cards to the hand.
	AddCards(seatCards []*SeatCard) bool
	// Takes cards (by indices) from the hand.
	RemoveCards(indices []int)
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

// Marks the card as shown.
func (seatCard *SeatCard) Show() {
	seatCard.shown = true
}

// Marks the card as hidden.
func (seatCard *SeatCard) Hide() {
	seatCard.shown = false
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

// Tries to sit a player, if given and with
// a non-empty stack, and only if no player
// is already sit. It is recommended for
// this method to be redefined on composition
// if more stuff is to be changed on sit.
func (seat *BaseSeat) Sit(player players.Player, stack uint64) bool {
	if seat.status == Free || player == nil || stack == 0 {
		return false
	} else {
		seat.player = player
		seat.stack = stack
		seat.status = Waiting
		return true
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
	seat.status = Free
	seat.flags = Nothing
	seat.cards = make([]*SeatCard, 0)
	return
}

// Sets the status of this seat. The "free"
// status cannot set, and cannot replaced,
// by this method.
func (seat *BaseSeat) SetStatus(status Status) bool {
	if seat.player == nil || status == Free {
		// Cannot set the status of a table
		// without player, and cannot set
		// the status of a table with player
		// to Free.
		return false
	} else {
		seat.status = status
		return true
	}
}

// Sets a flag on this seat. This can only
// be done to non-free seats.
func (seat *BaseSeat) SetFlag(flag Flags) bool {
	if seat.player == nil {
		return false
	} else {
		seat.flags |= flag
		return true
	}
}

// Clears a flag on this seat. This can only
// be done to non-free seats.
func (seat *BaseSeat) ClearFlag(flag Flags) bool {
	if seat.player == nil {
		return false
	} else {
		seat.flags &= ^flag
		return true
	}
}

// Takes chips from the seat. This can only
// be done to non-free seats that can afford
// the specified chips amount.
func (seat *BaseSeat) TakeChips(chips uint64) bool {
	if seat.player == nil || seat.stack < chips {
		return false
	} else {
		seat.stack -= chips
		return true
	}
}

// Gives chips to the seat. This can only be
// done to non-free seats that can receive
// the amount without overflowing. It is
// recommended that table levels and user
// accounts are defensively restricted to
// reach a case like this.
func (seat *BaseSeat) GiveChips(chips uint64) bool {
	if seat.player == nil || seat.stack > (^(uint64(0) - chips)) {
		return false
	} else {
		seat.stack += chips
		return true
	}
}

// Sets the chips to the seat. This can only
// be done to non-free seats, and care should
// be taken when calling this method, since it
// may cause overflow as well when receiving
// more money or popping the seat in a cash
// 1-table game.
func (seat *BaseSeat) SetStack(chips uint64) bool {
	if seat.player == nil {
		return false
	} else {
		seat.stack = chips
		return true
	}
}

// Adds non-nil cards.
func (seat *BaseSeat) AddCards(seatCards []*SeatCard) bool {
	if len(seatCards) == 0 {
		return true
	}

	for _, seatCard := range seatCards {
		if seatCard == nil || seatCard.Card() == nil {
			return false
		}
	}

	seat.cards = append(seat.cards, seatCards...)
	return true
}

// Removes cards by indices. A []int{-1} array
// clears the hand.
func (seat *BaseSeat) RemoveCards(indices []int) {
	// Test []int{-1} special case.
	if len(indices) == 1 && indices[0] < 0 {
		seat.cards = make([]*SeatCard, 0)
		return
	}

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

// Gets the actual cards from the seat.
func (seat *BaseSeat) Cards(revealed bool) []cards.Card {
	cards := make([]cards.Card, len(seat.cards))
	for index, seatCard := range seat.cards {
		cards[index] = seatCard.Card(revealed)
	}
	return cards
}
