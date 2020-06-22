package environment

import (
	"github.com/luismasuelli/poker-go/engine/games/tables/seats"
	"github.com/luismasuelli/poker-go/engine/misc"
)

// Combines the seats and watchers
// broadcasters in just one.
type Broadcaster struct {
	seats    []seats.Seat
	watchers map[misc.Notifiable]bool
	parent   misc.Notifiable
}

// Creates a new environment broadcaster with
// the involved seats and an empty list of
// watchers.
func NewBroadcaster(seats []seats.Seat, parent misc.Notifiable) *Broadcaster {
	return &Broadcaster{seats, map[misc.Notifiable]bool{}, parent}
}

// Notifies to both sit players and watchers.
func (broadcaster *Broadcaster) Notify(message interface{}) {
	for _, seat := range broadcaster.seats {
		if player := seat.Player(); player != nil {
			func() {
				defer func() { recover() }()
				player.Notify(message)
			}()
		}
	}
	for watcher, _ := range broadcaster.watchers {
		func() {
			defer func() { recover() }()
			watcher.Notify(message)
		}()
	}
	broadcaster.parent.Notify(message)
}

// Registers a new watcher. Returns false
// if already registered, or nil.
func (broadcaster *Broadcaster) Register(watcher misc.Notifiable) bool {
	if watcher == nil {
		return false
	} else if _, ok := broadcaster.watchers[watcher]; ok {
		return false
	} else {
		broadcaster.watchers[watcher] = true
		return true
	}
}

// Unregisters a watcher. Returns false if absent.
func (broadcaster *Broadcaster) Unregister(watcher misc.Notifiable) bool {
	if _, ok := broadcaster.watchers[watcher]; !ok {
		return false
	} else {
		delete(broadcaster.watchers, watcher)
		return true
	}
}