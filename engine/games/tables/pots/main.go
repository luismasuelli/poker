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

// Gets the pot amount.
func (pot *Pot) Amount() uint64 {
	return pot.amount
}

// Splits the pot in sub-pots of the given
// quantities and the same participants. If
// the quantities surpass the whole pot, a
// final pot is added with the remaining
// quantities.
func (pot *Pot) Split(amounts ...uint64) []*Pot {
	remainingAmount := pot.amount
	splitPots := make([]*Pot, 0)

	// Iterating involves consuming all the amounts
	// that can be consumed (i.e. are lower than the
	// total amount in the pot). The moment we find
	// a pot equal or larger to the remaining pot,
	// or the moment we consume all the pots, we
	// exit the loop.
	for _, amount := range amounts {
		if amount < remainingAmount {
			splitPots = append(splitPots, &Pot{amount, pot.seats})
			remainingAmount -= amount
		} else {
			break
		}
	}
	// In the end, we still have some remaining
	// amount. This will make us have such amount
	// be added to the end of the list, as a new
	// and final pot.
	splitPots = append(splitPots, &Pot{remainingAmount, pot.seats})
	return splitPots
}

// Divides equally the amount of the pot among
// the winners that are involved with this pot.
// An intersection of such players is considered,
// and the money is divided among them. If no
// given winners are involved in this pot, then
// the result is (0, 0).
func (pot *Pot) Award(winners []seats.Seat) ([]seats.Seat, uint64, uint8) {
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