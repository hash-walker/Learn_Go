/*
Type Sizes
Integers, uints, floats, and complex numbers all have type sizes.

Whole Numbers (No Decimal)
int  int8  int16  int32  int64
Copy icon
Positive Whole Numbers (No Decimal)
"uint" stands for "unsigned integer".

uint uint8 uint16 uint32 uint64 uintptr
Copy icon
Signed Decimal Numbers
float32 float64
Copy icon
Imaginary Numbers (Rarely Used)
complex64 complex128
Copy icon
What's the Deal With the Sizes?
The size (8, 16, 32, 64, 128, etc) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The "standard" sizes that should be used unless you have a specific performance need (e.g. using less memory) are:

int
uint
float64
complex128
Converting Between Types
Some types can be easily converted like this:

temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)
Copy icon
Casting a float to an integer in this way truncates the floating point portion.

Assignment
Our Textio customers want to know how long they have had accounts with us.

On line 7, create a accountAgeInt variable and assign it the value of accountAgeFloat truncated to an integer.
*/

package main

import "fmt"

func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
