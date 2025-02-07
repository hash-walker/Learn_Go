/*
The Error Interface
Because errors are just interfaces, you can build your own custom types that implement the error interface. Here’s an example of a userError struct that implements the error interface:

type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}
Copy icon
It can then be used as an error:

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}
Copy icon
Assignment
Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It’s become such a problem that we need to make a new type of error for division by zero.

Update the code so that the divideError type implements the error interface. Its Error() method should just return a string formatted in the following way:

can not divide DIVIDEND by zero
Copy icon
Where DIVIDEND is the actual dividend of the divideError. Use the %v verb to format the dividend as a float.
*/

package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (zErr divideError) Error() string{
	return fmt.Sprintf("can not divide %v by zero", zErr.dividend)
}  

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
