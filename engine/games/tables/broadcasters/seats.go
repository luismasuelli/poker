package broadcasters

import "github.com/luismasuelli/poker-go/engine/games/tables/seats"

// Broadcasts a given message to all
// the players sitting in the table.
// Meant to be used inside a table.
type SeatsBroadcaster struct {
	receivers []seats.Seat
}

// Creates a broadcaster for all the seats.
func NewSeatsBroadcaster(seats []seats.Seat) *SeatsBroadcaster {
	return &SeatsBroadcaster{seats}
}

// Broadcasts the message to all the seats.
func (seatsBroadcaster *SeatsBroadcaster) Notify(message interface{}) {
	for _, seat := range seatsBroadcaster.receivers {
		if player := seat.Player(); player != nil {
			func() {
				defer func() { recover() }()
				player.Notify(message)
			}()
		}
	}
}
