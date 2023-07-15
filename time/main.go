// https://www.sohamkamani.com/golang/time/

package main

import (
	"fmt"
	"time"
)

/*
A monotonic clock is a clock that continuously increases at a steady rate, and never goes backwards or jumps ahead.
It is often used in computer systems to measure elapsed time, and is considered to be more reliable than a traditional clock, which may be affected by changes in time due to daylight saving time adjustments, leap years, or system clock adjustments.
Monotonic clocks are used in many applications, such as logging events and timing system operations.
They are particularly useful in distributed systems, where different machines may have slightly different clock settings, and the use of a monotonic clock ensures that all machines are in sync.
*/

func main() {
	currentTime := time.Now()
	// 2023-07-15 10:41:08.846415 +0530 IST m=+0.000127584
	fmt.Println(currentTime.String())

	// 1689398447
	fmt.Println(currentTime.Unix())

	// 1689398635151
	fmt.Println(currentTime.UnixMilli())

	// 1689398535567768
	fmt.Println(currentTime.UnixMicro())

	// https://pkg.go.dev/time#pkg-constants
	fmt.Println(currentTime.Format(time.RFC3339))
}
