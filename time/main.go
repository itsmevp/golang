package main

/*
A monotonic clock is a clock that continuously increases at a steady rate, and never goes backwards or jumps ahead.
It is often used in computer systems to measure elapsed time, and is considered to be more reliable than a traditional clock, which may be affected by changes in time due to daylight saving time adjustments, leap years, or system clock adjustments.
Monotonic clocks are used in many applications, such as logging events and timing system operations.
They are particularly useful in distributed systems, where different machines may have slightly different clock settings, and the use of a monotonic clock ensures that all machines are in sync.
*/

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now().UTC()
	fmt.Println(current)                      // 2023-03-06 09:33:51.502829 +0530 IST m=+0.000114209
	fmt.Println(current.Format(time.RFC3339)) // 2023-03-06T04:07:15Z
	/*
		2023-03-06T04:07:15Z, the "Z" at the end stands for "Zulu" time or Coordinated Universal Time (UTC).
		This is a standard way of indicating that the timestamp is in UTC time zone, which is the primary time standard used worldwide in aviation, military, and other industries.
	*/
}
