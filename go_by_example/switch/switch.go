package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's Weekday")
	}

	t := time.Now()
	fmt.Println(t)

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

}
