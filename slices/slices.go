//Package slices has functions that return slices populated with the number of seconds in a minute
//or number of minutes in a day for example.
//The syntax for the name of the functions is, first letter "number of x", second letter "in y"
package slices

import (
	"strconv"
)

func SecondsMinutes() []string {
	slice := make([]string, 60)

	for i := 1; i < 60; i++ {
		slice[i] = strconv.Itoa(i)
	}

	return slice
}
