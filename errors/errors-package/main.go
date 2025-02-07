/*
The Errors Package
The Go standard library provides an “errors” package that makes it easy to deal with errors.

Read the godoc for the errors.New() function, but here’s a simple example:

var err error = errors.New("something went wrong")
Copy icon
Assignment
Textio’s software architects may have overcomplicated the requirements from the last coding assignment… oops. All we needed was a new generic error message that returns the string no dividing by 0 when a user attempts to get us to perform the taboo.

Complete the divide function. Use the errors.New() function to return an error when y == 0 that reads “no dividing by 0”.

Hint
Remember that it’s conventional to return the “zero” values of all other return values when you return a non-nil error in Go.
*/

package main

import (
	"errors"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}
