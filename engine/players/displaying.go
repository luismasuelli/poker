package players

// Displaying involves the player having
// display data (which is immutable most
// often) that other players will see
// when interacting with this player.
// An example is nickname and/or picture.
type Displaying interface {
	Display() interface{}
}
