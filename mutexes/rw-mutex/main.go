/*
RW Mutex
The standard library also exposes a sync.RWMutex

In addition to these methods:

Lock()
Unlock()
The sync.RWMutex also has these methods for concurrent reads:

RLock()
RUnlock()
The sync.RWMutex improves performance in read-intensive processes. Multiple goroutines can safely read from the map simultaneously, as many RLock() calls can occur at the same time. However, only one goroutine can hold a Lock(), and during this time, all RLock() operations are blocked.

Assignment
Let's update our same code from the last assignment, but this time we can speed it up by allowing readers to read from the map concurrently.

Run the new test suite. You'll notice that it hangs forever and you'll need to cancel it.

Update the val method to only lock the mutex for reading.
*/

package main

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}
