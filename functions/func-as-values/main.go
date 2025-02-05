/*
Functions As Values
Go supports first-class and higher-order functions, which are just fancy ways of saying "functions as values". Functions are just another type -- like ints and strings and bools.

Let's assume we have two simple functions:

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}
Copy icon
We can create a new aggregate function that accepts a function as its 4th argument:

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  firstResult := arithmetic(a, b)
  secondResult := arithmetic(firstResult, c)
  return secondResult
}
Copy icon
It calls the given arithmetic function (which could be add or mul, or any other function that accepts two ints and returns an int) and applies it to three inputs instead of two. It can be used like this:

func main() {
	sum := aggregate(2, 3, 4, add)
	// sum is 9
	product := aggregate(2, 3, 4, mul)
	// product is 24
}
Copy icon
Assignment
Complete the reformat function. It takes a message string and a formatter function as input:

Apply the given formatter three times to the message
Add a prefix of TEXTIO: to the result
Return the final string
For example, if the message is "General Kenobi" and the formatter adds a period to the end of the string, the final result should be

TEXTIO: General Kenobi...
*/

package main

func reformat(message string, formatter func(string) string) string {
	
	formattedMessage := formatter(message)
	formattedMessage = formatter(formattedMessage)
	formattedMessage = formatter(formattedMessage)

	return formattedMessage
}
