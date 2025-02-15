/*
Channels
Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.

Create a Channel
Like maps and slices, channels must be created before use. They also use the same make keyword:

ch := make(chan int)
Copy icon
Send Data to a Channel
ch <- 69
Copy icon
The <- operator is called the channel operator. Data flows in the direction of the arrow. This operation will block until another goroutine is ready to receive the value.

Receive Data from a Channel
v := <-ch
Copy icon
This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value in the channel to be read.

Blocking and Deadlocks
A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

Assignment
Run the program. You'll see that it deadlocks and never exits. The sendIsOld function is trying to send on a channel, but no other goroutines are running that can accept the value from the channel.

Fix the deadlock by spawning a goroutine to send the "is old" values.
*/

package main

import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
