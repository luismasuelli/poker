package sync

import (
	. "sync"
)

// This type is used when retrieved the status of a GatedLoop
// instance by calling its Status() method.
type GatedLoopStatus uint8

const (
	Created GatedLoopStatus = iota
	Waiting
	Running
	Destroyed
)

// A gated coroutine loops indefinitely in a loop performing
// two steps in order:
// - Test the gate.
// - Run the body.
// It also has before/after callbacks to run custom code and
// test for potential panics (those callbacks are optional).
type GatedLoop struct {
	gate   *Gate
	once   Once
	before func()
	while  func() bool
	do     func()
	after  func(panicked interface{})
	status GatedLoopStatus
}

// Runs the gated loop only once. The first time it runs,
// this method returns true and runs the goroutine. Such
// goroutine is the actual execution of the loop. Future
// calls to this method will not run the goroutine and
// will return false.
func (gated *GatedLoop) Run() bool {
	var result bool
	gated.once.Do(func() {
		result = true
		go func() {
			gated.status = Running
			defer func() {
				gated.status = Destroyed
				if gated.after != nil {
					gated.after(recover())
				}
			}()
			if gated.before != nil {
				gated.before()
			}
			for gated.while() {
				gated.status = Waiting
				gated.gate.Wait()
				gated.status = Running
				gated.do()
			}
		}()
	})
	return result
}

// Returns the status of a gated loop.
func (gated *GatedLoop) Status() GatedLoopStatus {
	return gated.status
}

// Returns the gate of a gated loop, so witnesses can be
// added and removed, and make the loop wait after each
// respective execution.
func (gated *GatedLoop) Gate() *Gate {
	return gated.gate
}

// Creates a gated runner, with its own gate and optional
// before/after callbacks.
func NewGated(before func(), while func() bool, do func(), after func(panicked interface{})) *GatedLoop {
	return &GatedLoop{NewGate(), Once{}, before, while, do, after, Created}
}
