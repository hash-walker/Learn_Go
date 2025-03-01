/*
Constraints
Sometimes you need the logic in your generic function to know something about the types it operates on. The example we used in the first exercise didn't need to know anything about the types in the slice, so we used the built-in any constraint:

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

Constraints are just interfaces that allow us to write generics that only operate within the constraints of a given interface type. In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.

Creating a Custom Constraint
Let's take a look at the example of a concat function. It takes a slice of values and concatenates the values into a string. This should work with any type that can represent itself as a string, even if it's not a string under the hood. For example, a user struct can have a .String() that returns a string with the user's name and age.

type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}

Assignment
We have different kinds of "line items" that we charge our customer's credit cards for. Line items can be things like "subscriptions" or "one-time payments" for email usage.

Complete the chargeForLineItem function. First, it should check if the user has a balance with enough funds to be able to pay for the cost of the newItem. If they don't, then return an "insufficient funds" error and zero values for the other return values.

If they do have enough funds:

Add the line item to the user's history by appending the newItem to the slice of oldItems. This new slice is your first return value.
Calculate the user's new balance by subtracting the cost of the new item from their balance. This is your second return value.
*/

package main

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	var zeroValue []T
	if balance < newItem.GetCost() {
		return zeroValue, 0, errors.New("insufficient funds")
	}

	oldItems = append(oldItems, newItem)

	return oldItems, balance - newItem.GetCost(), nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}
