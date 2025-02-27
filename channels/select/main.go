/*
Select
Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.

select {
case i, ok := <- chInts:
    fmt.Println(i)
case s, ok := <- chStrings:
    fmt.Println(s)
}

The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time one is chosen randomly. The ok variable in the example above refers to whether or not the channel has been closed by the sender yet.

Assignment
Complete the logMessages function.

Use an infinite for loop and a select statement to log the emails and sms messages as they come in order across the two channels. Add a condition to return from the function when one of the two channels closes, whichever is first.

Use the logSms and logEmail functions to log the messages.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case v, ok := <- chEmails:
			if !ok {
				return 
			}
			logEmail(v)

		case v, ok := <- chSms:
			if !ok {
				return 
			}
			logSms(v)
		}
	}
}

// don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func test(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}

func main() {
	rand.Seed(0)
	test(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	test(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}
