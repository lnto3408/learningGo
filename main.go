package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("time is ", n)

	t := time.Date(2022, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go lauched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2022")
	fmt.Printf("the type of parsedTime is %T\n", parsedTime)
}
