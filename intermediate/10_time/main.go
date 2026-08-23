package main

import (
	"fmt"
	"time"
)

const (
	DATE_FORMAT = "January 02 2006, 03:04 PM"
	TIME_FORMAT = "03:04 PM"
	TIME_INPUT  = "2026-06-12T12:09:36.608Z"
)

func main() {
	/* Current Date & Time */
	now := time.Now()

	fmt.Println(now.Unix())      // Seconds since epoch
	fmt.Println(now.UnixMilli()) // Milliseconds since epoch
	fmt.Println(now.UnixMicro()) // Microseconds since epoch
	fmt.Println(now.UnixNano())  // Nanoseconds since epoch

	//OUTPUT:
	// 1781269241
	// 1781269241268
	// 1781269241268472
	// 1781269241268472000

	/* Specific Date & Time */
	// fmt.Println(time.Date(2024, time.January, 24, 1, 2, 3, 4, time.UTC))

	/* Parsed Date & Time - time.Parse() needs 2 args, 1 is the format of the input and 2 is a actual input */

	t, e := time.Parse(time.RFC3339Nano, TIME_INPUT)

	if e != nil {
		fmt.Println(e)
		return
	}

	/* Formatting Parsed Date & Time */
	fmt.Println(t.Format(DATE_FORMAT)) // OUTPUT: June 12 2026, 12:09 PM

	/* Add & Remove Date to Parsed Date & Time */
	// t is immutable. so t.AddDate() doesn't mutate the actual item rather it returns a new updated value

	// #1 Store the value and then use it
	addedDate := t.AddDate(0, 0, 1)
	fmt.Println(addedDate.Format(DATE_FORMAT)) // OUTPUT: June 13 2026, 12:09 PM

	// #2 or use directly
	fmt.Println(t.AddDate(0, 0, -10).Format(DATE_FORMAT)) // OUTPUT: June 02 2026, 12:09 PM

	/* Add & Remove Time to Parsed Date & Time */
	addedTime := t.Add(1 * time.Hour)
	fmt.Println(addedTime.Format(TIME_FORMAT)) // OUTPUT: 01:09 PM

	fmt.Println(time.Now().Unix() - addedTime.Unix())

	/* Change Timezones */
	istZone, _ := time.LoadLocation("Asia/Kolkata")

	fmt.Println("UTC Time:", t.Format(time.DateTime))
	fmt.Println("IST Time:", t.In(istZone).Format(time.DateTime))
	// OUTPUT: 5 hours 30 minutes difference
	// UTC Time: 2026-06-12 12:09:36
	// IST Time: 2026-06-12 17:39:36

	/* Comparing Date & Time */
	t1 := time.Now()
	t2 := t1.Add(time.Hour)

	fmt.Println(t1.Before(t2)) // true
	fmt.Println(t1.Equal(t2))  // false
	fmt.Println(t1.After(t2))  // false

	/* Time Elapsed Count*/
	start := time.Now()

	time.Sleep(2 * time.Second)

	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
