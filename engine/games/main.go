package games

import "github.com/luismasuelli/poker-go/engine"

// Games have an ID, a caption and the current
// occupancy (registered / max. amount). Game
// implementations will have different methods
// to implement, besides these ones.
type Game interface {
	ID()        engine.GameID
	Caption()   string
	Occupancy() (uint32, uint32)
}
