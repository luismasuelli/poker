package pot

import (
	"github.com/luismasuelli/poker-go/engine/games/tables/environment"
	"github.com/luismasuelli/poker-go/engine/games/tables/pots"
	"github.com/luismasuelli/poker-go/engine/games/messages/games"
	"github.com/luismasuelli/poker-go/engine/games/rules/showdowns"
	"github.com/luismasuelli/poker-go/engine/games/messages/games/tables"
	"github.com/luismasuelli/poker-go/engine/games/messages/games/tables/seats"
	"time"
)

// Awards all the given pots to their winners. The
// given pots correspond to a single showdown instance
// (e.g. while Hold'Em has a single showdown, other
// modes may have more than one showdown, in a
// permanent or conditional way), so this function
// will be called once per showdown per hand cycle.
//
// For each showdown in certain hand (corresponding
// to certain table / game) this function is called
// exactly once (which could be twice in hi/lo modes).
//
// The seats among pot players are active/all-in that
// did not muck their hands.
func AwardModePots(gameID interface{}, tableID uint32, handID uint64, mode showdowns.Mode,
                   pots []*pots.Pot, podium showdowns.Podium, broadcaster *environment.Broadcaster) {
    // Iterate over all the side pots (and the
    // main point) for a given showdown.
    for potIndex, pot := range pots {
    	// Iterate over all the showdown positions.
    	// The first one having at least one seat
    	// committed with this pot, earn[s] the pot
    	// since it is/are the winner[s].
		for _, sameRankSeats := range podium {
			involvedWinners, amount, remainder := pot.Award(sameRankSeats)
			if len(involvedWinners) != 0 {
				// Divide this pot, and break.
				for index, seat := range involvedWinners {
					// Get the seat ID, player display,
					// money to award (including the
					// remaining chip).
					display := seat.Player().Display()
					seatID := seat.SeatID()
					prize := amount
					if uint8(index) < remainder {
						prize += 1
					}
					// Add the chips to the seat.
					// On overflow: fuck you. How
					// do you overflow when having
					// uint64 as max limit?
					seat.AddStack(prize)
					// Notify with a game message about this
					// prize ($player received from $pot an
					// amount of $chips).
					broadcaster.Notify(games.GameMessage{
						gameID,
						tables.TableMessage{
							tableID,
							seats.SeatMessage{
								seatID,
								seats.PlayerWonChips{
									display, handID, mode,
									uint8(potIndex), prize,
								},
							},
						},
					})
				}
				break
			}
			// Otherwise, keep looking among the podium
			// "levels" until at least one player, in the
			// same level, can claim the pot.
		}
	}
}

// Takes all the given podiums (one per showdown mode), and all
// the given pot sets (one pot set per showdown mode), and awards
// them considering the showdown mode they belong to.
//
// All the all-in/active players that did not muck their cards
// will be present among the podium, and all the modes that exist
// in the game will exist as entries, but if they are not present
// (e.g. for hi/lo games, the "lo" part may be absent) they will
// become "nil" in the podium and empty/nil in the pot lists,
// and will be ignored on iteration.
//
// This function is called once per hand, iterating all of the
// available modes in the podium (and pots).
func AwardPots(gameID interface{}, tableID uint32, handID uint64,
               podiums showdowns.Podiums, potSets showdowns.Pots,
			   broadcaster *environment.Broadcaster,
			   interval time.Duration) {
	for mode, podium := range podiums {
		if podium == nil {
			broadcaster.Notify(games.GameMessage{
				gameID,
				tables.TableMessage{
					tableID,
					tables.Showdown{handID, mode, true},
				},
			})
		} else {
			broadcaster.Notify(games.GameMessage{
				gameID,
				tables.TableMessage{
					tableID,
					tables.Showdown{handID, mode, false},
				},
			})
			potSet := potSets[mode]
			AwardModePots(gameID, tableID, handID, mode, potSet, podium, broadcaster)
			<-time.After(interval)
		}
	}
}