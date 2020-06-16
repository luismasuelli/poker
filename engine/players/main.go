package players

// A player makes use of the four interfaces
// in this package to identify itself, show
// itself, manage its assets and receive
// notifications.
//
// No particular implementation will be given
// for players, but instead implementations
// must follow the guidelines to appropriately
// implement the player contract and use them.
type Player interface {
	Identified
	Accounting
	Displaying
	Notifiable
}
