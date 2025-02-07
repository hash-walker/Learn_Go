/*
Formatting Strings Review
A convenient way to format strings in Go is by using the standard library’s fmt.Sprintf() function. It’s a string interpolation function, similar to Python’s f-strings. The %v substring uses the type’s default formatting, which is often what you want.

Default Values
const name = "Kim"
const age = 22
s := fmt.Sprintf("%v is %v years old.", name, age)
// s = "Kim is 22 years old."
Copy icon
The equivalent Python code:

name = "Kim"
age = 22
s = f"{name} is {age} years old."
# s = "Kim is 22 years old."
Copy icon
Rounding Floats
fmt.Printf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
fmt.Printf("I am %.2f years old", 10.523)
// I am 10.52 years old
Copy icon
Assignment
We need better error logs for our backend developers to help them debug their code.

Complete the getSMSErrorString() function. It should return a string with this format:

SMS that costs $COST to be sent to 'RECIPIENT' can not be sent
Copy icon
COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
RECIPIENT is the stringified representation of the recipient’s phone number
Be sure to include the $ symbol and the single quotes
*/

package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
}
