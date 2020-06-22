package pots

import "github.com/luismasuelli/poker-go/engine/games/tables/seats"

// This is a pot gathered in the table. This
// means it is basically dead money, and also
// considers which seat(s) put money for this
// pot. In the end, only active/all-in seats
// will be considered (for folded and sit-out
// seats can even leave the game). These pots
// will be notified when a seat left its active
// status (i.e. folded [-> sit out [-> left]])
// and will pop the seat from the pot.
type Pot struct {
	amount uint64
	seats  map[seats.Seat]bool
}

// Creates a new pot, with the amount and the
// involved sit players.
func NewPot(amount uint64, involvedSeats []seats.Seat) *Pot {
	involvedSeatsSet := map[seats.Seat]bool{}
	for _, seat := range involvedSeats {
		involvedSeatsSet[seat] = true
	}
	return &Pot{amount, involvedSeatsSet}
}

// This method is invoked when a seat has left
// the hand (or even the game). All-in player
// cannot leave the game, but those active &
// above the pot size may leave the game if
// pushed enough by other active, non all-in,
// sit players.
func (pot *Pot) SeatHasLeft(seat seats.Seat) {
	delete(pot.seats, seat)
}

// Divides equally the amount of the pot among
// all the involved, and remaining active or
// remaining, seats. If there are remainder
// chips, the amount of them will be returned
// in the second return result.
func (pot *Pot) Split() (uint64, uint8) {
	size := uint64(len(pot.seats))
	divided := pot.amount / size
	remainder := uint8(pot.amount % size )
	return divided, remainder
}