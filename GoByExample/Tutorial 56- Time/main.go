package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() // get the full time variable
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC) // create a time variable
	p(then)

	p(then.Year())       // Extract Year from time
	p(then.Month())      // Extract Month from time
	p(then.Day())        // Extract Day from time
	p(then.Hour())       // Extract Hours from time
	p(then.Minute())     // Extract Minutes from time
	p(then.Second())     // Extract Seconds from time
	p(then.Nanosecond()) // Extract Nanoseconds from time
	p(then.Location())   // Extract Timezone from time

	p(then.Weekday()) // Calculate which weekday time is

	p(then.Before(now)) // Check if time is before
	p(then.After(now))  // Check if time is after
	p(then.Equal(now))  // Check if tim is equal

	diff := now.Sub(then) // Calculate difference in time between two times in HH:MM:SS
	p(diff)

	p(diff.Hours())       // Difference in Hours
	p(diff.Minutes())     // Difference in Minutes
	p(diff.Seconds())     // Difference in Seconds
	p(diff.Nanoseconds()) // Difference in Nanoseconds

	p(then.Add(diff))  // Add Time
	p(then.Add(-diff)) // Remove Time
}
