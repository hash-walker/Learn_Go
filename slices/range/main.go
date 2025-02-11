/*
Range
Go provides syntactic sugar to iterate easily over elements of a slice:

for INDEX, ELEMENT := range SLICE {
}
Copy icon
Note: the element is a copy of the value at that index.

For example:

fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape
Copy icon
Assignment
We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function. If it finds any bad words in the message it should return the index of the first bad word in the msg slice. This will help us filter out naughty words from our messaging system. If no bad words are found, return -1 instead.

Use the range keyword.

Tip
Remember, you can ignore returned values with an underscore _ instead of creating an unused variable.
*/

package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, msgWord := range msg{
		for _, bdw := range badWords{
			if bdw == msgWord{
				return i
			}
		}
	}
	return -1
}