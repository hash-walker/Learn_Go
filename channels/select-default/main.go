/*
Select Default Case
The default case in a select statement executes immediately if no other channel has a value ready. A default case stops the select statement from blocking.

select {
case v := <-ch:
    // use v
default:
    // receiving from ch would block
    // so do something else
}

Tickers
time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
time.After() sends a value once after the duration has passed.
time.Sleep() blocks the current goroutine for the specified duration of time.
The functions take a time.Duration as an argument. For example:

time.Tick(500 * time.Millisecond)

If you don't add time.Millisecond (or another unit), it will default to nanoseconds. That's — taking a wild guess here — probably faster than you want it to be.

Read-Only Channels
A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:

func main() {
    ch := make(chan int)
    readCh(ch)
}

func readCh(ch <-chan int) {
    // ch can only be read from
    // in this function
}

Write-Only Channels
The same goes for write-only channels, but the arrow's position moves.

func writeCh(ch chan<- int) {
    // ch can only be written to
    // in this function
}

Assignment
Like all good back-end engineers, we frequently save backup snapshots of the Mailio database.

Complete the saveBackups function.

It should read values from the snapshotTicker and saveAfter channels simultaneously and continuously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds. After all, we want to show in the logs that the snapshot service is running.
*/

package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
	select {
		case <-snapshotTicker:
			takeSnapshot(logChan)	
		case <-saveAfter:
			saveSnapshot(logChan)
			return 
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
	}
}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
