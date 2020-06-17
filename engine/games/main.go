package games

// Games have an ID, a caption and the current
// occupancy (registered / max. amount). Game
// implementations will have different methods
// to implement, besides these ones.
//
// About the lifecycle: cash games live forever,
// and in the worst cases they can be halted and
// restarted later. Tournaments, on the other
// hand, do not last forever. They end, and will
// not be created again.
type Game interface {
	ID()        interface{}
	Caption()   string
	Occupancy() (uint32, uint32)
}
