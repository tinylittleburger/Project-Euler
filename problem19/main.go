package main

import (
	"fmt"
	"time"
)

func mustParse(layout, value string) time.Time {
	dt, err := time.Parse(layout, value)

	if err != nil {
		panic(err)
	}

	return dt
}

func main() {
	start := mustParse("Jan 2 2006", "Jan 1 1901")
	end := mustParse("Jan 2 2006", "Jan 1 2001")
	count := 0

	for dt := start; !dt.Equal(end); dt = dt.AddDate(0, 0, 1) {
		if dt.Day() == 1 && dt.Weekday() == time.Sunday {
			count++
		}
	}

	fmt.Println(count)
}
