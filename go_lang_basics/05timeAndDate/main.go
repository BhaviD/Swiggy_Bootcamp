package main

import (
	"fmt"
	"time"
)

func main()  {
	t := time.Date(2020, time.July, 21, 10, 10, 10, 10, time.UTC)
	fmt.Println("Time is %s \n", t)

	now := time.Now()
	fmt.Printf("Time Now is : %s\n", now)

	// Month
	fmt.Println("Month: ", t.Month())

	// Day
	fmt.Println("Dat: ", t.Day())

	// Week Day
	fmt.Println("Week Day:", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow: %s\n", tomorrow)

	fmt.Printf("Tomorrow Date: %v, Month: %v\n", tomorrow.Day(), tomorrow.Month())

	fmt.Printf("Tomorrow Date: %s", tomorrow.Format("01/02/2006"))
}