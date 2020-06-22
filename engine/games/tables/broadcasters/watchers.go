package broadcasters

import (
	"github.com/luismasuelli/poker-go/engine/misc"
)

// Broadcasts a given message to all
// the listeners watching a table.
// Meant to be used inside a table.
type WatchersBroadcaster struct {
	watchers map[misc.Notifiable]bool
}

// Creates a new watchers broadcaster.
func NewWatchersBroadcaster() *WatchersBroadcaster {
	return &WatchersBroadcaster{map[misc.Notifiable]bool{}}
}

// Registers a new watcher. Returns false
// if already registered, or nil.
func (watchersBroadcaster *WatchersBroadcaster) Register(watcher misc.Notifiable) bool {
	if watcher == nil {
		return false
	} else if _, ok := watchersBroadcaster.watchers[watcher]; ok {
		return false
	} else {
		watchersBroadcaster.watchers[watcher] = true
		return true
	}
}

// Unregisters a watcher. Returns false if absent.
func (watchersBroadcaster *WatchersBroadcaster) Unregister(watcher misc.Notifiable) bool {
	if _, ok := watchersBroadcaster.watchers[watcher]; !ok {
		return false
	} else {
		delete(watchersBroadcaster.watchers, watcher)
		return true
	}
}

// Broadcasts the message to all the watchers.
func (watchersBroadcaster *WatchersBroadcaster) Notify(message interface{}) {
	for watcher, _ := range watchersBroadcaster.watchers {
		func() {
			defer func() { recover() }()
			watcher.Notify(message)
		}()
	}
}
