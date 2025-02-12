/*
Mutations
Insert an Element
m[key] = elem
Copy icon
Get an Element
elem = m[key]
Copy icon
Delete an Element
delete(m, key)
Copy icon
Check If a Key Exists
elem, ok := m[key]
Copy icon
If key is in m, then ok is true and elem is the value as expected.

If key is not in the map, then ok is false and elem is the zero value for the map’s element type.

Assignment
It’s important to keep up with privacy regulations and to respect our user’s data. We need a function that will delete user records.

Complete the deleteIfNecessary function. The user struct has a scheduledForDeletion field that determines if they are scheduled for deletion or not.

If the user doesn’t exist in the map, return the error not found.
If they exist but aren’t scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
Note on Passing Maps
Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original, we don’t have a copy.
*/

package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	value, ok := users[name]

	if !ok{
		return false, errors.New("not found")
	}

	if value.scheduledForDeletion{
		delete(users, value.name)
		return true, nil
	}else{
		return false, nil
	}
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}