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
// the winners that are involved with this pot.
// An intersection of such players is considered,
// and the money is divided among them. If no
// given winners are involved in this pot, then
// the result is (0, 0).
func (pot *Pot) Split(winners []seats.Seat) ([]seats.Seat, uint64, uint8) {
	// Compute intersections to keep only the
	// players that are both winners and also
	// involved in the pot.
	involvedWinners := []seats.Seat(nil)
	for _, winner := range winners {
		if _, ok := pot.seats[winner]; ok {
			involvedWinners = append(involvedWinners, winner)
		}
	}

	size := uint64(len(involvedWinners))
	if size == 0 {
		return nil, 0, 0
	} else {
		divided := pot.amount / size
		remainder := uint8(pot.amount % size )
		return involvedWinners, divided, remainder
	}
}