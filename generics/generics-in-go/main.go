/*
Generics in Go
As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many cases. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't care about the types of values stored in the slice. Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.

func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

In Go 1.18 however, support for generics was released, effectively solving this problem!

Type Parameters
Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

In the example above, T is the name of the type parameter for the splitAnySlice function, and we've said that it must match the any constraint, which means it can be anything. This makes sense because the body of the function doesn't care about the types of things stored in the slice.

firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
fmt.Println(firstInts, secondInts)

Assignment
At Mailio we store all the emails for a campaign in memory as a slice. We store payments for a single user in the same way.

Complete the getLast() function. It should be a generic function that returns the last element from a slice, no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

Tip: Zero Value of a Type
Creating a variable that's the zero value of a type is easy:

var myZeroInt int

It's the same with generics, we just have a variable that represents the type:

var myZero T
*/

package main

func getLast[T any](s []T) T {
	var myZero T
	if len(s) == 0 {
		return myZero
	}

	return s[len(s)-1]
}