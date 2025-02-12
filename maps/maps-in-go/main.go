/*
Maps
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24
Copy icon
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
Copy icon
The len() function works on a map, it returns the total number of key/value pairs.

ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2
Copy icon
Assignment
We can speed up our contact-info lookups by using a map!

Key-based map lookup: O(1)
Slice brute-force search: O(n)
Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error. A user struct just contains a user’s name and phone number. The first name in the names slice pairs with the first phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string “invalid sizes”.
*/

package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {

	if len(names) != len(phoneNumbers){
		return map[string]user{},errors.New("invalid sizes")
	}

	newMap := make(map[string]user)

	for i, name := range names{
		newMap[name] = user{name: name, phoneNumber: phoneNumbers[i]}
	}

	return newMap,nil

}

type user struct {
	name        string
	phoneNumber int
}
