package showdowns

import (
	"github.com/luismasuelli/poker-go/engine/games/tables/seats"
	"github.com/luismasuelli/poker-go/engine/games/tables/pots"
)

// When showdown resolves, a structure of winners
// and losers will be determined. First, one has
// to tell which pots are being played: usually
// only one pot (with eventual side pots) is
// being played, but hi/lo games or pai gow will
// have more than one pot. So we have:
//   pot rule -> (pot structure, pot winners)
// This said, pots are delivered accordingly.

// These players tied. This means: they have the same
// value at showdown, but due to resolution algorithm,
// they will have a priority regarding getting the
// remainder of the chips. The seats at lower indices
// will have more priority regarding getting remainder
// chips than seats at higher indices (this is because
// the lower indices are the ones showing their hands
// first, which means giving away more information
// and so they deserve the remaining chips).
type PodiumPosition []seats.Seat

// These are the full showdown-ranked players. Each
// index will have one or more tying sit players, with
// the 0-index being the greatest showdown winner(s),
// having the best showdown score, and then going down
// one rank each index. This said, only active/all-in
// seats will be included in this ranking (which usually
// means at most 9 seats in total, and at most 4 seats
// per "tie" level).
type Podium []PodiumPosition

// Different games have different showdown modes.
type Mode uint8

// A complete showdown podium involves all the available
// modes. Standard poker games have only one entry in
// this dictionary, while hi/lo games have two entries
// if low hand (8/better) is available, and pai gow games
// have two permanent entries.
type Podiums map[Mode]Podium

// When a complete showdown podium is given, the showdown
// process will also have a mean to split all the existing
// pots among the available showdown modes. Each main or
// side pot will be split accordingly.
type Pots map[Mode][]*pots.Pot

const (
	// Standard showdowns are used when
	// game rules have only one showdown
	// mode.
	Standard Mode = iota
	// "High" and "Low" showdowns are used
	// in Hi/Lo games like Stud and Omaha.
	High
	Low
	// "Front" and "Back" showdowns are used
	// in modes like Pai gow.
	Front
	Back
)

// All the modes to check / iterate for each hand.
var ModesToCheck = []Mode{Standard, High, Low, Front, Back}