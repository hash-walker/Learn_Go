/*
Append
The built-in append function is used to dynamically add elements to a slice:

func append(slice []Type, elems ...Type) []Type
Copy icon
If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.

Notice that append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
Copy icon
Assignment
We’ve been asked to add a feature that extracts costs for a given day.

Complete the getDayCosts() function using the append() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day’s costs.

If there are no costs for that day, return an empty slice.
*/

package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	costSlice := []float64{}

	for _, cost := range costs{
		if cost.day == day {
			costSlice = append(costSlice, cost.value)
		}
	}

	return costSlice
}


