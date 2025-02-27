/*
Closing Channels in Go
Channels can be explicitly closed by a sender:

ch := make(chan int)

// do some stuff with the channel

close(ch)

Checking If a Channel Is Closed
Similar to the ok value when accessing data in a map, receivers can check the ok value when receiving from a channel to test if a channel was closed.

v, ok := <-ch

ok is false if the channel is empty and closed.

Don't Send on a Closed Channel
Sending on a closed channel will cause a panic. A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.

Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still be garbage collected if they're unused. You should close channels to indicate explicitly to a receiver that nothing else is going to come across.

Assignment
At Mailio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel. It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent
*/

package main

func countReports(numSentCh chan int) int {
	count := 0
	for {
		value, ok := <-numSentCh

		if !ok {
			break
		}
		count += value
	}

	return count
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}