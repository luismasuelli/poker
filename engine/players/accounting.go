package players

import (
	"github.com/luismasuelli/poker-go/engine/assets"
)

// Accounting involves managing players' assets and
// tracking those players by their assets. Assets
// are identified by a string ID and come in two
// flavors: Currencies and tickets. Tickets can be
// present in buy-ins and/or prizes and special
// promos, while currencies can -in addition- be
// the actual currency of cash games.
//
// This contract allows getting, adding, subtracting
// and checking among different types of assets.
//
// Working with nil assets is not defined, and should
// panic.
type Accounting interface {
	// Gets all the (accounting) players having a
	// particular ticket among their assets. By
	// contract, these players should satisfy:
	//   player.Get(Ticket, ticketID) ~ (> 0, true)
	// This is only meaningful only for exclusive
	// tickets (which go for a specific tournament
	// which will NOT repeat) and not for general
	// tickets (which may have multiple purposes.
	// It is not allowed to use a currency-typed
	// asset, and should panic.
	FindHaving(asset assets.Asset) []Accounting
	// Gets how many of an asset this users has.
	// If the asset does not exist, this method
	// must return 0, false. If the asset exists
	// but the player has none, this method must
	// return 0, true. Otherwise, the result must
	// be (quantity), true.
	Get(asset assets.Asset) (uint64, bool)
	// Adds some amount to a particular currency
	// among the player's assets. THIS METHOD MUST
	// PANIC ON OVERFLOW, and is responsibility of
	// the developer to disallow actions or entries
	// to games that could imply the user risks
	// getting money to the point of overflowing.
	// This method should issue a durable command
	// and be concurrency-safe.
	Add(asset assets.Asset, amount uint64)
	// Tries to take some amount from a particular
	// asset. If the player can afford it, this
	// method will return true and subtract the
	// quantity from the underlying asset. If the
	// asset does not exist or the player cannot
	// afford that amount to take, this method
	// returns false. This method should issue a
	// durable command and be concurrency-safe.
	Take(asset assets.Asset, amount uint64)
}
