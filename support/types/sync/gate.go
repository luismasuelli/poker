package sync

import (
	. "sync"
)

// A gate is a kind of custom wait group based
// on a map. As values are added and removed
// from and to the inner map, the wait group
// gains and loses one unit, respectively. This
// occurs in a concurrency-safe way, and serves
// to issue multiple simultaneous interested
// parties' requests on the same gate: the gate
// will unlock only when all the parties agree.
type Gate struct {
	entries   map[interface{}]bool
	waitGroup WaitGroup
	mutex     Mutex
}

// Tries to add an arbitrary witness value to the
// gate. If already present, this function returns
// false. If not present, the entry is added, the
// wait group is incremented, and this function
// returns true.
//
// This function is concurrency-safe.
func (gate *Gate) Enter(value interface{}) bool {
	gate.mutex.Lock()
	defer gate.mutex.Unlock()

	if _, ok := gate.entries[value]; !ok {
		gate.entries[value] = true
		gate.waitGroup.Add(1)
		return true
	} else {
		return false
	}
}

// Tries to remove an arbitrary witness value from the
// gate. If not present, this function returns false.
// If present, the entry is removed, the wait group
// is decremented, and this function returns true.
//
// This function is concurrency-safe.
func (gate *Gate) Leave(value interface{}) bool {
	gate.mutex.Lock()
	defer gate.mutex.Unlock()

	if _, ok := gate.entries[value]; ok {
		delete(gate.entries, value)
		gate.waitGroup.Done()
		return true
	} else {
		return false
	}
}

// Checks whether a specific witness is locking
// the gate, or not.
func (gate *Gate) Is(value interface{}) bool {
	gate.mutex.Lock()
	defer gate.mutex.Unlock()

	_, ok := gate.entries[value]
	return ok
}

// Locks until the gate is empty (no witnesses).
func (gate *Gate) Wait() {
	gate.waitGroup.Wait()
}

// Creates a new Gate, ready to be used, empty.
func NewGate() *Gate {
	return &Gate{map[interface{}]bool{}, WaitGroup{}, Mutex{}}
}
